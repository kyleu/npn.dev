package collection

import (
	"os"
	"path"

	"github.com/gofrs/uuid"

	"emperror.dev/errors"
	"github.com/kyleu/npn/npncore"
)

func (s *Service) Save(userID *uuid.UUID, originalKey string, newKey string, title string, description string) error {
	originalKey = npncore.Slugify(originalKey)
	if len(newKey) == 0 {
		newKey = "new"
	}
	newKey = npncore.Slugify(newKey)

	var orig *Collection
	var err error

	if len(originalKey) > 0 {
		orig, err = s.Load(userID, originalKey)
		if err != nil {
			return errors.Wrap(err, "unable to load original collection ["+originalKey+"]")
		}
		if orig != nil && originalKey != newKey {
			o := path.Join(s.files.Root(), s.dirFor(userID), originalKey)
			n := path.Join(s.files.Root(), s.dirFor(userID), newKey)
			err := os.Rename(o, n)
			if err != nil {
				return errors.Wrap(err, "unable to rename original collection ["+originalKey+"] in path ["+o+"]")
			}
		}
	}

	n := &Collection{
		Key:         newKey,
		Title:       title,
		Description: description,
	}

	if orig == nil {
		n.Owner = "system"
	} else {
		n.Owner = orig.Owner
		n.RequestOrder = orig.RequestOrder
	}
	n.Path = newKey

	p := path.Join(s.dirFor(userID), newKey, "collection.json")
	content := npncore.ToJSONBytes(n, s.logger, true)
	err = s.files.WriteFile(p, content, true)
	if err != nil {
		return errors.Wrap(err, "unable to save collection ["+newKey+"]")
	}

	return nil
}
