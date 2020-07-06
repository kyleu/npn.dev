package project

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/app/model/data"
	"github.com/kyleu/npn/app/util"
	"logur.dev/logur"
)

type Cache struct {
	files  *data.FileLoader
	data   map[string]*Project
	logger logur.Logger
}

func NewCache(logger logur.Logger) *Cache {
	return &Cache{files: data.NewFileLoader(logger), data: make(map[string]*Project), logger: logger}
}

func (c *Cache) List() []string {
	return c.files.ListJSON("project")
}

func (c *Cache) Summary(key string) (*Summary, error) {
	content, err := c.files.ReadFile("project/" + key + ".json")
	if err != nil {
		return nil, errors.Wrap(err, "unable to find project file with key ["+key+"]")
	}
	tgt := &Summary{}
	util.FromJSON([]byte(content), tgt, c.logger)
	return tgt, nil
}

func (c *Cache) Load(key string) (*Project, error) {
	content, err := c.files.ReadFile("project/" + key + ".json")
	if err != nil {
		return nil, errors.Wrap(err, "unable to find project file with key ["+key+"]")
	}
	tgt := &Project{}
	util.FromJSON([]byte(content), tgt, c.logger)
	return tgt, nil
}

func (c *Cache) Save(originalKey string, p *Project, overwrite bool) error {
	if len(originalKey) > 0 && originalKey != "new" {
		newProj, _ := c.Load(p.Key)
		if originalKey != p.Key && newProj != nil {
			return errors.New("remove the existing [" + originalKey + "] project before you overwrite it with this one")
		}
	}
	err := c.files.WriteFile("project/"+p.Key+".json", util.ToJSON(p, c.logger), overwrite)
	if err != nil {
		return errors.Wrap(err, "unable to write project")
	}
	if len(originalKey) > 0 && originalKey != "new" && originalKey != p.Key {
		err = c.Remove(originalKey)
		if err != nil {
			return errors.Wrap(err, "cannot remove original project")
		}
	}
	return nil
}

func (c *Cache) Summaries() (Summaries, error) {
	projectKeys := c.List()
	summaries := make(Summaries, 0, len(projectKeys))
	for _, key := range projectKeys {
		sch, err := c.Summary(key)
		if err != nil {
			return nil, err
		}
		summaries = append(summaries, sch)
	}
	return summaries, nil
}

func (c *Cache) Remove(key string) error {
	return errors.Wrap(c.files.Remove("project/"+key+".json"), "unable to remove project ["+key+"] file")
}
