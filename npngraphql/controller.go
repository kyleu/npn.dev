package npngraphql

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kyleu/npn/npncontroller/routes"

	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"

	"logur.dev/logur"

	"emperror.dev/errors"
	"github.com/graphql-go/graphql"
)

func RoutesGraphQL(app npnweb.AppInfo, r *mux.Router) {
	g := r.Path(routes.Name(npncore.KeyGraphQL)).Subrouter()
	g.Methods(http.MethodPost).Handler(routes.AddContext(r, app, http.HandlerFunc(GraphQLRun))).Name(routes.Name(npncore.KeyGraphQL))
	g.Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(GraphiQL))).Name(routes.Name(npncore.KeyGraphiQL))
	r.Path(routes.Name(npncore.KeyVoyager, "query")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(GraphQLVoyagerQuery))).Name(routes.Name(npncore.KeyVoyager, "query"))
	r.Path(routes.Name(npncore.KeyVoyager, "mutation")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(GraphQLVoyagerMutation))).Name(routes.Name(npncore.KeyVoyager, "mutation"))
}

func RoutesGraphQLAdmin(app npnweb.AppInfo, r *mux.Router) {
	g := r.Path(routes.Adm(npncore.KeyGraphQL)).Subrouter()
	g.Methods(http.MethodPost).Handler(routes.AddContext(r, app, http.HandlerFunc(AdminGraphQLRun))).Name(npnweb.AdminLink(npncore.KeyGraphQL))
	g.Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(AdminGraphiQL))).Name(npnweb.AdminLink(npncore.KeyGraphiQL))
	r.Path(routes.Adm(npncore.KeyVoyager, "query")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(AdminGraphQLVoyagerQuery))).Name(npnweb.AdminLink(npncore.KeyVoyager, "query"))
	r.Path(routes.Adm(npncore.KeyVoyager, "mutation")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(AdminGraphQLVoyagerMutation))).Name(npnweb.AdminLink(npncore.KeyVoyager, "mutation"))
}

func GraphQLRun(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		return run(w, r, ctx)
	})
}

func AdminGraphQLRun(w http.ResponseWriter, r *http.Request) {
	npncontroller.AdminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		return run(w, r, ctx)
	})
}

func run(w http.ResponseWriter, r *http.Request, ctx *npnweb.RequestContext) (string, error) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		e := ErrorResponseJSON(graphQLService.Logger, errors.Wrap(err, "cannot read JSON body for GraphQL"))
		return graphQLResponse(w, e, ctx.Logger)
	}
	err = r.Body.Close()
	if err != nil {
		e := ErrorResponseJSON(graphQLService.Logger, errors.Wrap(err, "cannot close body for GraphQL"))
		return graphQLResponse(w, e, ctx.Logger)
	}

	var req map[string]interface{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		e := ErrorResponseJSON(graphQLService.Logger, errors.Wrap(err, "error decoding JSON body for GraphQL"))
		return graphQLResponse(w, e, ctx.Logger)
	}

	op := npncore.MapGetString(req, "operationName", ctx.Logger)
	query := npncore.MapGetString(req, "query", ctx.Logger)
	v := mapGetMap(req, "variables", ctx.Logger)

	res, err := graphQLService.Run(op, query, v, ctx)
	if err != nil {
		e := ErrorResponseJSON(graphQLService.Logger, errors.Wrap(err, "error running GraphQL"))
		return graphQLResponse(w, e, ctx.Logger)
	}

	return graphQLResponse(w, res, ctx.Logger)
}

func graphQLResponse(w http.ResponseWriter, res *graphql.Result, logger logur.Logger) (string, error) {
	return npncontroller.RespondJSON(w, "", res, logger)
}

func mapGetMap(m map[string]interface{}, key string, logger logur.Logger) map[string]interface{} {
	retEntry := npncore.GetEntry(m, key, logger)
	ret, ok := retEntry.(map[string]interface{})
	if !ok {
		logger.Warn(fmt.Sprintf("key [%v] in map is type [%T], not map[string]interface{}", key, retEntry))
		return nil
	}
	return ret
}
