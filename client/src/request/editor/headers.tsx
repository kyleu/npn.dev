namespace request.editor {
  export function initHeadersEditor(el: HTMLTextAreaElement) {
    const parent = el.parentElement!;
    parent.appendChild(createHeadersEditor(el));
  }

  export function setHeaders(cache: Cache, headers: header.Header[] | undefined) {

  }

  function createHeadersEditor(el: HTMLTextAreaElement) {
    const container = <ul id={el.id + "-ul"} class="uk-list uk-list-divider" />;

    const curr = json.parse(el.value) as header.Header[];
    container.innerText = ""
    container.appendChild(mapHeader(el.id, "addHeaderRow"));
    if (curr) {
      for (let idx = 0; idx < curr.length; idx++) {
        addChild(el.id, idx, container, curr[idx]);
      }
    }

    return container;
  }

  export function removeHeaderRow(id: string, el: HTMLElement) {
    el.parentElement!.parentElement!.parentElement!.parentElement!.remove();
    parseHeaders(id);
    return false;
  }

  export function addHeaderRow(id: string) {
    const ul = dom.req("#" + id + "-ul");
    const idx = ul.children.length - 1;
    addChild(id, idx, ul, {k: '', v: ''});
    return false;
  }

  function addChild(elID: string, idx: number, container: HTMLElement, h: header.Header) {
    const ret = newChild(elID, idx, h, "removeHeaderRow")
    container.appendChild(ret);
    events(ret, () => parseHeaders(elID))
  }

  function parseHeaders(elID: string) {
    let ret: header.Header[] = parseMapParams(elID)
    const ta = dom.req<HTMLTextAreaElement>("#" + elID);
    ta.value = json.str(ret);
    check();
    return ret;
  }
}
