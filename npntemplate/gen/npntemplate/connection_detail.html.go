// Code generated by hero.
// source: github.com/kyleu/npn/npntemplate/html/connection/detail.html
// DO NOT EDIT!
package npntemplate

import (
	"fmt"
	"io"

	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/shiyanhui/hero"
)

func ConnectionDetail(model *npnconnection.Status, msg *npnconnection.Message, ctx *npnweb.RequestContext, w io.Writer) (int, error) {
	_buffer := hero.GetBuffer()
	defer hero.PutBuffer(_buffer)
	_buffer.WriteString(`<!DOCTYPE html>
<html lang="en">
<head>`)
	Head(ctx, _buffer)
	_buffer.WriteString(`</head>
<body>
`)
	Navbar(ctx, _buffer)

	_buffer.WriteString(`
<div id="content" data-uk-height-viewport="expand: true">`)
	Flash(ctx, _buffer)
	_buffer.WriteString(`
<div class="uk-section uk-section-small">
  <div class="uk-container">
    <div class="uk-card uk-card-body uk-card-default">
      <h3 class="uk-card-title">Socket Connection</h3>
      <table class="uk-table uk-table-divider uk-text-left">
        <caption class="hidden">socket list</caption>
        <tbody>
        <tr>
          <th scope="row">`)
	hero.EscapeHTML(npncore.Title(npncore.KeyID), _buffer)
	_buffer.WriteString(`</th>
          <td>`)
	hero.EscapeHTML(model.ID.String(), _buffer)
	_buffer.WriteString(`</td>
        </tr>
        <tr>
          <th scope="row">User ID</th>
          <td><a href="`)
	hero.EscapeHTML(ctx.Route(npnweb.AdminLink(npncore.KeyUser, npncore.KeyDetail), `id`, model.UserID.String()), _buffer)
	_buffer.WriteString(`">`)
	hero.EscapeHTML(model.UserID.String(), _buffer)
	_buffer.WriteString(`</a></td>
        </tr>
        </tbody>
      </table>
    </div>
    <div class="uk-card uk-card-body uk-card-default mt">
      <form class="uk-form-stacked" action="`)
	hero.EscapeHTML(ctx.Route(npnweb.AdminLink(npncore.KeyConnection, `post`)), _buffer)
	_buffer.WriteString(`" method="post">
        <input name="`)
	hero.EscapeHTML(npncore.KeyID, _buffer)
	_buffer.WriteString(`" type="hidden" value="`)
	hero.EscapeHTML(fmt.Sprintf("%v", model.ID), _buffer)
	_buffer.WriteString(`" />
        <fieldset class="uk-fieldset">
          <legend class="hidden">connection form</legend>
          <div class="uk-margin-small">
            <label class="uk-form-label" for="`)
	hero.EscapeHTML(npncore.KeySvc, _buffer)
	_buffer.WriteString(`">Service</label>
            <input class="uk-input" name="`)
	hero.EscapeHTML(npncore.KeySvc, _buffer)
	_buffer.WriteString(`" type="text" value="`)
	hero.EscapeHTML(msg.Svc, _buffer)
	_buffer.WriteString(`" />
          </div>

          <div class="uk-margin">
            <label class="uk-form-label" for="cmd">Command</label>
            <input class="uk-input" name="cmd" type="text" value="`)
	hero.EscapeHTML(msg.Cmd, _buffer)
	_buffer.WriteString(`" />
          </div>

          <div class="uk-margin">
            <label class="uk-form-label" for="param">Param</label>
            `)
	Textarea(npnweb.InputParams{Name: "param", Value: npncore.ToJSON(msg.Param, ctx.Logger)}, _buffer)
	_buffer.WriteString(`
          </div>

          <div class="mt">
            <button class="uk-button uk-button-default" type="submit">Send</button>
          </div>
        </fieldset>
      </form>
    </div>
  </div>
</div>
`)

	_buffer.WriteString(`</div></body>
</html>
`)
	return w.Write(_buffer.Bytes())

}
