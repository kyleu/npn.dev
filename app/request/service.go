package request

import (
	"fmt"
	"os"
	"path"
	"strings"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/libnpn/npncore"
	"github.com/kyleu/libnpn/npnuser"
	"logur.dev/logur"
)

type Service struct {
	multiuser bool
	files     npncore.FileLoader
	logger    logur.Logger
}

func NewService(multiuser bool, f npncore.FileLoader, logger logur.Logger) *Service {
	return &Service{multiuser: multiuser, files: f, logger: logger}
}

func (s *Service) LoadAll(userID *uuid.UUID, c string) (Requests, error) {
	p := s.dirFor(userID, c)
	files := s.files.ListJSON(p)
	ret := make(Requests, 0, len(files))
	for _, rk := range files {
		r, err := s.Load(userID, c, rk)
		if err != nil {
			return nil, errors.Wrap(err, "error loading request ["+rk+"]")
		}
		ret = append(ret, r)
	}
	return ret, nil
}

func (s *Service) List(userID *uuid.UUID, c string) (Summaries, error) {
	requests, err := s.LoadAll(userID, c)
	if err != nil {
		return nil, err
	}
	ret := make(Summaries, 0, len(requests))
	for idx, r := range requests {
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

func (s *Service) Load(userID *uuid.UUID, c string, f string) (*Request, error) {
	f = strings.TrimSuffix(f, ".json")
	p := path.Join(s.dirFor(userID, c), f+".json")
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

func (s *Service) Save(userID *uuid.UUID, coll string, originalKey string, req *Request) error {
	originalKey = npncore.Slugify(originalKey)
	if len(req.Key) == 0 {
		req.Key = "new"
	}
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
		orig, err := s.Load(userID, coll, req.Key)
		if err == nil && orig != nil {
			return errors.New("request file already exists in collection [" + coll + "] with key [" + req.Key + "]")
		}
	}

	p := s.requestPath(userID, coll, req.Key)

	if shouldDelete {
		o := path.Join(s.files.Root(), s.requestPath(userID, coll, originalKey))
		n := path.Join(s.files.Root(), p)
		err := os.Rename(o, n)
		if err != nil {
			return errors.Wrap(err, "unable to rename original request ["+originalKey+"] in path ["+o+"]")
		}
	}

	msg := npncore.ToJSONBytes(req, s.logger, true)

	err := s.saveHistory(userID, coll, req, p, msg)
	if err != nil {
		return errors.Wrap(err, "unable to save history")
	}

	err = s.files.WriteFile(p, msg, true)
	if err != nil {
		return errors.Wrap(err, "unable to write file")
	}

	return nil
}

func (s *Service) Delete(userID *uuid.UUID, coll string, key string) error {
	p := path.Join(s.dirFor(userID, coll), key+".json")
	return s.files.RemoveRecursive(p)
}

func (s *Service) requestPath(userID *uuid.UUID, coll string, key string) string {
	return path.Join(s.dirFor(userID, coll), key+".json")
}

func (s *Service) dirFor(userID *uuid.UUID, coll string) string {
	if (!s.multiuser) || userID == nil || *userID == npnuser.SystemUserID {
		return path.Join("collections", coll, "requests")
	}
	return path.Join("users", userID.String(), "collections", coll, "requests")
}
