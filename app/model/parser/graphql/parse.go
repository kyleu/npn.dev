package parsegraphql

import (
	"io/ioutil"
	"strings"

	parseutil "github.com/kyleu/npn/app/model/parser/util"

	"emperror.dev/errors"
	"github.com/kyleu/npn/app/model/schema"
)

func (p *GraphQLParser) Parse(paths []string) (*parseutil.ParseResponse, error) {
	md := schema.Metadata{Comments: nil, Origin: schema.OriginGraphQL, Source: paths[0]}
	return p.parse(paths, parseutil.NewParseResponse(paths, md))
}

func (p *GraphQLParser) parse(paths []string, ret *parseutil.ParseResponse) (*parseutil.ParseResponse, error) {
	rsp := ret
	var err error
	for _, path := range paths {
		rsp, err = p.parsePath(path, rsp)
		if err != nil {
			return nil, errors.Wrap(err, "error parsing graphql")
		}
	}
	return rsp, nil
}

func (p *GraphQLParser) parsePath(path string, ret *parseutil.ParseResponse) (*parseutil.ParseResponse, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	schStr := strings.TrimSpace(string(b))
	if len(schStr) == 0 {
		return nil, errors.New("empty file")
	}

	return p.parseSchema(path, schStr, ret)
}
