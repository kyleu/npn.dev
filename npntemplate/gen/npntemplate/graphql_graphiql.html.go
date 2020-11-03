// Code generated by hero.
// source: github.com/kyleu/npn/npntemplate/html/graphql/graphiql.html
// DO NOT EDIT!
package npntemplate

import (
	"io"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnuser"
	"github.com/kyleu/npn/npnweb"
	"github.com/shiyanhui/hero"
)

func GraphiQL(ctx *npnweb.RequestContext, w io.Writer) (int, error) {
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
  <div id="`)
	hero.EscapeHTML(npncore.KeyGraphiQL, _buffer)
	_buffer.WriteString(`" style="height: calc(100vh - 54px);"></div>
  <link rel="stylesheet" media="screen" href="/assets/vendor/graphiql/graphiql.min.css")>
  <script src="/assets/vendor/react/react.production.min.js"></script>
  <script src="/assets/vendor/react/react-dom.production.min.js"></script>
  <script src="/assets/vendor/graphiql/graphiql.min.js"></script>

  <script>
    function graphQLFetcher(graphQLParams) {
      return fetch("`)
	hero.EscapeHTML(ctx.Route(npnweb.AdminLink(npncore.KeyGraphQL)), _buffer)
	_buffer.WriteString(`", {
        method: "post",
        headers: {
          "Accept": "application/json",
          "Content-Type": "application/json",
        },
        body: JSON.stringify(graphQLParams, null, 2),
      }).then(function (response) {
        return response.text();
      }).then(function (responseBody) {
        try {
          return JSON.parse(responseBody);
        } catch (error) {
          return responseBody;
        }
      });
    }

    function voyagerLink(t, icon, route) {
      return React.createElement("a", {
        className: "theme profile-link uk-icon-link",
        "data-uk-icon": "icon: " + icon,
        href: route,
        title: "explore the " + t + " schema"
      });
    }

    window.addEventListener("load", function() {
      var q = voyagerLink("query", "settings", "`)
	hero.EscapeHTML(ctx.Route(npnweb.AdminLink(npncore.KeyVoyager, `query`)), _buffer)
	_buffer.WriteString(`");
      var s = React.createElement("span", { style: { display: "inline-block", width: "16px" } });
      var m = voyagerLink("mutation", "cog", "`)
	hero.EscapeHTML(ctx.Route(npnweb.AdminLink(npncore.KeyVoyager, `mutation`)), _buffer)
	_buffer.WriteString(`");
      ReactDOM.render(
        React.createElement(GraphiQL, {
          fetcher: graphQLFetcher,
          defaultVariableEditorOpen: true,
          editorTheme: "darcula"
        }, React.createElement(GraphiQL.Logo, {}, q, s, m)),
        document.getElementById("`)
	hero.EscapeHTML(npncore.KeyGraphiQL, _buffer)
	_buffer.WriteString(`")
      )
    }, false);
  </script>
  <style>
    .graphiql-container .execute-button-wrap {
      margin-left: 16px;
    }
  </style>
  `)
	if ctx.Profile.Theme == npnuser.ThemeDark {
		_buffer.WriteString(`
  <style>
    .cm-s-darcula  { font-family: Consolas, Menlo, Monaco, 'Lucida Console', 'Liberation Mono', 'DejaVu Sans Mono', 'Bitstream Vera Sans Mono', 'Courier New', monospace, serif;}
    .cm-s-darcula.CodeMirror { background: #2B2B2B; color: #A9B7C6; }

    .cm-s-darcula span.cm-meta { color: #BBB529; }
    .cm-s-darcula span.cm-number { color: #6897BB; }
    .cm-s-darcula span.cm-keyword { color: #CC7832; line-height: 1em; font-weight: bold; }
    .cm-s-darcula span.cm-def { color: #A9B7C6; font-style: italic; }
    .cm-s-darcula span.cm-variable { color: #A9B7C6; }
    .cm-s-darcula span.cm-variable-2 { color: #A9B7C6; }
    .cm-s-darcula span.cm-variable-3 { color: #9876AA; }
    .cm-s-darcula span.cm-type { color: #AABBCC; font-weight: bold; }
    .cm-s-darcula span.cm-property { color: #FFC66D; }
    .cm-s-darcula span.cm-operator { color: #A9B7C6; }
    .cm-s-darcula span.cm-string { color: #6A8759; }
    .cm-s-darcula span.cm-string-2 { color: #6A8759; }
    .cm-s-darcula span.cm-comment { color: #61A151; font-style: italic; }
    .cm-s-darcula span.cm-link { color: #CC7832; }
    .cm-s-darcula span.cm-atom { color: #CC7832; }
    .cm-s-darcula span.cm-error { color: #BC3F3C; }
    .cm-s-darcula span.cm-tag { color: #629755; font-weight: bold; font-style: italic; text-decoration: underline; }
    .cm-s-darcula span.cm-attribute { color: #6897bb; }
    .cm-s-darcula span.cm-qualifier { color: #6A8759; }
    .cm-s-darcula span.cm-bracket { color: #A9B7C6; }
    .cm-s-darcula span.cm-builtin { color: #FF9E59; }
    .cm-s-darcula span.cm-special { color: #FF9E59; }
    .cm-s-darcula span.cm-matchhighlight { color: #FFFFFF; background-color: rgba(50, 89, 48, .7); font-weight: normal;}
    .cm-s-darcula span.cm-searching { color: #FFFFFF; background-color: rgba(61, 115, 59, .7); font-weight: normal;}

    .cm-s-darcula .CodeMirror-cursor { border-left: 1px solid #A9B7C6; }
    .cm-s-darcula .CodeMirror-activeline-background { background: #323232; }
    .cm-s-darcula .CodeMirror-gutters { background: #313335; border-right: 1px solid #313335; }
    .cm-s-darcula .CodeMirror-guttermarker { color: #FFEE80; }
    .cm-s-darcula .CodeMirror-guttermarker-subtle { color: #D0D0D0; }
    .cm-s-darcula .CodeMirror-linenumber { color: #606366; }
    .cm-s-darcula .CodeMirror-matchingbracket { background-color: #3B514D; color: #FFEF28 !important; font-weight: bold; }

    .cm-s-darcula div.CodeMirror-selected { background: #214283; }

    .CodeMirror-hints.darcula {
      font-family: Menlo, Monaco, Consolas, 'Courier New', monospace;
      color: #9C9E9E;
      background-color: #3B3E3F !important;
    }

    .CodeMirror-hints.darcula .CodeMirror-hint-active {
      background-color: #494D4E !important;
      color: #9C9E9E !important;
    }
  </style>
  `)
	}

	_buffer.WriteString(`</div></body>
</html>
`)
	return w.Write(_buffer.Bytes())

}
