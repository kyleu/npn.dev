package controllers

import (
	"net/http"

	"github.com/kyleu/npn/gen/templates"

	"github.com/gorilla/mux"
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/imprt"
	"github.com/kyleu/libnpn/npncontroller"
	"github.com/kyleu/libnpn/npncore"
	"github.com/kyleu/libnpn/npnweb"
)

func ImportForm(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = "Import"
		return npncontroller.T(templates.ImportForm(ctx, w))
	})
}

func ImportDetail(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		key := mux.Vars(r)[npncore.KeyKey]
		svc := app.Svc(ctx.App)
		cfg, out, err := svc.Import.Load(key)
		if err != nil {
			return npncontroller.EResp(err)
		}
		ctx.Title = "Import"
		ret := map[string]interface{}{"cfg": cfg, "out": out}
		return npncontroller.RespondJSON(w, "", ret, ctx.Logger)
	})
}

func ImportUpload(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		err := r.ParseMultipartForm(1024 * 1024 * 2)
		if err != nil {
			return npncontroller.EResp(err, "unable to parse multipart form")
		}

		v := r.MultipartForm.File["file"]

		files := make([]imprt.File, 0)
		for _, file := range v {
			ct, ok := file.Header["Content-Type"]
			if !ok || len(ct) == 0 {
				ct = []string{"text/plain"}
			}
			files = append(files, imprt.File{
				Filename:    file.Filename,
				Size:        file.Size,
				ContentType: ct[0],
			})
		}

		svc := app.Svc(ctx.App)
		importKey := npncore.RandomString(16)
		err = svc.Import.Create(importKey, files)
		if err != nil {
			return npncontroller.EResp(err, "unable to create import")
		}

		for _, file := range v {
			f, err := file.Open()
			if err != nil {
				return npncontroller.EResp(err, "unable to open uploaded file ["+file.Filename+"]")
			}
			err = svc.Import.WriteFile(importKey, file.Filename, f)
			if err != nil {
				return npncontroller.EResp(err, "unable to write import file ["+file.Filename+"]")
			}
		}

		return ctx.Route("import.detail", npncore.KeyKey, importKey), nil
	})
}
