namespace request.editor {
  export function initHeadersEditor(el: HTMLTextAreaElement) {
    const parent = el.parentElement!;
    parent.appendChild(createHeadersEditor(el));
  }

  export function setHeaders(cache: Cache, headers: header.Header[] | undefined) {

  }

  function createHeadersEditor(el: HTMLTextAreaElement) {
    const container = <ul id={el.id + "-ul"} class="uk-list uk-list-divider" />;

    const header = <li>
      <div data-uk-grid="">
        <div class="uk-width-1-4">Name</div>
        <div class="uk-width-1-4">Value</div>
        <div class="uk-width-1-2">
          <div class="right">
            <a class={style.linkColor} href="" onclick={"return request.editor.addHeaderRow('" + el.id + "')"} title="new header">
              <span data-uk-icon="icon: plus" />
            </a>
          </div>
          Description
        </div>
      </div>
    </li>;

    const curr = json.parse(el.value) as header.Header[];
    container.innerText = ""
    container.appendChild(header);
    if (curr) {
      for (let idx = 0; idx < curr.length; idx++) {
        addChild(el.id, idx, container, curr[idx]);
      }
    }

    return container;
  }

  export function addHeaderRow(id: string) {
    const ul = dom.req("#" + id + "-ul");
    const idx = ul.children.length - 1;
    addChild(id, idx, ul, {k: '', v: ''});
    return false;
  }

  export function removeHeaderRow(id: string, el: HTMLElement) {
    el.parentElement!.parentElement!.parentElement!.parentElement!.remove();
    parseHeaders(id);
    return false;
  }

  function addChild(elID: string, idx: number, container: HTMLElement, h: header.Header) {
    const ret = <li>
      <div data-uk-grid="">
        <div class="uk-width-1-4">
          <input class="uk-input" data-field={idx + "-key"} type="text" value={h.k} />
        </div>
        <div class="uk-width-1-4">
          <input class="uk-input" data-field={idx + "-value"} type="text" value={h.v} />
        </div>
        <div class="uk-width-1-2">
          <div class="right" style="margin-top: 6px;">
            <a class={style.linkColor} href="" onclick={"return request.editor.removeHeaderRow('" + elID + "', this);"} title="new header"><span data-uk-icon="icon: close" /></a>
          </div>
          <input style="width: calc(100% - 48px);" class="uk-input" data-field={idx + "-desc"} type="text" value={h.desc} />
        </div>
      </div>
    </li>

    events(ret, () => parseHeaders(elID))

    container.appendChild(ret);
  }

  function parseHeaders(elID: string) {
    const ta = dom.req<HTMLTextAreaElement>("#" + elID);
    const ul = dom.req("#" + elID + "-ul");
    const inputs = dom.els<HTMLInputElement>("input", ul);
    let ret: header.Header[] = []
    for (const i of inputs) {
      const field = i.dataset["field"] || "";
      const dash = field.lastIndexOf("-");
      const idx = parseInt(field.substring(0, dash), 10);
      const key = field.substring(dash + 1);
      if (!ret[idx]) {
        ret[idx] = {k: "", v: ""};
      }
      switch (key) {
        case "key":
          ret[idx].k = i.value.trim();
          break;
        case "value":
          ret[idx].v = i.value.trim();
          break;
        case "desc":
          const desc = i.value.trim();
          if (desc.length > 0) {
            ret[idx].desc = desc;
          }
          break;
        default:
          throw "unknown key [" + key + "]";
      }
    }

    ret = ret.filter(x => x.k.length > 0);
    ta.value = json.str(ret);

    request.form.checkEditor(elID.substr(0, elID.lastIndexOf("-")));

    return ret;
  }
}
