package npncore

import (
	"fmt"
	"sort"
	"strings"

	"github.com/gofrs/uuid"
	"logur.dev/logur"

	"emperror.dev/errors"
)

// Returns the value associated to the key in the provided map, logging a detailed message and returning nil if not found
func GetEntry(m map[string]interface{}, key string, logger logur.Logger) interface{} {
	retEntry, ok := m[key]
	if !ok {
		if logger != nil {
			keys := make([]string, 0, len(m))
			for k := range m {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			logger.Warn(fmt.Sprintf("no key [%v] in map from available keys [%v]", key, strings.Join(keys, ", ")))
		}
		return nil
	}
	return retEntry
}

// Returns the string value associated to the key in the provided map, logging a detailed message and returning nil if not found or value is not a string
func MapGetString(m map[string]interface{}, key string, logger logur.Logger) string {
	retEntry := GetEntry(m, key, logger)
	ret, ok := retEntry.(string)
	if !ok {
		logger.Warn(fmt.Sprintf("key [%v] in map is type [%T], not string", key, retEntry))
		return ""
	}
	return ret
}

// Returns the bool value associated to the key in the provided map, logging a detailed message and returning nil if not found or value is not a boolean
func MapGetBool(m map[string]interface{}, key string, logger logur.Logger) bool {
	retEntry := GetEntry(m, key, logger)
	ret, ok := retEntry.(bool)
	if !ok {
		logger.Warn(fmt.Sprintf("key [%v] in map is type [%T], not bool", key, retEntry))
		return false
	}
	return ret
}

// Returns the UUID value associated to the key in the provided map, logging a detailed message and returning nil if not found or value is not a UUID
func MapGetUUID(m map[string]interface{}, key string, logger logur.Logger) *uuid.UUID {
	retEntry := GetEntry(m, key, logger)
	ret, ok := retEntry.(uuid.UUID)
	if !ok {
		s, ok := retEntry.(string)
		if !ok {
			logger.Warn(fmt.Sprintf("key [%v] in map is type [%T], not uuid", key, retEntry))
			return nil
		}
		r, e := uuid.FromString(s)
		if e != nil {
			logger.Warn(fmt.Sprintf("key [%v] in map with value [%v] is not a valid uuid", key, s))
			return nil
		}
		ret = r
	}
	return &ret
}

// Returns the map resulting for applying the arguments, alternating between keys and values
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

// Returns the map resulting for applying the arguments, alternating between string keys and values
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

// Returns the map resulting for applying the arguments, alternating between string keys and string values
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

// Returns a string representing the provided map, in HStore format
func MapToDBString(m map[string]string) string {
	ret := make([]string, 0, len(m))
	for k, v := range m {
		ret = append(ret, fmt.Sprintf("%v=>%v", k, v))
	}
	return strings.Join(ret, ",")
}

// Returns a map representing the provided HStore format string
func StringToDBMap(s string) (map[string]string, error) {
	if len(strings.TrimSpace(s)) == 0 {
		return make(map[string]string, 0), nil
	}

	parts := strings.Split(s, ",")
	ret := make(map[string]string, len(parts))

	for _, p := range parts {
		idx := strings.Index(p, "=>")
		if idx == -1 {
			return nil, errors.New(`no "=>" in string [` + p + "]")
		}
		k := strings.TrimSpace(p[0:idx])
		v := strings.TrimSpace(p[idx+2:])
		ret[k] = v
	}
	return ret, nil
}
