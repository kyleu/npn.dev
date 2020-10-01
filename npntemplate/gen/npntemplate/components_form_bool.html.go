// Code generated by hero.
// source: /Users/kyle/kyleu/npn/npntemplate/html/components/form/bool.html
// DO NOT EDIT!
package npntemplate

import (
	"bytes"

	"github.com/shiyanhui/hero"
)

func InputBool(name string, selected bool, buffer *bytes.Buffer) {
	if selected {
		buffer.WriteString(`
  <label><input class="uk-radio" type="radio" name="`)
		hero.EscapeHTML(name, buffer)
		buffer.WriteString(`" value="true" checked /> True</label>
  <label><input class="uk-radio" type="radio" name="`)
		hero.EscapeHTML(name, buffer)
		buffer.WriteString(`" value="false" /> False</label>
`)
	} else {
		buffer.WriteString(`
  <label><input class="uk-radio" type="radio" name="`)
		hero.EscapeHTML(name, buffer)
		buffer.WriteString(`" value="true" /> True</label>
  <label><input class="uk-radio" type="radio" name="`)
		hero.EscapeHTML(name, buffer)
		buffer.WriteString(`" value="false" checked /> False</label>
`)
	}

}
