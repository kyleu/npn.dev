package npncore

import (
	"encoding/csv"
	"fmt"
	"net/url"
	"strings"

	"emperror.dev/errors"
)

// A helper class, wrapping map[string]interface{}
type Data map[string]interface{}

// Returns a boolean indicating if the data has a value associated to the provided key
func (d Data) HasKey(s string) bool {
	_, ok := d[s]
	return ok
}

// Returns a string representation of this Data's contents, mostly used for debugging
func (d Data) String() string {
	ret := make([]string, 0, len(d))
	for k, v := range d {
		ret = append(ret, fmt.Sprintf("%v = %v", k, v))
	}
	return strings.Join(ret, ", ")
}

// Returns a string representation of this Data's contents, URL-encoded in query string format
func (d Data) ToQueryString() string {
	params := url.Values{}
	for k, v := range d {
		params.Add(k, fmt.Sprint(v))
	}
	return params.Encode()
}

// Returns a new shallow copy of this Data
func (d Data) Clone() Data {
	ret := make(Data, len(d))
	for k, v := range d {
		ret[k] = v
	}
	return ret
}

// Returns the value associated to this path, traversing child objects if a path like "x.y.z" is passed
func (d Data) GetPath(path string) interface{} {
	r := csv.NewReader(strings.NewReader(path))
	r.Comma = '.'
	fields, err := r.Read()
	if err != nil {
		return err
	}
	return getPath(d, fields)
}

func getPath(i interface{}, path []string) interface{} {
	if len(path) == 0 {
		return i
	}
	switch t := i.(type) {
	case Data:
		ret, ok := t[path[0]]
		if !ok {
			return nil
		}
		return getPath(ret, path[1:])
	case map[string]interface{}:
		ret, ok := t[path[0]]
		if !ok {
			return nil
		}
		return getPath(ret, path[1:])
	default:
		return nil
	}
}

// Sets the value associated to this path, traversing child objects if a path like "x.y.z" is passed
func (d Data) SetPath(path string, val interface{}) interface{} {
	r := csv.NewReader(strings.NewReader(path))
	r.Comma = '.'
	fields, err := r.Read()
	if err != nil {
		return err
	}
	return setPath(d, fields, val)
}

func setPath(i interface{}, path []string, val interface{}) error {
	work := i
	for idx, p := range path {
		if idx == len(path)-1 {
			switch t := work.(type) {
			case Data:
				t[p] = val
			case map[string]interface{}:
				t[p] = val
			default:
				return errors.New(fmt.Sprintf("unhandled [%T]", t))
			}
		} else {
			switch t := work.(type) {
			case Data:
				t[p] = map[string]interface{}{}
				work = t[p]
			case map[string]interface{}:
				t[p] = map[string]interface{}{}
				work = t[p]
			default:
				return errors.New(fmt.Sprintf("unhandled [%T]", t))
			}
		}
	}
	return nil
}

// Returns a string representation of the value associated to the provided key
func (d Data) GetString(key string) string {
	return fmt.Sprint(d.GetPath(key))
}
