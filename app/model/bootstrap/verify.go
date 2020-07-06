package bootstrap

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/app/model/project"
	"github.com/kyleu/npn/app/util"
	"logur.dev/logur"
	"os"
)

func Verify(prototype *Prototype, cfg *project.Project, logger logur.Logger) error {
	_, err := os.Open(cfg.RootPath)
	if err != nil {
		return errors.Wrap(err, "destination ["+cfg.RootPath+"] is missing")
	}

	for _, cmd := range prototype.BuildCmds {
		parsed, err := util.Template(cmd, cfg)
		if err != nil {
			return err
		}
		err = util.RunProcess(parsed, cfg.RootPath, logger, nil, nil, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
