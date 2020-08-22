package npncore

import (
	"fmt"
	"net/url"
	"strings"
)

type Data map[string]interface{}

func (d Data) HasKey(s string) bool {
	_, ok := d[s]
	return ok
}

func (d Data) String() string {
	ret := make([]string, 0, len(d))
	for k, v := range d {
		ret = append(ret, fmt.Sprintf("%v = %v", k, v))
	}
	return strings.Join(ret, ", ")
}

func (d Data) ToQueryString() string {
	params := url.Values{}
	for k, v := range d {
		params.Add(k, fmt.Sprintf("%v", v))
	}
	return params.Encode()
}

func (d Data) Clone() Data {
	ret := make(Data, len(d))
	for k, v := range d {
		ret[k] = v
	}
	return ret
}

func (d Data) GetPath(path string) interface{} {
	parts := strings.Split(path, ".")
	return getPath(d, parts)
}

func getPath(i interface{}, path []string) interface{} {
	if len(path) == 0 {
		return i
	}
	t, ok := i.(map[string]interface{})
	if ok {
		ret, ok := t[path[0]]
		if !ok {
			return nil
		}
		return getPath(ret, path[1:])
	}
	return nil
}

func (d Data) GetString(k string) string {
	v, ok := d[k]
	if !ok {
		return ""
	}
	return fmt.Sprintf("%v", v)
}
