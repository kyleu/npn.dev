namespace request.form {
  export function extractRequest(reqID: string): request.Request {
    const key = gv(reqID, "key");
    const title = gv(reqID, "title");
    const desc = gv(reqID, "description");

    const url = gv(reqID, "url");
    let proto = request.urlToPrototype(url);
    proto.method = gv(reqID, "method");
    proto.query = json.parse(gv(reqID, "queryparams"));
    proto.headers = json.parse(gv(reqID, "headers"));
    proto.auth = json.parse(gv(reqID, "auth"));
    proto.body = json.parse(gv(reqID, "body"));
    proto.options = json.parse(gv(reqID, "options"));

    return {key: key, title: title, description: desc, prototype: proto};
  }

  export function checkEditor(reqID: string) {
    const o = request.getRequest(collection.cache.active!, reqID);
    let changed = false;
    if (o) {
      const n = extractRequest(reqID);
      const diff = request.diff(o, n);
      // console.log(o, n, diff);
      changed = diff.length > 0;
    } else {
      changed = true;
    }

    dom.setDisplay("#save-panel", changed)
  }

  export function saveCurrentRequest(reqID: string) {
    const req = extractRequest(reqID);
    const msg = {"coll": collection.cache.active, "orig": reqID, "req": req};
    socket.send({svc: services.request.key, cmd: command.client.saveRequest, param: msg});
  }

  function gv(r: string, k: string) {
    return dom.req<HTMLInputElement>(`#${r}-${k}`).value;
  }
}
