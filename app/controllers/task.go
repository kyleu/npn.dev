package controllers

import (
	"net/http"
	"strings"

	"emperror.dev/errors"
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/project"
	"github.com/kyleu/npn/app/task"
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"

	"github.com/gorilla/mux"
	"github.com/kyleu/npn/gen/templates"
)

const projectDetailRoute = "project.detail"

func TaskRun(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		p, t, err := loadTask(r, ctx)
		if err != nil {
			return npncontroller.EResp(err)
		}
		tsk := task.FromString(t.T)
		tr := run(ctx.App, tsk, p.Key, t.Options)
		return npncontroller.T(templates.TaskResults(tr, ctx, w))
	})
}

func TaskRunAll(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		p, err := loadProject(r, ctx)
		if err != nil {
			return npncontroller.EResp(err)
		}
		ret := task.Results{}
		schemata, err := app.Schemata(ctx.App).LoadAll(p.SchemaKeys)
		if err != nil {
			return npncontroller.EResp(err)
		}
		for _, td := range p.Tasks {
			tsk := task.FromString(td.T)
			r := task.RunTask(p, schemata, tsk, td.Options, ctx.Logger)
			ret = append(ret, r...)
		}
		ctx.Breadcrumbs = projectBreadcrumbs(ctx, ctx.Route(projectDetailRoute, npncore.KeyKey, p.Key), p.Key, "", "all")
		return npncontroller.T(templates.TaskResults(ret, ctx, w))
	})
}

func TaskAdd(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		p, t, err := loadTask(r, ctx)
		if err != nil {
			return npncontroller.EResp(err)
		}
		td := &project.TaskDefinition{Key: "new", T: t.Key, Options: nil}
		return npncontroller.T(templates.TaskForm(p, td, ctx, w))
	})
}

func TaskEdit(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		p, td, err := loadTask(r, ctx)
		if err != nil {
			return npncontroller.EResp(err)
		}
		return npncontroller.T(templates.TaskForm(p, td, ctx, w))
	})
}

func TaskDelete(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		p, td, err := loadTask(r, ctx)
		if err != nil {
			return npncontroller.EResp(err)
		}
		p.Tasks = p.Tasks.Without(td.Key)

		err = app.Projects(ctx.App).Save(p.Key, p, true)
		if err != nil {
			return npncontroller.EResp(err, "cannot save project")
		}

		redir := ctx.Route(projectDetailRoute, npncore.KeyKey, p.Key)
		return npncontroller.FlashAndRedir(true, "Deleted task ["+td.Key+"]", redir, w, r, ctx)
	})
}

func TaskSave(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		projectKey := mux.Vars(r)[npncore.KeyKey]
		p, err := app.Projects(ctx.App).Load(projectKey)
		if err != nil || p == nil {
			return npncontroller.EResp(err)
		}

		_ = r.ParseForm()
		taskType := r.Form["taskType"][0]
		originalTaskKey := r.Form["originalTaskKey"][0]
		newTaskKey := r.Form["newTaskKey"][0]

		options := make(npncore.Entries, 0, len(r.Form))
		for k, v := range r.Form {
			if strings.HasPrefix(k, "opt-") && len(v[0]) > 0 {
				options = append(options, &npncore.Entry{K: strings.TrimPrefix(k, "opt-"), V: v[0]})
			}
		}

		var originalTask *project.TaskDefinition
		if len(originalTaskKey) == 0 || originalTaskKey == "new" {
			originalTask = &project.TaskDefinition{}
		} else {
			originalTask = p.Tasks.Get(originalTaskKey)
		}
		newTask := originalTask.Clone()
		newTask.Key = newTaskKey
		newTask.T = taskType
		newTask.Options = options

		p.Tasks = p.Tasks.Replacing(originalTaskKey, newTask)

		err = app.Projects(ctx.App).Save(p.Key, p, true)
		if err != nil {
			return npncontroller.EResp(err, "cannot save project")
		}

		redir := ctx.Route(projectDetailRoute, npncore.KeyKey, projectKey)
		return npncontroller.FlashAndRedir(true, "Saved task", redir, w, r, ctx)
	})
}

func loadProject(r *http.Request, ctx *npnweb.RequestContext) (*project.Project, error) {
	projectKey := mux.Vars(r)[npncore.KeyKey]
	return app.Projects(ctx.App).Load(projectKey)
}

func loadTask(r *http.Request, ctx *npnweb.RequestContext) (*project.Project, *project.TaskDefinition, error) {
	taskKey := mux.Vars(r)["task"]

	p, err := loadProject(r, ctx)
	if err != nil {
		return nil, nil, err
	}

	t := p.Tasks.Get(taskKey)
	if t == nil {
		t = &project.TaskDefinition{Key: taskKey, T: taskKey}
	}
	ctx.Breadcrumbs = projectBreadcrumbs(ctx, ctx.Route(projectDetailRoute, npncore.KeyKey, p.Key), p.Key, "", t.Key)

	return p, t, nil
}

func run(a npnweb.AppInfo, t task.Task, projectKey string, options npncore.Entries) task.Results {
	proj, err := app.Projects(a).Load(projectKey)
	if err != nil {
		err = errors.Wrap(err, projErr(projectKey))
		return task.ErrorResults(t, proj, options, err)
	}
	schemata, err := app.Schemata(a).LoadAll(proj.SchemaKeys)
	if err != nil {
		err = errors.Wrap(err, "cannot load schema for project ["+proj.Key+"]")
		return task.ErrorResults(t, proj, options, err)
	}
	return task.RunTask(proj, schemata, t, options, a.Logger())
}
