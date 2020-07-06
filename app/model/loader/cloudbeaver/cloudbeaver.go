package cloudbeaver

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"emperror.dev/errors"
	"github.com/kyleu/npn/app/model/schema"
	"github.com/kyleu/npn/app/util"
	"logur.dev/logur"
)

type Loader struct {
	url    string
	logger logur.Logger
}

type navNode struct {
	ID       string     `json:"id"`
	Type     string     `json:"type"`
	Children []*navNode `json:"children,omitempty"`
}

func NewLoader(secure bool, host string, port int, logger logur.Logger) *Loader {
	protocol := "http"
	if secure {
		protocol = "https"
	}
	u := fmt.Sprintf("%v://%v:%v/dbeaver/gql", protocol, host, port)
	return &Loader{url: u, logger: logger}
}

func (l *Loader) Detect() (map[string]string, error) {
	_, err := l.call(gqlOpenSession, nil)
	if err != nil {
		return nil, err
	}

	return l.loadRoot()
}

func (l *Loader) Crawl() (*schema.Schema, interface{}, error) {
	_, err := l.call(gqlOpenSession, nil)
	if err != nil {
		return nil, nil, err
	}

	md := &schema.Metadata{Description: "loaded from cloudbeaver", Origin: schema.OriginDatabase, Source: l.url}
	sch := schema.NewSchema("loading", []string{l.url}, md)

	tgt, err := l.load(sch, "Root", "/", []string{})
	if err != nil {
		return nil, nil, err
	}

	return sch, tgt, nil
}

func (l *Loader) call(query string, variables interface{}) (interface{}, error) {
	client := http.Client{}
	reqBody := util.ToJSON(map[string]interface{}{"query": query, "variables": variables}, l.logger)
	req, err := http.NewRequest(http.MethodPost, l.url, strings.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.AddCookie(&http.Cookie{Name: "DBEAVER_SESSION_ID", Value: "node01q27oq6n5gl8z5odzhke5zsdk0.node0"})
	rsp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "can't post GraphQL")
	}
	if rsp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("GraphQL response returned status [%v]", rsp.StatusCode))
	}

	defer func() { _ = rsp.Body.Close() }()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "can't read GraphQL response")
	}

	tgt := make(map[string]interface{})
	util.FromJSON(body, &tgt, l.logger)

	data, ok := tgt["data"]
	if ok {
		return data, nil
	}

	return tgt, nil
}
