namespace request.editor {
  export interface Cache {
    readonly key: HTMLInputElement;
    readonly title: HTMLInputElement;
    readonly desc: HTMLTextAreaElement;
    readonly method: HTMLSelectElement;
    readonly url: HTMLInputElement;
    readonly urlView: HTMLSpanElement;
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
      urlView: dom.req<HTMLSpanElement>(id("urlview")),
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
    initURLEditor(cache.url, cache.urlView);
    initAuthEditor(cache.auth);
    initQueryParamsEditor(cache.qp);
    initHeadersEditor(cache.headers);
    initBodyEditor(cache.body);
    initOptionsEditor(cache.options);
  }

  export function events(e: HTMLElement, f: (key: string) => void) {
    e.onchange = () => {
      f("change");
      return true;
    };
    e.onkeyup = () => {
      f("keyup");
      return true;
    };
    e.onblur = () => {
      f("blur");
      return true;
    };
  }

  export function check() {
    request.form.checkEditor(collection.cache.active!, request.cache.active!);
  }

  function wireEvents(cache: Cache) {
    events(cache.key, check);
    events(cache.title, check);
    events(cache.desc, check);
    events(cache.method, check);

    events(cache.url, function (key: string) {
      const p = prototypeFromURL(cache.url.value);
      setURL(cache, p);
      if (key == "blur") {
        request.editor.toggleURLEditor(request.cache.active!, false);
      }
      check();
    });

    events(cache.auth, function () {
      let auth: auth.Auth | undefined;
      try {
        auth = json.parse(cache.auth.value);
      } catch (e) {
        console.warn("invalid auth JSON [" + cache.auth.value + "]")
      }
      setAuth(cache.url, cache.urlView, auth);
      check();
    });

    events(cache.qp, function () {
      let qp: QueryParam[]
      try {
        qp = json.parse(cache.qp.value);
      } catch (e) {
        console.warn("invalid qp JSON [" + cache.qp.value + "]")
        qp = [];
      }
      setQueryParams(cache.url, cache.urlView, qp);
      check();
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
      check();
    });

    events(cache.body, function () {
      let b: rbody.Body | undefined
      try {
        b = json.parse(cache.body.value);
      } catch (e) {
        console.warn("invalid body JSON [" + cache.body.value + "]")
      }
      setBody(cache, b);
      check();
    });
  }
}
