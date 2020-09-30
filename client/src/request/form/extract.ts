namespace request.form {
  export function extractRequest(): request.Request {
    const key = gv("key");
    const title = gv("title");
    const desc = gv("description");

    const url = gv("url");
    let proto = request.urlToPrototype(url);
    proto.method = gv("method");
    proto.query = json.parse(gv("queryparams"));
    proto.headers = json.parse(gv("headers"));
    proto.auth = json.parse(gv("auth"));
    proto.body = json.parse(gv("body"));
    proto.options = json.parse(gv("options"));

    console.log(url, proto);

    return {key: key, title: title, description: desc, prototype: proto};
  }

  function gv(k: string) {
    return dom.req<HTMLInputElement>(`#${request.cache.active}-${k}`).value;
  }
}
