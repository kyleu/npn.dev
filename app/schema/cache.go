package schema

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/app/util"
	"logur.dev/logur"
)

type Cache struct {
	files  *util.FileLoader
	data   map[string]*Schema
	logger logur.Logger
}

func NewCache(logger logur.Logger) *Cache {
	return &Cache{files: util.NewFileLoader(logger), data: make(map[string]*Schema), logger: logger}
}

func (c *Cache) List() []string {
	return c.files.ListJSON("schema")
}

func (c *Cache) Summary(key string) (*Summary, error) {
	content, err := c.files.ReadFile("schema/" + key + ".json")
	if err != nil {
		return nil, errors.Wrap(err, "unable to find schema file with key ["+key+"]")
	}
	tgt := &Summary{}
	util.FromJSON([]byte(content), tgt, c.logger)
	return tgt, nil
}

func (c *Cache) Load(key string) (*Schema, error) {
	content, err := c.files.ReadFile("schema/" + key + ".json")
	if err != nil {
		return nil, errors.Wrap(err, "unable to find schema file with key ["+key+"]")
	}
	tgt := &Schema{}
	util.FromJSON([]byte(content), tgt, c.logger)
	return tgt, nil
}

func (c *Cache) Save(sch *Schema, overwrite bool) error {
	return c.files.WriteFile("schema/"+sch.Key+".json", util.ToJSON(sch, c.logger), overwrite)
}

func (c *Cache) Summaries() (Summaries, error) {
	schemaKeys := c.List()
	ret := make(Summaries, 0, len(schemaKeys))
	for _, key := range schemaKeys {
		sch, err := c.Summary(key)
		if err != nil {
			return nil, err
		}
		ret = append(ret, sch)
	}
	return ret, nil
}
