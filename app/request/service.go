package request

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnuser"
	"logur.dev/logur"
	"os"
	"path"
	"strings"
)

type Service struct {
	files  npncore.FileLoader
	logger logur.Logger
}

func NewService(f npncore.FileLoader, logger logur.Logger) *Service {
	return &Service{files: f, logger: logger}
}

func (s *Service) ListRequests(userID *uuid.UUID, c string) (Summaries, error) {
	p := dirFor(userID, c)
	files := s.files.ListJSON(p)
	ret := make(Summaries, 0, len(files))
	for idx, rk := range files {
		r, err := s.LoadRequest(userID, c, rk)
		if err != nil {
			return nil, errors.Wrap(err, "error loading request ["+rk+"]")
		}
		url := ""
		if r.Prototype != nil {
			url = r.Prototype.URLString()
		}
		ret = append(ret, &Summary{
			Key:         r.Key,
			Title:       r.Title,
			Description: r.Description,
			URL:         url,
			Order:       idx,
		})
	}
	return ret, nil
}

func (s *Service) LoadRequest(userID *uuid.UUID, c string, f string) (*Request, error) {
	f = strings.TrimSuffix(f, ".json")
	p := path.Join(dirFor(userID, c), f+".json")
	content, err := s.files.ReadFile(p)
	if err != nil {
		return nil, err
	}
	ret, err := FromString(f, string(content))
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s *Service) SaveRequest(userID *uuid.UUID, coll string, originalKey string, req *Request) error {
	originalKey = npncore.Slugify(originalKey)
	slug := npncore.Slugify(req.Key)
	if slug != req.Key {
		s.logger.Debug(fmt.Sprintf("renaming request key from [%v] to [%v]", req.Key, slug))
		if len(req.Title) == 0 {
			req.Title = req.Key
		}
		req.Key = slug
	}

	shouldDelete := len(originalKey) > 0 && req.Key != originalKey

	if shouldDelete {
		orig, err := s.LoadRequest(userID, coll, req.Key)
		if err == nil && orig != nil {
			return errors.New("file already exists in collection [" + coll + "] with key [" + req.Key + "]")
		}
	}

	p := requestPath(userID, coll, req.Key)

	if shouldDelete {
		o := path.Join(s.files.Root(), requestPath(userID, coll, originalKey))
		n := path.Join(s.files.Root(), p)
		err := os.Rename(o, n)
		if err != nil {
			return errors.Wrap(err, "unable to rename original request ["+originalKey+"] in path ["+o+"]")
		}
	}

	msg := npncore.ToJSON(req, s.logger)

	err := s.saveHistory(userID, coll, req, p, msg)
	if err != nil {
		return errors.Wrap(err, "unable to save history")
	}

	err = s.files.WriteFile(p, []byte(msg), true)
	if err != nil {
		return errors.Wrap(err, "unable to write file")
	}

	return nil
}

func (s *Service) DeleteRequest(userID *uuid.UUID, coll string, key string) error {
	p := path.Join(dirFor(userID, coll), key+".json")
	return s.files.RemoveRecursive(p)
}

func requestPath(userID *uuid.UUID, coll string, key string) string {
	return path.Join(dirFor(userID, coll), key+".json")
}

func dirFor(userID *uuid.UUID, coll string) string {
	if userID == nil || *userID == npnuser.SystemUserID {
		return "collections"
	}
	return path.Join("users", userID.String(), "collections", coll, "requests")
}
