namespace rbody {
  export function renderHTMLPreview(el: HTMLAnchorElement) {
    const container = editorContent(el, true);
    const iframe = document.createElement("iframe");
    iframe.style.width = "100%";
    iframe.style.minHeight = "720px";

    const html = previewHTMLFor(container[1], container[0]);

    // iframe.src = "data:text/html;charset=utf-8," + encodeURI(html);
    container[2].innerHTML = "";
    container[2].appendChild(iframe)

    const idoc = iframe.contentDocument || iframe.contentWindow!.document;

    idoc.open();
    idoc.write(html);
    idoc.close();
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
    const headIdx = ret.indexOf("<head");
    if (headIdx > -1) {
      const headEnd = ret.indexOf(">", headIdx);
      if (headEnd > -1) {
        const base = `<base href="${baseURL}" target="_blank">`;
        ret = ret.substr(0, headEnd + 1) + base + ret.substr(headEnd + 1);
      }
    }
    return ret;
  }

  export function iframeError() {
    return <div class="uk-container">
      <div class="uk-section uk-section-small">
        <h3>Rendering Error</h3>
        <p>This page indicates that the HTML preview was unable to render.</p>
      </div>
    </div>;
  }
}
