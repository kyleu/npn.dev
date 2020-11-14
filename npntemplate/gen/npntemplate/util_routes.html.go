// Code generated by hero.
// source: github.com/kyleu/npn/npntemplate/html/util/routes.html
// DO NOT EDIT!
package npntemplate

import (
	"io"
	"strings"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/shiyanhui/hero"
)

func RoutesList(ctx *npnweb.RequestContext, w io.Writer) (int, error) {
	_buffer := hero.GetBuffer()
	defer hero.PutBuffer(_buffer)
	_buffer.WriteString(`<!DOCTYPE html>
<html lang="en">
<head>`)
	Head(ctx, _buffer)
	_buffer.WriteString(`</head>
<body>
`)
	Navbar(ctx, _buffer)

	_buffer.WriteString(`
<div id="content" data-uk-height-viewport="expand: true">`)
	Flash(ctx, _buffer)
	_buffer.WriteString(`
<div class="uk-section uk-section-small">
  <div class="uk-container">
    <div class="uk-card uk-card-body uk-card-default">
      <table class="uk-table uk-table-divider uk-table-small">
        <caption class="hidden">route listing</caption>
        <thead>
          <tr>
            <th scope="col">`)
	hero.EscapeHTML(npncore.Title(npncore.KeyName), _buffer)
	_buffer.WriteString(`</th>
            <th scope="col">Methods</th>
            <th scope="col">Path</th>
          </tr>
        </thead>
        <tbody>
          `)
	for _, r := range npnweb.ExtractRoutes(ctx.Routes) {
		if len(r.Methods) == 0 {
			_buffer.WriteString(`
            <tr>
              <th scope="row" colspan="3"><div class="mt">`)
			hero.EscapeHTML(strings.TrimPrefix(r.Path, "/"), _buffer)
			_buffer.WriteString(`</div></th>
            </tr>
            `)
		} else {
			_buffer.WriteString(`
            <tr>
              <td>`)
			hero.EscapeHTML(r.Name, _buffer)
			_buffer.WriteString(`</td>
              <td>`)
			hero.EscapeHTML(r.Methods, _buffer)
			_buffer.WriteString(`</td>
              <td>
              `)
			if strings.Contains(r.Methods, "GET") && len(npncore.PathParams(r.Path)) == 0 {
				_buffer.WriteString(`
                <a href="`)
				hero.EscapeHTML(r.Path, _buffer)
				_buffer.WriteString(`">`)
				hero.EscapeHTML(r.Path, _buffer)
				_buffer.WriteString(`</a>
              `)
			} else {
				hero.EscapeHTML(r.Path, _buffer)
			}
			_buffer.WriteString(`
              </td>
            </tr>
            `)
		}
	}
	_buffer.WriteString(`
        </tbody>
      </table>
    </div>
  </div>
</div>
`)

	_buffer.WriteString(`</div></body>
</html>
`)
	return w.Write(_buffer.Bytes())

}
