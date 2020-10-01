// Code generated by hero.
// source: /Users/kyle/kyleu/npn/npntemplate/html/util/message.html
// DO NOT EDIT!
package npntemplate

import (
	"io"

	"github.com/kyleu/npn/npnweb"
	"github.com/shiyanhui/hero"
)

func StaticMessage(message string, ctx *npnweb.RequestContext, w io.Writer) (int, error) {
	_buffer := hero.GetBuffer()
	defer hero.PutBuffer(_buffer)
	_buffer.WriteString(`<!DOCTYPE html>
<html lang="en">
<head>`)
	Head(ctx, _buffer)
	InitDom(ctx, _buffer)
	_buffer.WriteString(`</head>
<body class="`)
	hero.EscapeHTML(ctx.Profile.Theme.CSS, _buffer)
	_buffer.WriteString(`">
`)
	Navbar(ctx, _buffer)

	_buffer.WriteString(`
<div id="content" data-uk-height-viewport="expand: true">`)
	Flash(ctx, _buffer)
	_buffer.WriteString(`
<div class="uk-section uk-section-small">
  <div class="uk-container">
    <div>
      <h1 class="uk-heading-hero">`)
	hero.EscapeHTML(message, _buffer)
	_buffer.WriteString(`</h1>
    </div>
  </div>
</div>
`)

	_buffer.WriteString(`</div></body>
</html>
`)
	return w.Write(_buffer.Bytes())

}
