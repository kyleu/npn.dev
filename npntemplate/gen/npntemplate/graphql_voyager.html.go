// Code generated by hero.
// source: /Users/kyle/kyleu/npn/npntemplate/html/graphql/voyager.html
// DO NOT EDIT!
package npntemplate

import (
	"io"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/shiyanhui/hero"
)

func GraphQLVoyager(root string, ctx *npnweb.RequestContext, w io.Writer) (int, error) {
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
  <div id="`)
	hero.EscapeHTML(npncore.KeyVoyager, _buffer)
	_buffer.WriteString(`" style="height: calc(100vh - 54px);"></div>
  <script src="/vendor/react/react.production.min.js"></script>
  <script src="/vendor/react/react-dom.production.min.js"></script>

  <link rel="stylesheet" media="screen" href="/vendor/graphql-voyager/voyager.css")>
  <script src="/vendor/graphql-voyager/voyager.min.js"></script>

  <script>
    function introspectionProvider(query) {
      return fetch('`)
	hero.EscapeHTML(ctx.Route(npnweb.AdminLink(npncore.KeyGraphQL)), _buffer)
	_buffer.WriteString(`', {
        method: 'post',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({query: query}, null, 2),
      }).then(response => response.json());
    }

    GraphQLVoyager.init(document.getElementById('`)
	hero.EscapeHTML(npncore.KeyVoyager, _buffer)
	_buffer.WriteString(`'), {introspection: introspectionProvider, displayOptions: { rootType: '`)
	hero.EscapeHTML(root, _buffer)
	_buffer.WriteString(`' }})
  </script>
`)

	_buffer.WriteString(`</div></body>
</html>
`)
	return w.Write(_buffer.Bytes())

}
