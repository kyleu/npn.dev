// Code generated by hero.
// DO NOT EDIT!
package npntemplate

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/shiyanhui/hero"
)

func FileBrowse(dir []string, files []os.FileInfo, ctx *npnweb.RequestContext, w io.Writer) (int, error) {
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
    <div class="uk-card uk-card-body uk-card-default">
      <h3 class="uk-card-title"><a class="theme" href="`)
	hero.EscapeHTML(ctx.Route(npncore.KeyFile), _buffer)
	_buffer.WriteString(`">/</a>`)
	hero.EscapeHTML(strings.Join(dir, "/"), _buffer)
	_buffer.WriteString(`</h3>
      <table class="uk-table uk-table-divider">
        <caption class="hidden">file listing</caption>
        <thead>
          <tr>
            <th scope="col">Name</th>
            <th scope="col">Size</th>
          </tr>
        </thead>
        <tbody>
          `)
	for _, file := range files {
		_buffer.WriteString(`
            <tr>
              <td>
                <a class="theme" href="`)
		hero.EscapeHTML(ctx.Route(npncore.KeyFile), _buffer)
		hero.EscapeHTML(strings.Join(append(dir, file.Name()), `/`), _buffer)
		_buffer.WriteString(`">
                  `)

		icon := "file-text"
		if file.IsDir() {
			icon = "folder"
		}

		_buffer.WriteString(`<span data-uk-icon="icon: `)
		hero.EscapeHTML(icon, _buffer)
		_buffer.WriteString(`"c class="uk-margin-small-right"></span>`)
		hero.EscapeHTML(file.Name(), _buffer)
		_buffer.WriteString(`
                </a>
              </td>
              <td class="uk-table-shrink uk-text-nowrap">
                `)
		if !file.IsDir() {
			hero.EscapeHTML(fmt.Sprintf("%v", file.Size()), _buffer)
		}
		_buffer.WriteString(`
              </td>
            </tr>
          `)
	}
	_buffer.WriteString(`
        </tbody>
      </table>
    </div>
  </div>
</div>
`)

	_buffer.WriteString(`</div></body>
</html>
`)
	return w.Write(_buffer.Bytes())

}