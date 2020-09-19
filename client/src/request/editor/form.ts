namespace request.editor {
  export interface Cache {
    readonly url: HTMLInputElement
    readonly auth: HTMLTextAreaElement
    readonly qp: HTMLTextAreaElement
    readonly headers: HTMLTextAreaElement
    readonly body: HTMLTextAreaElement
  }

  export function wireForm(prefix: string) {
    const id = function (k: string): string {
      return "#" + prefix + "-" + k;
    }

    const cache: Cache = {
      url: dom.req<HTMLInputElement>(id("url")),
      auth: dom.req<HTMLTextAreaElement>(id("auth")),
      qp: dom.req<HTMLTextAreaElement>(id("queryparams")),
      headers: dom.req<HTMLTextAreaElement>(id("headers")),
      body: dom.req<HTMLTextAreaElement>(id("body"))
    }
    initEditors(prefix, cache);
    wireEvents(cache);
  }

  function initEditors(prefix: string, cache: Cache) {
    initURLEditor(cache.url);
    initAuthEditor(cache.auth);
    initQueryParamsEditor(cache.qp);
    initHeadersEditor(cache.headers);
    initBodyEditor(cache.body);
    initOptionsEditor(prefix);
  }

  function events(e: HTMLElement, f: () => void) {
    e.onchange = f;
    e.onkeyup = f;
    e.onblur = f;
  }

  function wireEvents(cache: Cache) {
    events(cache.url, function () {
      setURL(cache, prototypeFromURL(cache.url.value));
    });

    events(cache.auth, function () {
      let auth: auth.Auth[];
      try {
        auth = JSON.parse(cache.auth.value);
      } catch (e) {
        console.log("invalid auth JSON [" + cache.auth.value + "]")
        auth = [];
      }
      setAuth(cache, auth);
    });

    events(cache.qp, function () {
      let qp: QueryParam[]
      try {
        qp = JSON.parse(cache.qp.value);
      } catch (e) {
        console.log("invalid qp JSON [" + cache.qp.value + "]")
        qp = [];
      }
      setQueryParams(cache, qp);
    });

    events(cache.headers, function () {
      let h: Header[]
      try {
        h = JSON.parse(cache.headers.value);
      } catch (e) {
        console.log("invalid headers JSON [" + cache.headers.value + "]")
        h = [];
      }
      setHeaders(cache, h);
    });

    events(cache.body, function () {
      let b: body.Body | undefined
      try {
        b = JSON.parse(cache.body.value);
      } catch (e) {
        console.log("invalid body JSON [" + cache.body.value + "]")
      }
      setBody(cache, b);
    });
  }
}
