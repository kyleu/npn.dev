// Code generated by hero.
// source: github.com/kyleu/npn/npntemplate/html/layout/navbar-common.html
// DO NOT EDIT!
package npntemplate

import (
	"bytes"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/shiyanhui/hero"
)

func NavbarCommon(section string, showToggle bool, ctx *npnweb.RequestContext, buffer *bytes.Buffer) {
	buffer.WriteString(`
<header>
  <div data-uk-sticky="sel-target: .uk-navbar-container; cls-active: data-uk-navbar-sticky; media: 960">
    <nav id="navbar" class="uk-navbar-container `)
	hero.EscapeHTML(ctx.Profile.NavColor, buffer)
	buffer.WriteString(`-bg" data-uk-navbar>
      <div id="breadcrumbs" class="uk-navbar-left">
        `)
	if len(npnweb.IconContent) > 0 {
		buffer.WriteString(`<a title="`)
		hero.EscapeHTML(npncore.AppName, buffer)
		buffer.WriteString(`" class="uk-navbar-item uk-logo" href="`)
		hero.EscapeHTML(ctx.Route(`home`), buffer)
		buffer.WriteString(`">
          `)
		buffer.WriteString(npnweb.IconContent)
		buffer.WriteString(`
        </a>`)
	}
	BreadcrumbDisplay(ctx, buffer)
	buffer.WriteString(`
      </div>
      <div class="uk-navbar-right">
        <ul class="uk-navbar-nav">
          `)
	buffer.WriteString(npnweb.NavbarContent)
	buffer.WriteString(`<li class="uk-margin-small-right">
            `)
	if len(ctx.Profile.Picture) == 0 || ctx.Profile.Picture == "none" {
		buffer.WriteString(`<a href="`)
		hero.EscapeHTML(ctx.Route(npncore.KeyProfile), buffer)
		buffer.WriteString(`" data-uk-icon="icon:user" title="Profile"></a>`)
	} else {
		buffer.WriteString(`<a href="`)
		hero.EscapeHTML(ctx.Route(npncore.KeyProfile), buffer)
		buffer.WriteString(`" title="Profile">
              <img class="uk-border-circle" alt="user profile" src="`)
		hero.EscapeHTML(ctx.Profile.Picture, buffer)
		buffer.WriteString(`" />
            </a>`)
	}
	buffer.WriteString(`
          </li>
        </ul>
        `)
	if showToggle {
		buffer.WriteString(`<a href="" data-uk-toggle="target: #nav-offcanvas;" data-uk-navbar-toggle-icon="" class="uk-hidden@m uk-icon uk-margin-right"></a>`)
	}
	buffer.WriteString(`
      </div>
    </nav>
  </div>
</header>
`)

}
