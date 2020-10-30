namespace request.editor {
  export function initURLEditor(el: HTMLInputElement, view: HTMLSpanElement) {

  }

  export function setURL(cache: Cache, p: Prototype | undefined) {
    if (!p) {
      dom.setValue(cache.qp, "[]");
      return;
    }
    dom.setContent(cache.urlView, prototypeToHTML(p));
    updateQueryParams(cache, p.query);
    updateBasicAuth(cache, p.auth);
  }

  export function toggleURLEditor(id: string, edit: boolean) {
    dom.setDisplay("#" + id + "-link", !edit);
    dom.setDisplay("#" + id + "-edit", edit);
    if (edit) {
      dom.req<HTMLInputElement>("#" + id + "-url").focus();
    }
  }
}
