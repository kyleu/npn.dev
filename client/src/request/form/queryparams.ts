namespace request.form {
  export function initQueryParamsEditor(el: HTMLTextAreaElement) {

  }

  export function setQueryParams(cache: Cache, qp: QueryParam[] | undefined) {
    let ret: string[] = [];
    if (qp) {
      for (let p of qp) {
        ret.push(encodeURIComponent(p.k) + '=' + encodeURIComponent(p.v))
      }
    }

    const url = new URL(cache.url.value);
    url.search = ret.join("&")
    cache.url.value = url.toString();
  }

  export function updateQueryParams(cache: Cache, qp: QueryParam[] | undefined) {
    cache.qp.value = JSON.stringify(qp, null, 2);
  }
}
