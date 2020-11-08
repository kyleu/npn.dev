package collection

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofrs/uuid"

	"emperror.dev/errors"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/npncore"
)

const shouldSaveHistory = true

type RequestSummary struct {
	Key         string `json:"key,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
	Order       int    `json:"order,omitempty"`
}

func (r *RequestSummary) TitleWithFallback() string {
	if len(r.Title) == 0 {
		return r.Key
	}
	return r.Title
}

type RequestSummaries []*RequestSummary

func (s *Service) ListRequests(userID *uuid.UUID, c string) (RequestSummaries, error) {
	p := path.Join(dirFor(userID), c, "requests")
	files := s.files.ListJSON(p)
	ret := make(RequestSummaries, 0, len(files))
	for idx, rk := range files {
		r, err := s.LoadRequest(userID, c, rk)
		if err != nil {
			return nil, errors.Wrap(err, "error loading request ["+rk+"]")
		}
		url := ""
		if r.Prototype != nil {
			url = r.Prototype.URLString()
		}
		ret = append(ret, &RequestSummary{
			Key:         r.Key,
			Title:       r.Title,
			Description: r.Description,
			URL:         url,
			Order:       idx,
		})
	}
	return ret, nil
}

func (s *Service) LoadRequest(userID *uuid.UUID, c string, f string) (*request.Request, error) {
	f = strings.TrimSuffix(f, ".json")
	p := path.Join(dirFor(userID), c, "requests", f+".json")
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

func (s *Service) SaveRequest(userID *uuid.UUID, coll string, originalKey string, req *request.Request) error {
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

func requestPath(userID *uuid.UUID, coll string, key string) string {
	return path.Join(dirFor(userID), coll, "requests", key+".json")
}

func historyPath(userID *uuid.UUID, coll string, key string) string {
	return path.Join(dirFor(userID), coll, "history", "requests", key)
}

func (s *Service) DeleteRequest(userID *uuid.UUID, coll string, key string) error {
	p := path.Join(dirFor(userID), coll, "requests", key+".json")
	return s.files.RemoveRecursive(p)
}
