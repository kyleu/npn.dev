package cloudbeaver

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/app/model/schema"
)

func (l *Loader) load(sch *schema.Schema, parentType string, path string, encountered []string) (*navNode, error) {
	var err error
	var children []*navNode
	switch parentType {
	case "Table", "View":
		// noop
	default:
		children, err = l.navChildren(sch, path, encountered)
	}
	if err != nil {
		return nil, errors.Wrap(err, "error loading children of ["+path+"]")
	}

	var filterKeys []string
	switch parentType {
	case "Schema", "Database":
		filterKeys = []string{"Schemas", "Tables", "Views", "Indexes"}
	}
	if len(filterKeys) > 0 {
		filtered := make([]*navNode, 0)
		for _, child := range children {
			for _, key := range filterKeys {
				if key == child.Type {
					filtered = append(filtered, child)
				}
			}
		}
		children = filtered
	}

	return &navNode{
		ID:       path,
		Type:     parentType,
		Children: children,
	}, nil
}

func (l *Loader) navChildren(sch *schema.Schema, path string, encountered []string) ([]*navNode, error) {
	tgt, err := l.call(gqlNavChildren(path), nil)
	if err != nil {
		return nil, err
	}
	if tgt == nil {
		return []*navNode{}, nil
	}
	kidsRaw, ok := tgt.(map[string]interface{})["navNodeChildren"]
	if !ok {
		return nil, errors.New("no [navNodeChildren] present in response")
	}
	kids := make([]*navNode, 0, len(kidsRaw.([]interface{})))
	for _, kidRaw := range kidsRaw.([]interface{}) {
		kidMap := kidRaw.(map[string]interface{})
		kidID := kidMap["id"].(string)
		kidType := kidMap["nodeType"].(string)
		for _, e := range encountered {
			if e == kidID {
				return nil, errors.New("cycle detected")
			}
		}
		kid, err := l.load(sch, kidType, kidID, append(encountered, kidID))
		if err != nil {
			return nil, errors.Wrap(err, "error loading child of ["+path+"]")
		}
		kids = append(kids, kid)
	}
	return kids, nil
}
