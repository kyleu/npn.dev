// Code generated by hero.
// source: github.com/kyleu/npn/npntemplate/html/components/bool.html
// DO NOT EDIT!
package npntemplate

import (
	"bytes"

	"github.com/shiyanhui/hero"
)

func Boolean(b bool, trueTitle string, falseTitle string, buffer *bytes.Buffer) {
	if b {
		buffer.WriteString(`
<div class="icon" title="`)
		hero.EscapeHTML(trueTitle, buffer)
		buffer.WriteString(`" data-uk-icon="icon: check"></div>
`)
	} else {
		buffer.WriteString(`
<div class="icon" title="`)
		hero.EscapeHTML(falseTitle, buffer)
		buffer.WriteString(`" data-uk-icon="icon: close"></div>
`)
	}

}
