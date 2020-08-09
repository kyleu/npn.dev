package bootstrap

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/app/project"
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
	"os"
)

func Verify(prototype *Prototype, cfg *project.Project, logger logur.Logger) error {
	_, err := os.Open(cfg.RootPath)
	if err != nil {
		return errors.Wrap(err, "destination ["+cfg.RootPath+"] is missing")
	}

	for _, cmd := range prototype.BuildCmds {
		parsed, err := npncore.Template(cmd, cfg)
		if err != nil {
			return err
		}
		_, err = npncore.RunProcess(parsed, cfg.RootPath, logger, nil, nil, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
