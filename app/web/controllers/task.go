package controllers

import (
	"github.com/kyleu/npn/app/model/task"
	"net/http"

	"github.com/kyleu/npn/app/model/project"

	"github.com/gorilla/mux"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/app/web"
	"github.com/kyleu/npn/app/web/act"
	"github.com/kyleu/npn/gen/templates"
)

func TaskRun(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		p, t, err := loadTask(r, ctx)
		if err != nil {
			return act.EResp(err)
		}
		tsk := task.FromString(t.T)
		tr, err := ctx.App.RunTask(tsk, p.Key, t.Options)
		if err != nil {
			return act.EResp(err, "error running task ["+t.Key+"] for project ["+p.Key+"]")
		}
		return act.T(templates.TaskResult(tr, ctx, w))
	})
}

func TaskAdd(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		p, t, err := loadTask(r, ctx)
		if err != nil {
			return act.EResp(err)
		}
		t = &project.Task{Key: "new", T: t.Key, Options: nil}
		return act.T(templates.TaskForm(p, t, ctx, w))
	})
}

func TaskEdit(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		p, t, err := loadTask(r, ctx)
		if err != nil {
			return act.EResp(err)
		}
		return act.T(templates.TaskForm(p, t, ctx, w))
	})
}

func TaskSave(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		projectKey := mux.Vars(r)[util.KeyKey]
		// taskKey := mux.Vars(r)[util.KeyTask]
		redir := ctx.Route(util.KeyProject+".detail", util.KeyKey, projectKey)
		return act.FlashAndRedir(true, "Saved task", redir, w, r, ctx)
	})
}

func loadTask(r *http.Request, ctx *web.RequestContext) (*project.Project, *project.Task, error) {
	projectKey := mux.Vars(r)[util.KeyKey]
	taskKey := mux.Vars(r)[util.KeyTask]

	p, err := ctx.App.Projects.Load(projectKey)
	if err != nil {
		return nil, nil, err
	}

	t := p.Tasks.Get(taskKey)
	if t == nil {
		t = &project.Task{Key: taskKey, T: taskKey}
	}
	ctx.Breadcrumbs = projectBreadcrumbs(ctx, ctx.Route(util.KeyProject+".detail", util.KeyKey, projectKey), projectKey, "", t.Key)

	return p, t, nil
}
