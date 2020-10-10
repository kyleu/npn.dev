namespace request.editor {
  export function initQueryParamsEditor(el: HTMLTextAreaElement) {
    const parent = el.parentElement!;
    parent.appendChild(createQueryParamsEditor(el));
  }

  export function setQueryParams(el: HTMLInputElement, qp: QueryParam[] | undefined) {
    let ret: string[] = [];
    if (qp) {
      for (let p of qp) {
        ret.push(encodeURIComponent(p.k) + '=' + encodeURIComponent(p.v))
      }
    }

    const url = new URL(el.value);
    url.search = ret.join("&")
    el.value = url.toString();
  }

  export function updateQueryParams(cache: Cache, qp: QueryParam[] | undefined) {
    cache.qp.value = json.str(qp);
    updateFn(cache.qp, dom.req("#" + cache.qp.id + "-ul"))
  }

  function header(id: string) {
    return <li>
      <div data-uk-grid="">
        <div class="uk-width-1-4">Name</div>
        <div class="uk-width-1-4">Value</div>
        <div class="uk-width-1-2">
          <div class="right">
            <a class={style.linkColor} href="" onclick={"return request.editor.addHeaderRow('" + id + "')"} title="new header">
              <span data-uk-icon="icon: plus" />
            </a>
          </div>
          Description
        </div>
      </div>
    </li>;
  }

  function updateFn(el: HTMLTextAreaElement, container: HTMLElement) {
    const curr = json.parse(el.value) as header.Header[];
    container.innerText = ""
    container.appendChild(header(el.id));
    if (curr) {
      for (let idx = 0; idx < curr.length; idx++) {
        addChild(el.id, idx, container, curr[idx]);
      }
    }
  }

  function createQueryParamsEditor(el: HTMLTextAreaElement) {
    const container = <ul id={el.id + "-ul"} class="uk-list uk-list-divider" />;
    updateFn(el, container);
    return container;
  }

  export function addQueryParamRow(id: string) {
    const ul = dom.req("#" + id + "-ul");
    const idx = ul.children.length - 1;
    addChild(id, idx, ul, {k: '', v: ''});
    return false;
  }

  export function removeQueryParamRow(id: string, el: HTMLElement) {
    el.parentElement!.parentElement!.parentElement!.parentElement!.remove();
    parseQueryParams(id);
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

    events(ret, () => parseQueryParams(elID))

    container.appendChild(ret);
  }

  function parseQueryParams(elID: string) {
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
    setQueryParams(dom.req("#" + elID.replace("queryparams", "url")), ret);

    request.form.checkEditor(elID.substr(0, elID.lastIndexOf("-")));

    return ret;
  }
}
