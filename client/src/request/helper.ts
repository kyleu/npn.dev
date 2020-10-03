namespace request {
  export function getActiveRequest() {
    return getRequest(collection.cache.active!, cache.active!);
  }

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

  export function onRequestMessage(cmd: string, param: any) {
    switch (cmd) {
      case command.server.requestDetail:
        const req = param as Request
        log.info("received details for request [" + req.key + "]");
        cache.updateRequest(req)
        if (cache.active === req.key) {
          renderActiveRequest(collection.cache.active!);
          renderAction(collection.cache.active!, cache.active, cache.action, cache.extra);
        }
        break;
      case command.server.callResult:
        const result = param as call.Result;
        call.setResult(result);
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
