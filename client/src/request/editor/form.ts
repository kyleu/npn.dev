namespace request.editor {
  export interface Cache {
    readonly key: HTMLInputElement;
    readonly title: HTMLInputElement;
    readonly desc: HTMLTextAreaElement;
    readonly method: HTMLSelectElement;
    readonly url: HTMLInputElement;
    readonly auth: HTMLTextAreaElement;
    readonly qp: HTMLTextAreaElement;
    readonly headers: HTMLTextAreaElement;
    readonly body: HTMLTextAreaElement;
    readonly options: HTMLTextAreaElement;
  }

  export function wireForm(prefix: string) {
    const id = (k: string): string => {
      return "#" + prefix + "-" + k;
    }

    const cache: Cache = {
      key: dom.req<HTMLInputElement>(id("key")),
      title: dom.req<HTMLInputElement>(id("title")),
      desc: dom.req<HTMLTextAreaElement>(id("description")),
      url: dom.req<HTMLInputElement>(id("url")),
      method: dom.req<HTMLSelectElement>(id("method")),
      auth: dom.req<HTMLTextAreaElement>(id("auth")),
      qp: dom.req<HTMLTextAreaElement>(id("queryparams")),
      headers: dom.req<HTMLTextAreaElement>(id("headers")),
      body: dom.req<HTMLTextAreaElement>(id("body")),
      options: dom.req<HTMLTextAreaElement>(id("options"))
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
    initOptionsEditor(cache.options);
  }

  export function events(e: HTMLElement, f: () => void) {
    const x = () => {
      f();
      return true;
    }
    e.onchange = x;
    e.onkeyup = x;
    e.onblur = x;
  }

  function wireEvents(cache: Cache) {
    events(cache.key, function () {
      request.form.checkEditor(request.cache.active!);
    });
    events(cache.title, function () {
      request.form.checkEditor(request.cache.active!);
    });
    events(cache.desc, function () {
      request.form.checkEditor(request.cache.active!);
    });

    events(cache.method, function () {
      request.form.checkEditor(request.cache.active!);
    });
    events(cache.url, function () {
      const p = prototypeFromURL(cache.url.value);
      setURL(cache, p);
      request.form.checkEditor(request.cache.active!);
    });

    events(cache.auth, function () {
      let auth: auth.Auth[];
      try {
        auth = json.parse(cache.auth.value);
      } catch (e) {
        console.warn("invalid auth JSON [" + cache.auth.value + "]")
        auth = [];
      }
      setAuth(cache, auth);
    });

    events(cache.qp, function () {
      let qp: QueryParam[]
      try {
        qp = json.parse(cache.qp.value);
      } catch (e) {
        console.warn("invalid qp JSON [" + cache.qp.value + "]")
        qp = [];
      }
      setQueryParams(cache.url, qp);
    });

    events(cache.headers, function () {
      let h: header.Header[]
      try {
        h = json.parse(cache.headers.value);
      } catch (e) {
        console.warn("invalid headers JSON [" + cache.headers.value + "]")
        h = [];
      }
      setHeaders(cache, h);
    });

    events(cache.body, function () {
      let b: rbody.Body | undefined
      try {
        b = json.parse(cache.body.value);
      } catch (e) {
        console.warn("invalid body JSON [" + cache.body.value + "]")
      }
      setBody(cache, b);
    });
  }
}
