package sandbox

import (
	"fmt"
	"time"

	"github.com/kyleu/npn/app/project"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"

	"emperror.dev/errors"
	"github.com/kyleu/npn/app/bootstrap"
	"golang.org/x/text/language"
)

var Bootstrap = Register(&Sandbox{
	Key:         "bootstrap",
	Title:       "Bootstrap",
	Description: "Packages the bootstrap projects for release",
	DevOnly:     true,
	Resolve: func(ctx *npnweb.RequestContext) (string, interface{}, error) {
		err := bootstrap.PersistAll()
		if err != nil {
			return npncore.KeyError, "persist-error", err
		}

		var ret []string
		for _, proto := range bootstrap.AllPrototypes {
			p := &project.Project{
				Key:         proto.Key,
				Title:       proto.Key,
				RootPath:    "./_projects/" + proto.Key,
				Description: proto.Description,
				Options:     map[string]interface{}{},
			}

			startNanos := time.Now().UnixNano()
			err = bootstrap.Extract(proto, p, ctx.Logger)
			if err != nil {
				return npncore.KeyError, "extract-error", errors.Wrap(err, "can't extract ["+proto.Key+"]")
			}
			err = bootstrap.Verify(proto, p, ctx.Logger)
			if err != nil {
				return npncore.KeyError, "verify-error", errors.Wrap(err, "can't verify ["+proto.Key+"]")
			}
			delta := (time.Now().UnixNano() - startNanos) / int64(time.Microsecond)
			ms := npncore.MicrosToMillis(language.AmericanEnglish, int(delta))
			ret = append(ret, fmt.Sprintf("extracted and verified [%v] in [%v]", proto.Key, ms))
		}

		return "OK", ret, nil
	},
})
