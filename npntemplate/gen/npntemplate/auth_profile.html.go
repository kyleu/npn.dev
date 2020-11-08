// Code generated by hero.
// source: github.com/kyleu/npn/npntemplate/html/auth/profile.html
// DO NOT EDIT!
package npntemplate

import (
	"io"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnservice/auth"
	"github.com/kyleu/npn/npnuser"
	"github.com/kyleu/npn/npnweb"
	"github.com/shiyanhui/hero"
)

func Profile(auths auth.Records, ref string, ctx *npnweb.RequestContext, w io.Writer) (int, error) {
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
    <div class="uk-grid-match uk-grid-small" data-uk-grid>
      `)

	var width = "1-1"
	if ctx.App.Auth().Enabled() {
		width = "2-3"
	}

	_buffer.WriteString(`
      <div class="uk-width-`)
	hero.EscapeHTML(width, _buffer)
	_buffer.WriteString(`@m">
        <div id="profile-card" class="uk-card uk-card-body uk-card-default">
          <h3 class="uk-card-title">Profile</h3>
          <form class="uk-form-horizontal" action="`)
	hero.EscapeHTML(ctx.Route(`profile.save`), _buffer)
	_buffer.WriteString(`" method="post">
            <input type="hidden" name="ref" value="`)
	hero.EscapeHTML(ref, _buffer)
	_buffer.WriteString(`" />
            <div class="uk-margin">
              <label class="uk-form-label" for="username">Username</label>
              <div class="uk-form-controls">
                <input class="uk-input" id="username" name="username" type="text" value="`)
	hero.EscapeHTML(ctx.Profile.Name, _buffer)
	_buffer.WriteString(`" />
              </div>
            </div>

            <div class="uk-margin">
              <label class="uk-form-label">Theme</label>
              `)
	for _, t := range npnuser.AllThemes {
		_buffer.WriteString(`
              <div class="uk-form-controls uk-form-controls-text">
                <label for="theme-`)
		hero.EscapeHTML(t.String(), _buffer)
		_buffer.WriteString(`">
                  `)
		if t == ctx.Profile.Theme {
			_buffer.WriteString(`
                  <input class="uk-radio" id="theme-`)
			hero.EscapeHTML(t.String(), _buffer)
			_buffer.WriteString(`" name="`)
			hero.EscapeHTML(npncore.KeyTheme, _buffer)
			_buffer.WriteString(`" type="radio" onchange="if (this.checked) { window.setTheme(this.value); }" value="`)
			hero.EscapeHTML(t.String(), _buffer)
			_buffer.WriteString(`" checked="checked" />
                  `)
		} else {
			_buffer.WriteString(`
                  <input class="uk-radio" id="theme-`)
			hero.EscapeHTML(t.String(), _buffer)
			_buffer.WriteString(`" name="`)
			hero.EscapeHTML(npncore.KeyTheme, _buffer)
			_buffer.WriteString(`" type="radio" onchange="if (this.checked) { window.setTheme(this.value); }" value="`)
			hero.EscapeHTML(t.String(), _buffer)
			_buffer.WriteString(`" />
                  `)
		}
		hero.EscapeHTML(t.String(), _buffer)
		_buffer.WriteString(`
                </label>
              </div>
              `)
	}
	_buffer.WriteString(`
            </div>

            `)
	ProfileColors("nav", "Nav", ctx.Profile.Settings.NavColor, _buffer)
	ProfileColors("link", "Link", ctx.Profile.Settings.LinkColor, _buffer)
	_buffer.WriteString(`

            <div class="mt">
              <button class="right uk-button uk-button-default mt" type="submit">Save Changes</button>
            </div>
          </form>
        </div>
      </div>
      `)
	if ctx.App.Auth().Enabled() {
		_buffer.WriteString(`
        <div class="uk-width-1-3@m">
          `)
		AuthSigninList(auths, ctx, _buffer)
		_buffer.WriteString(`
        </div>
      `)
	}
	_buffer.WriteString(`
    </div>
  </div>
</div>
<script>
  function setNavColor(el, c) {
    document.getElementById("nav-color").value = c;
    const nb = document.getElementById("navbar");
    nb.className = ` + "`" + `${c}-bg uk-navbar-container uk-navbar` + "`" + `;
    const colors = document.querySelectorAll(".nav_swatch");
    colors.forEach(function (i) {
      i.classList.remove("active");
    });
    el.classList.add("active");
  }

  function setLinkColor(el, c) {
    document.getElementById("link-color").value = c;
    const links = document.querySelectorAll(".profile-link");
    links.forEach(l => {
      l.classList.forEach(x => {
        if (x.indexOf("-fg") > -1) {
          l.classList.remove(x);
        }
        l.classList.add(` + "`" + `${c}-fg` + "`" + `);
      });
    });
    const colors = document.querySelectorAll(".link_swatch");
    colors.forEach(function (i) {
      i.classList.remove("active");
    });
    el.classList.add("active");
  }

  function setPicture(p) {
    document.getElementById("self-picture-input").value = p;
  }
</script>
`)

	_buffer.WriteString(`</div></body>
</html>
`)
	return w.Write(_buffer.Bytes())

}
