package collection

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/npncore"
	"os"
	"path"
	"strings"
	"time"
)

const shouldSaveHistory = true

func (s *Service) ListRequests(c string) (request.Requests, error) {
	p := path.Join(rootDir, c, "requests")
	files := s.files.ListJSON(p)
	ret := make(request.Requests, 0, len(files))
	for _, rk := range files {
		r, err := s.LoadRequest(c, rk)
		if err != nil {
			return nil, errors.Wrap(err, "error loading request ["+rk+"]")
		}
		ret = append(ret, r)
	}
	return ret, nil
}

func (s *Service) LoadRequest(c string, f string) (*request.Request, error) {
	f = strings.TrimSuffix(f, ".json")
	p := path.Join(rootDir, c, "requests", f+".json")
	content, err := s.files.ReadFile(p)
	if err != nil {
		return nil, err
	}
	ret, err := request.FromString(f, string(content))
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s *Service) SaveRequest(coll string, originalKey string, req *request.Request) error {
	shouldDelete := len(originalKey) > 0 && req.Key != originalKey

	if shouldDelete {
		orig, err := s.LoadRequest(coll, req.Key)
		if err == nil && orig != nil {
			return errors.New("file already exists in collection [" + coll + "] with key [" + req.Key + "]")
		}
	}

	p := requestPath(coll, req.Key)

	if shouldDelete {
		o := path.Join(s.files.Root(), requestPath(coll, originalKey))
		n := path.Join(s.files.Root(), p)
		err := os.Rename(o, n)
		if err != nil {
			return errors.Wrap(err, "unable to rename original request [" + originalKey + "] in path [" + o + "]")
		}
	}

	if shouldSaveHistory {
		hp := historyPath(coll, req.Key)
		now := time.Now()
		hfn := path.Join(hp, npncore.ToDateString(&now) + ".json")
		err := s.files.CopyFile(p, hfn)
		if err != nil {
			return errors.Wrap(err, "unable to create request history ["+hp+"]")
		}
	}

	msg := npncore.ToJSON(req, s.logger)
	err := s.files.WriteFile(p, []byte(msg), true)
	if err != nil {
		return errors.Wrap(err, "unable to write file")
	}

	return nil
}

func requestPath(coll string, key string) string {
	return path.Join(rootDir, coll, "requests", key+".json")
}

func historyPath(coll string, key string) string {
	return path.Join(rootDir, coll, "history", "requests", key)
}

func (s *Service) DeleteRequest(coll string, key string) error {
	p := path.Join(rootDir, coll, "requests", key+".json")
	return s.files.RemoveRecursive(p)
}
