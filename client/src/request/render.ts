namespace request {
  export function renderActiveRequest(coll: string) {
    if (cache.active) {
      const req = getRequest(coll, cache.active);
      if (req) {
        dom.setContent("#request-panel", request.form.renderFormPanel(coll, req));
        request.editor.wireForm(req.key);
      } else {
        const summ = getSummary(coll, cache.active);
        if (summ) {
          dom.setContent("#request-panel", request.renderSummaryPanel(coll, summ));
          const param = {coll: coll, req: summ.key};
          socket.send({svc: services.request.key, cmd: command.client.getRequest, param: param});
        }
      }
    } else {
      console.warn("no active request")
    }
  }

  export function renderAction(coll: string, reqKey: string, action: string | undefined, extra: string[]) {
    // TODO const req = request.form.getRequest();
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
        const req = getRequest(coll, reqKey);
        if (!req) {
          return;
        }
        call.prepare(coll, request.form.extractRequest());
        // call.prepare(coll, req);
        dom.setContent(ra, renderActionCall(coll, reqKey));
        break;
      case "transform":
        dom.setContent(ra, renderActionTransform(coll, reqKey, extra[0]));
        break;
      default:
        console.warn("unhandled request action [" + action + "]");
        dom.setContent(ra, request.renderActionUnknown(action, extra));
    }
    dom.setDisplay(re, action === undefined);
    dom.setDisplay(ra, action !== undefined);
  }
}

