package parser

import (
	"emperror.dev/errors"
	"encoding/xml"
	"fmt"
	"io"
	"logur.dev/logur"
	"os"
	"strings"
)

func parseXML(path string, onStart func(xml.StartElement, *xml.Decoder) error) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()
	d := xml.NewDecoder(f)

	for {
		tok, err := d.Token()
		if tok == nil || err == io.EOF {
			break // EOF means we're done.
		} else if err != nil {
			return errors.Wrap(err, "error decoding XML token: %+v")
		}

		switch e := tok.(type) {
		case xml.StartElement:
			err = onStart(e, d)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func attrValue(se xml.StartElement, name string, logger logur.Logger) string {
	for _, a := range se.Attr {
		if a.Name.Local == name {
			return a.Value
		}
	}
	available := make([]string, 0, len(se.Attr))
	for _, a := range se.Attr {
		available = append(available, a.Name.Local)
	}
	msg := "missing attribute [%v] in element [%v] from available [%v]"
	logger.Warn(fmt.Sprintf(msg, name, strings.Join(available, ", "), se.Name.Local))
	return ""
}
