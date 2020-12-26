package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	errors "emperror.dev/errors"
	"github.com/kyleu/npn/app/transform"
	"github.com/kyleu/libnpn/npncontroller"
	"github.com/kyleu/libnpn/npncore"
	"github.com/kyleu/libnpn/npnweb"
)

func allow(secret string, r *http.Request) bool {
	if len(secret) > 0 {
		c, _ := r.Cookie("secret")
		if c == nil || c.Value != secret {
			return false
		}
	}
	return true
}

func Enable(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		http.SetCookie(w, &http.Cookie{Name: "secret", Value: ctx.App.Secret()})
		return ctx.Route("workspace"), nil
	})
}

func Testbed(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		out, err := writeTransformerScript()
		if err != nil {
			return npncontroller.EResp(err)
		}

		_, _ = w.Write([]byte(out))
		return "", nil
	})
}

type tx struct {
	Key         string `json:"key"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func writeTransformerScript() (string, error) {
	fn := "./ui/src/util/transformers.ts"

	info, err := os.Stat(fn)
	if err != nil {
		return "", err
	}

	x, err := ioutil.ReadFile(fn)
	if err != nil {
		return "", err
	}
	fileContent := string(x)

	key := func(s string, k string) string {
		return fmt.Sprintf("  // %v of generated %v code", s, k)
	}

	replace := func(k string, content string) error {
		sKey := key("start", k)
		sIdx := strings.Index(fileContent, sKey)
		if sIdx == -1 {
			return errors.New("no index for [" + sKey + "]")
		}

		eKey := key("end", k)
		eIdx := strings.Index(fileContent, eKey)
		if eIdx == -1 {
			return errors.New("no index for [" + eKey + "]")
		}

		fileContent = fileContent[0:sIdx+len(sKey)+1] + content + "," + fileContent[eIdx-1:]

		return nil
	}

	toJSON := func(x transform.Transformer) string {
		str := npncore.ToJSONCompact(&tx{Key: x.Key(), Title: x.Title(), Description: x.Description()}, nil)
		return strings.ReplaceAll(strings.ReplaceAll(str, `","`, `", "`), `":"`, `": "`)
	}

	reqSection := []string{}
	for _, x := range transform.AllRequestTransformers {
		reqSection = append(reqSection, fmt.Sprintf("  %v", toJSON(x)))
	}
	err = replace("request", strings.Join(reqSection, ",\n"))
	if err != nil {
		return "", err
	}

	collSection := []string{}
	for _, x := range transform.AllCollectionTransformers {
		collSection = append(collSection, fmt.Sprintf("  %v", toJSON(x)))
	}
	err = replace("collection", strings.Join(collSection, ",\n"))
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(fn, []byte(fileContent), info.Mode())
	if err != nil {
		return "", err
	}

	return fileContent, nil
}
