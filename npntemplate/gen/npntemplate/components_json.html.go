// Code generated by hero.
// DO NOT EDIT!
package npntemplate

import (
	"bytes"

	"github.com/kyleu/npn/npncore"
	"github.com/shiyanhui/hero"
	"logur.dev/logur"
)

func JSON(t interface{}, logger logur.Logger, buffer *bytes.Buffer) {
	buffer.WriteString(`
<pre><code>`)
	hero.EscapeHTML(npncore.ToJSON(t, logger), buffer)
	buffer.WriteString(`</code></pre>
`)

}