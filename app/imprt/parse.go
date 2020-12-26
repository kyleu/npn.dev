package imprt

import (
	"encoding/json"
	"strings"

	"github.com/kyleu/npn/app/transform"

	"emperror.dev/errors"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/libnpn/npncore"
)

func parse(filename string, contentType string, content []byte) (string, interface{}, error) {
	pkey := initialPhaseKey(filename, contentType)
	p := parsePhase(&phase{Key: pkey, Value: content}, 0)
	if p == nil {
		return "unhandled", content, errors.New("nil phase for [" + contentType + "]")
	}
	return p.Key, p.Value, p.Error
}

func initialPhaseKey(filename string, contentType string) string {
	switch {
	case contentType == "application/json" || strings.HasSuffix(filename, ".json"):
		return "json"
	case contentType == "application/x-yaml" || strings.HasSuffix(filename, ".yaml"):
		return "yaml"
	default:
		return "bytes"
	}
}

func parsePhase(p *phase, depth int) *phase {
	if p == nil {
		return &phase{Key: "error", Error: errors.New("nil phase"), Final: true}
	}
	if depth > 32 {
		return &phase{Key: "error", Error: errors.New("recursion limit reached for phase [" + p.Key + "]"), Final: true}
	}
	var ret *phase
	switch p.Key {
	case "json":
		ret = parseJSON(p.Value.([]byte))
	case "yaml":
		ret = parseYAML(p.Value.([]byte))
	case "string":
		ret = parseString(p.Value.(string))
	}
	if ret == nil {
		return errorPhase(errors.New("import for ["+p.Key+"] didn't return a phase"), p.Value)
	}
	if ret.Error != nil {
		return ret
	}
	if ret.Final {
		return ret
	}
	return parsePhase(ret, depth+1)
}

func parseJSON(content []byte) *phase {
	var obj map[string]interface{}
	err := npncore.FromJSON(content, &obj)
	if err == nil {
		return parseJSONObject(obj, content)
	}

	var arr []json.RawMessage
	err = npncore.FromJSON(content, &arr)
	if err == nil {
		ret := make([]*phase, 0)
		for _, e := range arr {
			ret = append(ret, parseJSON(e))
		}
		return &phase{Key: "set", Value: ret, Final: true}
	}

	str, err := npncore.FromJSONString(content)
	if err == nil {
		return &phase{Key: "string", Value: str}
	}

	s := string(content)
	if strings.HasPrefix(s, "http:") || strings.HasPrefix(s, "https:") || strings.HasPrefix(s, "ws:") || strings.HasPrefix(s, "wss:") {
		return &phase{Key: "string", Value: s}
	}

	return errorPhase(errors.New("unhandled JSON"), string(content))
}

func parseYAML(content []byte) *phase {
	c := string(content)
	println(c)
	if strings.Contains(c, "openapi") {
		oapi, err := transform.OpenAPIImport(content)
		if err != nil {
			return errorPhase(errors.Wrap(err, "unhandled OpenAPI error"), c)
		}
		coll, err := transform.OpenAPIToFullCollection(oapi)
		if err != nil {
			return errorPhase(errors.Wrap(err, "error transforming OpenAPI"), c)
		}
		return &phase{Key: "request", Value: coll, Final: true}
	}
	return errorPhase(errors.New("unhandled YAML"), c)
}

func parseJSONObject(obj map[string]interface{}, content []byte) *phase {
	_, pok := obj["prototype"]
	_, mok := obj["method"]
	_, dok := obj["domain"]
	if pok || (mok && dok) {
		ret, err := request.FromString("import", string(content))
		if err == nil {
			return &phase{Key: "full", Value: ret, Final: true}
		}
	}
	return &phase{Key: "unhandled JSON object", Value: obj, Final: true}
}

func parseString(s string) *phase {
	ret, err := request.FromString("import", s)
	if err == nil {
		return &phase{Key: "request", Value: ret, Final: true}
	}
	return nil
}
