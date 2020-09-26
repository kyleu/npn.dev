namespace request.form {
  export function getRequest(): request.Request {
    const key = gv("key");
    const title = gv("title");
    const desc = gv("description");

    let proto = request.urlToPrototype(gv("url"))
    proto.method = gv("method");
    proto.query = json.parse(gv("queryparams"));
    proto.headers = json.parse(gv("headers"));
    proto.auth = json.parse(gv("auth"));
    proto.body = json.parse(gv("body"));
    proto.options = json.parse(gv("options"));

    return {key: key, title: title, description: desc, prototype: proto};
  }

  function gv(k: string) {
    return dom.req<HTMLInputElement>(`#${request.cache.active}-${k}`).value;
  }
}
