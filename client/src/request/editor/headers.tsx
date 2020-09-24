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
            <a class={style.linkColor} href="" onclick={"request.editor.addChild(dom.req('#" + el.id + "-ul" + "'), {k: '', v: ''});return false;"} title="new header"><span data-uk-icon="icon: plus" /></a>
          </div>
          Description
        </div>
      </div>
    </li>;

    const updateFn = function() {
      const curr = JSON.parse(el.value) as header.Header[];
      container.innerText = ""
      container.appendChild(header);
      if (curr) {
        for (let h of curr) {
          addChild(container, h);
        }
      }
    }

    updateFn();

    return container;
  }

  export function addChild(container: HTMLElement, h: header.Header) {
    container.appendChild(<li>
      <div data-uk-grid="">
        <div class="uk-width-1-4">{h.k}</div>
        <div class="uk-width-1-4">{h.v}</div>
        <div class="uk-width-1-2">
          <div class="right">
            <a class={style.linkColor} href="" onclick="return false;" title="new header"><span data-uk-icon="icon: close" /></a>
          </div>
          {h.desc ? h.desc : ""}
        </div>
      </div>
    </li>)
  }
}
