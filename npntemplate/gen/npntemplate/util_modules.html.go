// Code generated by hero.
// source: github.com/kyleu/npn/npntemplate/html/util/modules.html
// DO NOT EDIT!
package npntemplate

import (
	"io"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/shiyanhui/hero"
)

func ModulesList(ctx *npnweb.RequestContext, w io.Writer) (int, error) {
	_buffer := hero.GetBuffer()
	defer hero.PutBuffer(_buffer)
	_buffer.WriteString(`<!DOCTYPE html>
<html lang="en">
<head>
`)
	Head(ctx, _buffer)
	_buffer.WriteString(`
<style>
  .nav-b { background-color: `)
	hero.EscapeHTML(ctx.Profile.Settings.NavB, _buffer)
	_buffer.WriteString(` !important; }
  .nav-f { color: `)
	hero.EscapeHTML(ctx.Profile.Settings.NavF, _buffer)
	_buffer.WriteString(` !important; }
  .body-b { background-color: `)
	hero.EscapeHTML(ctx.Profile.Settings.BodyB, _buffer)
	_buffer.WriteString(` !important; }
  .body-l { color: `)
	hero.EscapeHTML(ctx.Profile.Settings.BodyL, _buffer)
	_buffer.WriteString(` !important; }
</style>
</head>
<body class="body-b">
`)
	Navbar(ctx, _buffer)

	_buffer.WriteString(`
<div id="content" data-uk-height-viewport="expand: true">`)
	Flash(ctx, _buffer)
	_buffer.WriteString(`
<div class="uk-section uk-section-small">
  <div class="uk-container">
    `)
	bi := npnweb.ExtractModules()
	_buffer.WriteString(`
    <div class="uk-card uk-card-body uk-card-default">
      <table class="uk-table uk-table-divider uk-table-small">
        <caption class="hidden">module listing</caption>
        <thead>
          <tr>
            <th scope="col">`)
	hero.EscapeHTML(npncore.Title(npncore.KeyName), _buffer)
	_buffer.WriteString(`</th>
            <th scope="col">Version</th>
          </tr>
        </thead>
        <tbody>
          `)
	for _, m := range bi.Deps {
		_buffer.WriteString(`
          <tr>
            <td><a target="_blank" rel="noopener noreferrer" href="https://`)
		hero.EscapeHTML(m.Path, _buffer)
		_buffer.WriteString(`">`)
		hero.EscapeHTML(m.Path, _buffer)
		_buffer.WriteString(`</a></td>
            <td title="`)
		hero.EscapeHTML(m.Sum, _buffer)
		_buffer.WriteString(`">`)
		hero.EscapeHTML(m.Version, _buffer)
		_buffer.WriteString(`</td>
          </tr>
          `)
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
