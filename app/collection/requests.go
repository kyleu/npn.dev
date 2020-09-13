package collection

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/npncore"
	"path"
	"strings"
)

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
	ret, err := request.FromString(f, content)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s *Service) SaveRequest(coll string, originalKey string, req *request.Request) error {
	if len(originalKey) > 0 && req.Key != originalKey {
		orig, err := s.LoadRequest(coll, req.Key)
		if err == nil && orig != nil {
			return errors.New("file already exists in collection [" + coll + "] with key [" + req.Key + "]")
		}
	}

	msg := npncore.ToJSON(req, s.logger)
	p := path.Join(rootDir, coll, "requests", req.Key+".json")
	err := s.files.WriteFile(p, msg, true)
	if err != nil {
		return errors.Wrap(err, "unable to write file")
	}

	if len(originalKey) > 0 && req.Key != originalKey {
		err = s.DeleteRequest(coll, originalKey)
		if err != nil {
			return errors.Wrap(err, "unable to delete ["+coll+"/"+originalKey+"]")
		}
	}

	return nil
}

func (s *Service) DeleteRequest(coll string, key string) error {
	p := path.Join(rootDir, coll, "requests", key+".json")
	return s.files.RemoveRecursive(p)
}
