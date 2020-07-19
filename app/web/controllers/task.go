package controllers

import (
	"github.com/kyleu/npn/app/task"
	"net/http"

	"github.com/kyleu/npn/app/project"

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
		td := &project.TaskDefinition{Key: "new", T: t.Key, Options: nil}
		return act.T(templates.TaskForm(p, td, ctx, w))
	})
}

func TaskEdit(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		p, td, err := loadTask(r, ctx)
		if err != nil {
			return act.EResp(err)
		}
		return act.T(templates.TaskForm(p, td, ctx, w))
	})
}

func TaskSave(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		projectKey := mux.Vars(r)[util.KeyKey]
		p, err := ctx.App.Projects.Load(projectKey)
		if err != nil || p == nil {
			return act.EResp(err)
		}

		_ = r.ParseForm()
		taskType := r.Form["taskType"][0]
		originalTaskKey := r.Form["originalTaskKey"][0]
		newTaskKey := r.Form["newTaskKey"][0]

		println(originalTaskKey + " / " + newTaskKey)

		var originalTask *project.TaskDefinition
		if len(originalTaskKey) == 0 || originalTaskKey == "new" {
			originalTask = &project.TaskDefinition{}
		} else {
			originalTask = p.Tasks.Get(originalTaskKey)
		}
		newTask := originalTask.Clone()
		newTask.Key = newTaskKey
		newTask.T = taskType

		p.Tasks = p.Tasks.Plus(originalTaskKey, newTask)

		err = ctx.App.Projects.Save(p.Key, p, true)
		if err != nil {
			return act.EResp(err, "cannot save project")
		}

		redir := ctx.Route(util.KeyProject+".detail", util.KeyKey, projectKey)
		return act.FlashAndRedir(true, "Saved task", redir, w, r, ctx)
	})
}

func loadTask(r *http.Request, ctx *web.RequestContext) (*project.Project, *project.TaskDefinition, error) {
	projectKey := mux.Vars(r)[util.KeyKey]
	taskKey := mux.Vars(r)[util.KeyTask]

	p, err := ctx.App.Projects.Load(projectKey)
	if err != nil {
		return nil, nil, err
	}

	t := p.Tasks.Get(taskKey)
	if t == nil {
		t = &project.TaskDefinition{Key: taskKey, T: taskKey}
	}
	ctx.Breadcrumbs = projectBreadcrumbs(ctx, ctx.Route(util.KeyProject+".detail", util.KeyKey, projectKey), projectKey, "", t.Key)

	return p, t, nil
}
