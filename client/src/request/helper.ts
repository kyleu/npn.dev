namespace request {
  export function getSummary(coll: string, key: string) {
    for (let req of cache.summaries.get(coll) || []) {
      if (req.key === key) {
        return req
      }
    }
    return undefined
  }

  export function getRequest(coll: string, key: string) {
    for (let req of cache.requests.get(coll) || []) {
      if (req.key === key) {
        return req
      }
    }
    return undefined
  }

  export function toSummary(r: Request, order: number): Summary {
    return {key: r.key, title: r.title, description: r.description, url: prototypeToURL(r.prototype), order: order};
  }

  export function onRequestMessage(cmd: string, param: any) {
    const coll = collection.cache.active!;
    switch (cmd) {
      case command.server.requestDetail:
        const req = param as Request
        log.info("received details for request [" + req.key + "]");
        cache.updateRequest(coll, req)
        if (cache.active === req.key) {
          renderActiveRequest(coll);
          renderAction(coll, cache.active, cache.action, cache.extra);
        }
        break;
      case command.server.requestAdded:
        const ra = param as Request;
        log.info("received details for new request [" + ra.key + "]");
        cache.updateRequest(coll, ra);
        nav.navigate(`/c/${coll}/${ra.key}`);
        break;
      case command.server.requestDeleted:
        const rd = param as string;
        log.info("received details for deleted request [" + rd + "]");
        cache.removeRequest(coll, rd);
        nav.navigate(`/c/${coll}`);
        break;
      case command.server.callResult:
        const result = param as call.Result;
        call.setResult(result, location.hash);
        const path = `r/` + result.id
        // TODO history.replaceState(path, "", "/" + path);
        break;
      case command.server.transformResult:
        transform.setResult(param as transform.Result);
        break;
      default:
        console.warn(`unhandled request command [${cmd}]`);
    }
  }
}
