namespace request {
  export function renderActiveRequest(coll: string) {
    if (cache.active) {
      render(coll, cache.active);
    } else {
      console.warn("no active request")
    }
  }

  export function render(coll: string, reqKey: string) {
    const req = getRequest(coll, reqKey);
    if (req) {
      dom.setContent("#request-panel", request.form.renderFormPanel(coll, req));
      request.editor.wireForm(req.key);
    } else {
      const summ = getSummary(coll, reqKey);
      if (summ) {
        dom.setContent("#request-panel", request.renderSummaryPanel(coll, summ));
        const param = {coll: coll, req: summ.key};
        socket.send({svc: services.request.key, cmd: command.client.getRequest, param: param});
      }
    }
  }

  export function renderAction(coll: string, reqKey: string, action: string | undefined, extra: string[]) {
    const re = dom.opt(".request-editor");
    const ra = dom.opt(".request-action");
    if(!re || !ra) {
      return;
    }
    switch (action) {
      case undefined:
        dom.setContent(ra, renderActionEmpty());
        break;
      case "call":
        // call.prepare(coll, getRequest(coll, reqKey));
        call.prepare(coll, request.form.extractRequest(request.cache.active!));
        dom.setContent(ra, renderActionCall(coll, reqKey));
        break;
      case "delete":
        const del = {coll: coll, req: reqKey};
        socket.send({svc: services.request.key, cmd: command.client.deleteRequest, param: del});
        break;
      case "transform":
        const req = request.form.extractRequest(request.cache.active!)
        dom.setContent(ra, transform.renderRequest(coll, reqKey, extra[0]));
        const tx = {coll: coll, req: reqKey, fmt: extra[0], proto: req.prototype};
        socket.send({svc: services.request.key, cmd: command.client.transform, param: tx});
        break;
      default:
        console.warn("unhandled request action [" + action + "]");
        dom.setContent(ra, request.renderActionUnknown(action, extra));
    }
    dom.setDisplay(re, action === undefined);
    dom.setDisplay(ra, action !== undefined);
  }
}

