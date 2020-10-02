// Code generated by hero.
// DO NOT EDIT!
package npntemplate

import (
	"bytes"

	"github.com/kyleu/npn/npnweb"
	"github.com/shiyanhui/hero"
)

func InitDom(ctx *npnweb.RequestContext, buffer *bytes.Buffer) {
	buffer.WriteString(`
<script>window.addEventListener("load", function() { dom.initDom('`)
	hero.EscapeHTML(ctx.Profile.Theme.String(), buffer)
	buffer.WriteString(`', '`)
	hero.EscapeHTML(ctx.Profile.LinkColor, buffer)
	buffer.WriteString(`') }, false);</script>
`)

}
