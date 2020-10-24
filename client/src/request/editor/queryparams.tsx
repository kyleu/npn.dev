namespace request.editor {
  export function initQueryParamsEditor(el: HTMLTextAreaElement) {
    const parent = el.parentElement!;
    parent.appendChild(createQueryParamsEditor(el));
  }

  export function setQueryParams(el: HTMLInputElement, view: HTMLSpanElement, qp: QueryParam[] | undefined) {
    let ret: string[] = [];
    if (qp) {
      for (let p of qp) {
        let x = encodeURIComponent(p.k);
        if (p.v.length > 0) {
          x += '=' + encodeURIComponent(p.v);
        }
        ret.push(x);
      }
    }

    const orig = el.value;

    let f = "";
    const fIdx = orig.indexOf("#");
    if (fIdx > -1) {
      f = orig.substr(fIdx + 1);
    }

    const qIdx = orig.indexOf("?");
    let url = "";
    if (qIdx === -1) {
      url = orig + "?" + ret.join("&");
    } else {
      url = orig.substr(0, qIdx) + "?" + ret.join("&");
    }
    if (f.length > 0) {
      url += "#" + encodeURIComponent(f);
    }
    dom.setValue(el, url);
    dom.setContent(view, prototypeToHTML(prototypeFromURL(url)));
  }

  export function updateQueryParams(cache: Cache, qp: QueryParam[] | undefined) {
    dom.setValue(cache.qp, json.str(qp));
    updateFn(cache.qp, dom.req("#" + cache.qp.id + "-ul"))
  }

  function updateFn(el: HTMLTextAreaElement, container: HTMLElement) {
    const curr = json.parse(el.value) as QueryParam[];
    dom.clear(container);
    container.appendChild(mapHeader(el.id, "addQueryParamRow"));
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

  export function removeQueryParamRow(id: string, el: HTMLElement) {
    el.parentElement!.parentElement!.parentElement!.parentElement!.remove();
    parseQueryParams(id);
    return false;
  }

  export function addQueryParamRow(id: string) {
    const ul = dom.req("#" + id + "-ul");
    const idx = ul.children.length - 1;
    addChild(id, idx, ul, {k: '', v: ''});
    return false;
  }

  function addChild(elID: string, idx: number, container: HTMLElement, qp: QueryParam) {
    const ret = newChild(elID, idx, qp, "removeQueryParamRow");
    container.appendChild(ret);
    events(ret, () => parseQueryParams(elID))
  }

  function parseQueryParams(elID: string) {
    let ret: QueryParam[] = parseMapParams(elID);
    const ta = dom.req<HTMLTextAreaElement>("#" + elID);
    dom.setValue(ta, json.str(ret));
    const e = dom.req<HTMLInputElement>("#" + elID.replace("-queryparams", "-url"));
    const v = dom.req("#" + elID.replace("-queryparams", "-urlview"));
    setQueryParams(e, v, ret);
    check();
    return ret;
  }
}
