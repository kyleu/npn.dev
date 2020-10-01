// Code generated by hero.
// source: /Users/kyle/kyleu/npn/npntemplate/html/components/form/tag-editor.html
// DO NOT EDIT!
package npntemplate

import (
	"bytes"
	"strings"

	"github.com/shiyanhui/hero"
)

func TagEditor(name string, opts []string, buffer *bytes.Buffer) {
	buffer.WriteString(`
<input id="model-`)
	hero.EscapeHTML(name, buffer)
	buffer.WriteString(`-input" name="`)
	hero.EscapeHTML(name, buffer)
	buffer.WriteString(`" type="hidden" value="`)
	hero.EscapeHTML(strings.Join(opts, `,`), buffer)
	buffer.WriteString(`" />
<div class="tag-editor" data-key="`)
	hero.EscapeHTML(name, buffer)
	buffer.WriteString(`" data-uk-sortable="cls-no-drag: add-item">
  `)
	for _, s := range opts {
		buffer.WriteString(`
    <span class="item">
      <span class="value" onclick="tags.editTag(this.parentElement);">`)
		hero.EscapeHTML(s, buffer)
		buffer.WriteString(`</span>
      <span class="editor"></span>
      <span class="close" data-uk-icon="icon: close; ratio: 0.6;" onclick="tags.removeTag(this);"></span>
    </span>
  `)
	}
	buffer.WriteString(`
  <span class="add-item uk-icon" data-uk-icon="icon: plus; ratio: 0.8;" onclick="tags.addTag(this);"></span>
  <div class="clear"></div>
</div>
`)

}
