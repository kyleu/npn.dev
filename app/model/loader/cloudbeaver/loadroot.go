package cloudbeaver

import (
	"emperror.dev/errors"
)

func (l *Loader) loadRoot() (map[string]string, error) {
	ret := make(map[string]string)

	err := l.navChildrenRoot(ret)
	if err != nil {
		return nil, errors.Wrap(err, "error loading root children")
	}

	return ret, nil
}

func (l *Loader) navChildrenRoot(ret map[string]string) error {
	tgt, err := l.call(gqlNavChildren("/"), nil)
	if err != nil {
		return err
	}
	if tgt == nil {
		return nil
	}
	kidsRaw, ok := tgt.(map[string]interface{})["navNodeChildren"]
	if !ok {
		return errors.New("no [navNodeChildren] present in response")
	}
	for _, kidRaw := range kidsRaw.([]interface{}) {
		kidMap := kidRaw.(map[string]interface{})
		kidID := kidMap["id"].(string)
		ret[kidID] = kidID
	}
	return nil
}
