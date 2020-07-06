package util

import (
	"fmt"

	"emperror.dev/errors"
)

func MapFromPairs(x ...interface{}) (map[interface{}]interface{}, error) {
	if len(x)%2 != 0 {
		return nil, errors.New(fmt.Sprintf("observed [%v] args, need an even number", len(x)))
	}
	ret := make(map[interface{}]interface{}, len(x)/2)
	for i := 0; i < len(x)-1; i += 2 {
		k := x[i]
		v := x[i+1]
		ret[k] = v
	}
	return ret, nil
}

func StringKeyMapFromPairs(x ...interface{}) (map[string]interface{}, error) {
	curr, err := MapFromPairs(x...)
	if err != nil {
		return nil, err
	}
	ret := make(map[string]interface{}, len(curr))
	for k, v := range curr {
		s, ok := k.(string)
		if !ok {
			return nil, errors.New(fmt.Sprintf("key [%v: %T] is not a string", k, k))
		}
		ret[s] = v
	}
	return ret, nil
}

func StringMapFromPairs(x ...interface{}) (map[string]string, error) {
	curr, err := StringKeyMapFromPairs(x...)
	if err != nil {
		return nil, err
	}
	ret := make(map[string]string, len(curr))
	for k, v := range curr {
		s, ok := v.(string)
		if !ok {
			return nil, errors.New(fmt.Sprintf("value [%v: %T] is not a string", k, k))
		}
		ret[k] = s
	}
	return ret, nil
}
