// Code generated by hero.
// source: github.com/kyleu/npn/npntemplate/html/components/table-header.html
// DO NOT EDIT!
package npntemplate

import (
	"bytes"
	"net/url"

	"github.com/kyleu/npn/npncore"
	"github.com/shiyanhui/hero"
)

func TableHeader(section string, key string, title string, params *npncore.Params, req *url.URL, buffer *bytes.Buffer) {
	if params == nil {
		params = &npncore.Params{Key: section}
		buffer.WriteString(`
  <th scope="col" class="uk-transition-toggle uk-text-nowrap">
    <a class="theme" href="?`)
		hero.EscapeHTML(params.CloneOrdering(&npncore.Ordering{Column: key, Asc: true}).ToQueryString(req), buffer)
		buffer.WriteString(`">
      `)
		hero.EscapeHTML(title, buffer)
		buffer.WriteString(` <span class="uk-transition-fade" data-uk-icon="icon: chevron-down"></span>
    </a>
  </th>
`)
	} else {
		o := params.GetOrdering(key)
		if o == nil {
			buffer.WriteString(`
    <th scope="col" class="uk-transition-toggle uk-text-nowrap">
      <a class="theme" href="?`)
			hero.EscapeHTML(params.CloneOrdering(&npncore.Ordering{Column: key, Asc: true}).ToQueryString(req), buffer)
			buffer.WriteString(`">
        `)
			hero.EscapeHTML(title, buffer)
			buffer.WriteString(` <span class="uk-transition-fade" data-uk-icon="icon: chevron-down"></span>
      </a>
    </th>
  `)
		} else if o.Asc {
			buffer.WriteString(`
    <th scope="col" class=" uk-text-nowrap">
      <a class="theme" href="?`)
			hero.EscapeHTML(params.CloneOrdering(&npncore.Ordering{Column: key, Asc: false}).ToQueryString(req), buffer)
			buffer.WriteString(`">
        `)
			hero.EscapeHTML(title, buffer)
			buffer.WriteString(` <span data-uk-icon="icon: chevron-down"></span>
      </a>
    </th>
  `)
		} else {
			buffer.WriteString(`
    <th scope="col" class=" uk-text-nowrap">
      <a class="theme" href="?`)
			hero.EscapeHTML(params.CloneOrdering(&npncore.Ordering{Column: key, Asc: true}).ToQueryString(req), buffer)
			buffer.WriteString(`">
        `)
			hero.EscapeHTML(title, buffer)
			buffer.WriteString(`
        <span data-uk-icon="icon: chevron-up"></span>
      </a>
    </th>
  `)
		}
	}

}
