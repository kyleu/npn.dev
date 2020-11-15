package request

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npncore"
	"os"
	"path"
	"path/filepath"
	"time"
)

const shouldSaveHistory = true

func (s *Service) saveHistory(userID *uuid.UUID, coll string, req *Request, p string, msg string) error {
	if shouldSaveHistory {
		hp := historyPath(userID, coll, req.Key)
		now := time.Now()
		hfn := path.Join(hp, npncore.ToDateString(&now)+".json")
		hd := filepath.Dir(hfn)
		err := s.files.CreateDirectory(hd)
		if err != nil {
			return errors.Wrap(err, "unable to create request history directory ["+hd+"]")
		}

		x, _ := os.Stat(p)
		if x == nil {
			err = s.files.WriteFile(hfn, []byte(msg), true)
		} else {
			err = s.files.CopyFile(p, hfn)
		}
		if err != nil {
			return errors.Wrap(err, "unable to create request history ["+hp+"]")
		}
	}

	err := s.files.WriteFile(p, []byte(msg), true)
	if err != nil {
		return errors.Wrap(err, "unable to write file")
	}

	return nil
}

func historyPath(userID *uuid.UUID, coll string, key string) string {
	return path.Join("users", userID.String(), "collections", coll, "history", "requests", key)
}
