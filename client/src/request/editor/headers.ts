namespace request.editor {
  export function initHeadersEditor(el: HTMLTextAreaElement) {
    const parent = el.parentElement!;
    parent.appendChild(createHeadersEditor(el));
  }

  export function setHeaders(cache: Cache, headers: Header[] | undefined) {

  }
}
