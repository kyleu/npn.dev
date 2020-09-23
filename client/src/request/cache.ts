namespace request {
  class Cache {
    requests: Map<string, Request[]> = new Map();
    active?: string;
    action?: string;
    extra: string[] = [];

    setCollectionRequests(coll: collection.Collection, requests: request.Request[]) {
      this.requests.set(coll.key, requests);
      if (coll.key === collection.cache.active) {
        dom.setContent("#collection-panel", collection.renderCollection(coll, requests));
        for (let req of requests) {
          if (this.active === req.key) {
            renderActiveRequest(collection.cache.active, req);
            if (this.action) {
              renderActiveAction(collection.cache.active, req, this.action);
            }
          }
        }
      }
    }

    setActiveRequest(key: string | undefined) {
      if (!collection.cache.active) {
        console.warn("no active collection");
        return;
      }
      if (this.active !== key) {
        this.active = key;
        if (this.active) {
          const r = getActiveRequest()
          if (r) {
            renderActiveRequest(collection.cache.active, r);
          }
        } else {

        }
      }
    }

    setActiveAction(act: string | undefined, extra: string[]) {
      if (!collection.cache.active) {
        console.warn("no active collection");
        return;
      }

      if (this.action !== act) {
        this.action = act;
        const r = getActiveRequest()
        if (r) {
          renderActiveAction(collection.cache.active, r, this.action);
        }
      }
      if (this.extra.length === extra.length && this.extra.every(function(value, index) { return value === extra[index]})) {
        // same
      } else {
        this.extra = extra;
        log.info("Extra: " + this.extra);
        // TODO setActionExtra(this.action, this.extra);
      }
    }
  }

  function renderActiveRequest(coll: string, req: request.Request) {
    log.info("Request: " + req.key)
    dom.setText("#active-request-title", req.title ? req.title : req.key);
    dom.setContent("#active-request", request.form.renderForm(coll, req));
    request.editor.wireForm(req.key);
  }

  function renderActiveAction(coll: string, req: request.Request, action: string | undefined) {
    log.info("Action: " + action)
    switch (action) {
      case undefined:
        dom.setContent("#request-action", request.renderEmpty(req));
        break;
      default:
        console.warn("unhandled request action [" + action + "]")
        dom.setContent("#request-action", request.renderSplash(req));
    }
  }

  function getActiveRequest() {
    const coll = collection.cache.active;
    if (!coll) {
      return undefined
    }
    for (let req of cache.requests.get(coll) || []) {
      if (req.key === cache.active) {
        return req
      }
    }
    return undefined
  }

  export const cache = new Cache();
}
