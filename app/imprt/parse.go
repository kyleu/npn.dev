package imprt

import (
	"encoding/json"
	"strings"

	"github.com/ghodss/yaml"

	"emperror.dev/errors"
	"github.com/kyleu/libnpn/npncore"
	"github.com/kyleu/npn/app/request"
)

func parse(filename string, contentType string, content []byte) (string, interface{}, error) {
	pkey := initialPhaseKey(filename, contentType)
	p := parsePhase(filename, &phase{Key: pkey, Value: content}, 0)
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

func parsePhase(filename string, p *phase, depth int) *phase {
	if p == nil {
		return &phase{Key: "error", Error: errors.New("nil phase"), Final: true}
	}
	if depth > 32 {
		return &phase{Key: "error", Error: errors.New("recursion limit reached for phase [" + p.Key + "]"), Final: true}
	}
	var ret *phase
	switch p.Key {
	case "json", "yaml":
		ret = parseContent(filename, p.Value.([]byte))
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
	return parsePhase(filename, ret, depth+1)
}

func parseContent(filename string, content []byte) *phase {
	var obj map[string]interface{}
	err := yaml.Unmarshal(content, &obj)
	if err == nil {
		return parseObject(filename, obj, content)
	}

	var arr []json.RawMessage
	err = yaml.Unmarshal(content, &arr)
	if err == nil {
		ret := make([]*phase, 0)
		for _, e := range arr {
			ret = append(ret, parseContent(filename, e))
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

func parseObject(filename string, obj map[string]interface{}, content []byte) *phase {
	_, pok := obj["prototype"]
	_, mok := obj["method"]
	_, dok := obj["domain"]
	if pok || (mok && dok) {
		ret, err := request.FromString(npncore.KeyImport, string(content))
		if err != nil {
			return errorPhase(errors.Wrap(err, "parse-error"), obj)
		}
		return &phase{Key: "full", Value: ret, Final: true}
	}

	_, ook := obj["openapi"]
	if ook {
		return parseOpenAPI3(content, filename)
	}
	_, sok := obj["swagger"]
	if sok {
		return parseOpenAPI2(content, filename)
	}

	infoInt, iok := obj["info"]
	if iok {
		info, iok := infoInt.(map[string]interface{})
		if iok {
			sch, iok := info["schema"]
			if iok && strings.Contains(sch.(string), "postman") {
				return parsePostman(content, filename)
			}
		}
	}

	return errorPhase(errors.New("unhandled-json"), obj)
}

func parseString(s string) *phase {
	ret, err := request.FromString(npncore.KeyImport, s)
	if err == nil {
		return &phase{Key: "request", Value: ret, Final: true}
	}
	return nil
}
