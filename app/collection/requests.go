package collection

import (
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/npncore"
	"path"
	"strings"
)

func (s *Service) ListRequests(key string) ([]string, error) {
	p := path.Join(rootDir, key, "requests")
	return s.files.ListJSON(p), nil
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

func (s *Service) SaveRequest(coll string, req *request.Request) error {
	msg := npncore.ToJSON(req, s.logger)
	p := path.Join(rootDir, coll, "requests", req.Key + ".json")
	return s.files.WriteFile(p, msg, true)
}

func (s *Service) DeleteRequest(coll string, key string) error {
	p := path.Join(rootDir, coll, "requests", key + ".json")
	return s.files.RemoveRecursive(p)
}
