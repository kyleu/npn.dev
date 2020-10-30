declare const Prism: any;

namespace rbody {
  export function renderBody(url: string, b?: rbody.Body) {
    if (!b) {
      return <div>No body</div>;
    }
    switch (b.type) {
      case "json":
        return renderJSON(b.config as rbody.JSONConfig);
      case "html":
        const baseURL = request.baseURL(url);
        return renderHTML(b.config as rbody.HTMLConfig, baseURL);
      case "image":
        return renderImage(b.config as { readonly type: string, readonly content: string });
      case "raw":
        return renderRaw(b.config as rbody.RawConfig);
      case "error":
        return renderError(b.config as { readonly message: string });
      default:
        return <div>unhandled body type [{b.type}]</div>;
    }
  }

  function renderHTML(h: rbody.HTMLConfig, baseURL: string) {
    let html: string = Prism.highlight(h.content, Prism.languages.html, 'html');
    return <div class="html-body">
      <span class="base-url hidden">{baseURL}</span>
      <span class="preview-link right">(<a class={style.linkColor} href="" onclick="rbody.renderHTMLPreview(this);return false">preview</a>)</span>
      <span class="text-link right hidden">(<a class={style.linkColor} href="" onclick="rbody.renderHTMLText(this);return false">text</a>)</span>
      <em>HTML</em>
      <div class="text-content prism mt" style="overflow: auto; max-height: 720px;">
        <pre class="language-html" data-start="0" style="white-space:pre-wrap;"><code dangerouslySetInnerHTML={{__html: html}} /></pre>
      </div>
      <div class="preview-content mt hidden" style="overflow: auto; max-height: 720px; border: 1px solid #666;" />
    </div>;
  }

  function renderJSON(j: rbody.JSONConfig) {
    let html: string = Prism.highlight(JSON.stringify(j.msg, null, 2), Prism.languages.javascript, 'javascript');
    return <div>
      <em>JSON</em>
      <div class="mt">
        <pre class="language-html" data-start="0" style="white-space:pre-wrap;"><code dangerouslySetInnerHTML={{__html: html}} /></pre>
      </div>
    </div>;
  }

  function renderImage(i: { readonly type: string; readonly content: string }) {
    const dataURL = `data:${i.type};base64,${i.content}`;
    return <img alt="response image" src={dataURL} />;
  }

  function renderRaw(r: rbody.RawConfig) {
    return <div><em>{r.type ? r.type : "Unknown Type"}</em><pre>{json.str(r)}</pre></div>;
  }

  function renderError(err: { readonly message: string }) {
    return <div><em>Error</em><pre>{err.message}</pre></div>;
  }
}
