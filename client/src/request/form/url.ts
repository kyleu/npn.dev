namespace request.form {
  export function initURLEditor(el: HTMLInputElement) {

  }

  export function setURL(cache: Cache, u: Prototype | undefined) {
    if (!u) {
      cache.qp.value = "[]"
      return;
    }
    updateQueryParams(cache, u.query);
    updateBasicAuth(cache, u.auth)
  }
}
