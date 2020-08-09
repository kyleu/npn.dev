// Code generated by hero.
// source: /Users/kyle/kyleu/npn/npntemplate/html/util/error.html
// DO NOT EDIT!
package npntemplate

import (
	"io"
	"net/http"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/shiyanhui/hero"
)

func InternalServerError(ed *npncore.ErrorDetail, r *http.Request, ctx *npnweb.RequestContext, w io.Writer) (int, error) {
	_buffer := hero.GetBuffer()
	defer hero.PutBuffer(_buffer)
	_buffer.WriteString(`<!DOCTYPE html>
<html lang="en">
<head>
`)
	Head(ctx, _buffer)
	_buffer.WriteString(`
</head>
<body class="`)
	hero.EscapeHTML(ctx.Profile.Theme.CSS, _buffer)
	_buffer.WriteString(`">
`)
	Navbar(ctx, _buffer)

	_buffer.WriteString(`
<div id="content" data-uk-height-viewport="expand: true">
`)
	Flash(ctx, _buffer)
	_buffer.WriteString(`
<div class="uk-section uk-section-small">
  <div class="uk-container">
    <div>
      <h1 class="uk-heading-hero">`)
	hero.EscapeHTML(ed.Message, _buffer)
	_buffer.WriteString(`</h1>
      <div class="uk-text-lead">Internal Server Error</div>
      `)
	errorStack(ed, ctx, _buffer)
	_buffer.WriteString(`
    </div>
    `)
	cause := ed.Cause
	for cause != nil {
		_buffer.WriteString(`
      <div class="uk-text-lead">Caused by</div>
      <div class="uk-text-lead">`)
		hero.EscapeHTML(cause.Message, _buffer)
		_buffer.WriteString(`</div>
      `)
		errorStack(cause, ctx, _buffer)
		cause = cause.Cause
	}
	_buffer.WriteString(`
  </div>
</div>
`)

	_buffer.WriteString(`
</div>
`)
	InitDom(ctx, _buffer)
	_buffer.WriteString(`
</body>
</html>
`)
	return w.Write(_buffer.Bytes())

}
