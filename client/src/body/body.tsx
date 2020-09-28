namespace rbody {
  export function renderBody(requestKey: string, b?: rbody.Body) {
    if (!b) {
      return <div>No body</div>;
    }
    switch (b.type) {
      case "json":
        return renderJSON(b.config as rbody.JSONConfig);
      case "html":
        const req = request.getRequest(collection.cache.active!, requestKey);
        const baseURL = request.prototypeBaseURL(req?.prototype);
        console.log(baseURL);
        return renderHTML(b.config as rbody.HTMLConfig, baseURL);
      default:
        return <div>TODO: {b.type}</div>;
    }
  }

  function renderHTML(h: rbody.HTMLConfig, baseURL: string) {
    return <div class="html-body">
      <span class="base-url hidden">{baseURL}</span>
      <span class="preview-link right">(<a class={style.linkColor} href="" onclick="rbody.renderHTMLPreview(this);return false">preview</a>)</span>
      <span class="text-link right hidden">(<a class={style.linkColor} href="" onclick="rbody.renderHTMLText(this);return false">text</a>)</span>
      <em>HTML</em>
      <pre class="text-content" style="overflow: auto; max-height: 720px;">{h.content}</pre>
      <div class="preview-content hidden" style="overflow: auto; max-height: 720px;" />
    </div>;
  }

  function renderJSON(j: rbody.JSONConfig) {
    return <div><em>JSON</em><pre>{json.str(j.msg)}</pre></div>;
  }

  export function renderHTMLPreview(el: HTMLAnchorElement) {
    const container = editorContent(el, true);
    const iframe = document.createElement("iframe");
    iframe.style.width = "100%";
    iframe.style.minHeight = "720px";

    const html = previewHTMLFor(container[1], container[0]);

    // iframe.src = "data:text/html;charset=utf-8," + encodeURI(html);
    container[2].innerHTML = "";
    container[2].appendChild(iframe)

    iframe.contentWindow!.document.open();
    iframe.contentWindow!.document.write(html);
    iframe.contentWindow!.document.close();
  }

  export function renderHTMLText(el: HTMLAnchorElement) {
    editorContent(el, false);
  }

  function editorContent(el: HTMLAnchorElement, preview: boolean): [string, HTMLPreElement, HTMLElement] {
    const container = el.parentElement!.parentElement!
    if (!container.classList.contains("html-body")) {
      throw "container is not class [html-body]";
    }

    const baseURLEl = dom.req(".base-url", container);
    const tLink = dom.req(".text-link", container);
    const tContent = dom.req<HTMLPreElement>(".text-content", container);
    const pLink = dom.req(".preview-link", container);
    const pContent = dom.req(".preview-content", container);

    dom.setDisplay(tLink, preview);
    dom.setDisplay(tContent, !preview);
    dom.setDisplay(pLink, !preview);
    dom.setDisplay(pContent, preview);

    return [baseURLEl.innerText, tContent, pContent];
  }

  function previewHTMLFor(e: HTMLPreElement, baseURL: string) {
    let ret = e.innerText;
    const headIdx = ret.indexOf("<head")
    if (headIdx > -1) {
      const headEnd = ret.indexOf(">", headIdx);
      if (headEnd > -1) {
        const base = `<base href="${baseURL}" target="_blank">`;
        ret = ret.substr(0, headEnd + 1) + base + ret.substr(headEnd + 1);
      }
    }
    return ret;
  }
}
