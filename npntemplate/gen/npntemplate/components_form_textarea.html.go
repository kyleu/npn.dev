// Code generated by hero.
// source: github.com/kyleu/npn/npntemplate/html/components/form/textarea.html
// DO NOT EDIT!
package npntemplate

import (
	"bytes"

	"github.com/kyleu/npn/npnweb"
	"github.com/shiyanhui/hero"
)

func Textarea(p npnweb.InputParams, buffer *bytes.Buffer) {
	buffer.WriteString(`
<div class="textarea-emoji">
  <div class="picker-toggle"><span class="uk-icon" data-uk-icon="happy"></span></div>
  <textarea id="`)
	hero.EscapeHTML(p.ID, buffer)
	buffer.WriteString(`" class="uk-textarea `)
	hero.EscapeHTML(p.Cls, buffer)
	buffer.WriteString(`" name="`)
	hero.EscapeHTML(p.Name, buffer)
	buffer.WriteString(`" placeholder="`)
	hero.EscapeHTML(p.Placeholder, buffer)
	buffer.WriteString(`">`)
	hero.EscapeHTML(p.Value, buffer)
	buffer.WriteString(`</textarea>
</div>
`)

}
