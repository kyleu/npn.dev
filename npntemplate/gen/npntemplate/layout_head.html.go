// Code generated by hero.
// source: /Users/kyle/kyleu/npn/npntemplate/html/layout/head.html
// DO NOT EDIT!
package npntemplate

import (
	"bytes"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/shiyanhui/hero"
)

func Head(ctx *npnweb.RequestContext, buffer *bytes.Buffer) {
	buffer.WriteString(`
<title>`)
	hero.EscapeHTML(ctx.Title, buffer)
	buffer.WriteString(`</title>
<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1.0">

<meta property="og:title" content="`)
	hero.EscapeHTML(ctx.Title, buffer)
	buffer.WriteString(`">
<meta property="og:type" content="website">
<meta property="og:locale" content="en_US">

`)
	if ctx.App.Debug() {
		buffer.WriteString(`
<link rel="stylesheet" media="screen" href="/vendor/uikit/uikit.css">
<link rel="stylesheet" media="screen" href="/assets/`)
		hero.EscapeHTML(npncore.AppKey, buffer)
		buffer.WriteString(`.css">
<script src="/vendor/uikit/uikit.js"></script>
<script src="/vendor/uikit/uikit-icons.js"></script>
<script src="/assets/`)
		hero.EscapeHTML(npncore.AppKey, buffer)
		buffer.WriteString(`.js"></script>
`)
	} else {
		buffer.WriteString(`
<link rel="stylesheet" media="screen" href="/vendor/uikit/uikit.min.css")>
<link rel="stylesheet" media="screen" href="/assets/`)
		hero.EscapeHTML(npncore.AppKey, buffer)
		buffer.WriteString(`.min.css")>
<script src="/vendor/uikit/uikit.min.js"></script>
<script src="/vendor/uikit/uikit-icons.min.js"></script>
<script src="/assets/`)
		hero.EscapeHTML(npncore.AppKey, buffer)
		buffer.WriteString(`.min.js"></script>
`)
	}

}
