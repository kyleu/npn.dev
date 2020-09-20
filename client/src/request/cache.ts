namespace request {
  class Cache {
    requests: Map<string, Request[]> = new Map();
    active?: string;
    action?: string;

    setCollectionRequests(coll: string, requests: Request[]) {
      this.requests.set(coll, requests);
      if (coll === collection.cache.active) {
        dom.setContent("#request-list", view.renderRequests(coll, requests));
        for (let req of requests) {
          if (this.active === req.key) {
            renderActiveRequest(collection.cache.active, req, this.action);
          }
        }
      }
    }

    setActiveRequest(key: string | undefined) {
      if (!collection.cache.active) {
        console.warn("no active collection");
        return;
      }
      const coll = collection.cache.active;
      const reqs = this.requests.get(coll) || [];

      if (this.active !== key) {
        this.active = key;

        for (let req of reqs) {
          if (req.key === this.active) {
            renderActiveRequest(coll, req, this.action);
          }
        }
      }
    }

    setActiveAction(act: string | undefined) {
      if (!collection.cache.active) {
        console.warn("no active collection");
        return;
      }
      const coll = collection.cache.active;
      const reqs = this.requests.get(coll) || [];

      if (this.action !== act) {
        this.action = act;

        for (let req of reqs) {
          if (req.key === this.active) {
            renderActiveAction(coll, req, this.action);
          }
        }
      }
    }
  }

  function renderActiveRequest(coll: string, req: request.Request, action: string | undefined) {
    dom.setContent("#active-request", view.renderRequestDetail(coll, req));
    renderActiveAction(coll, req, action)
  }

  function renderActiveAction(coll: string, req: request.Request, action: string | undefined) {
    switch (action) {
      case undefined:
        dom.setContent("#request-action", request.renderEmpty(req));
        break;
      case "edit":
        dom.setContent("#request-action", request.form.renderForm(coll, req));
        break;
      default:
        console.warn("unhandled request action [" + action + "]")
        dom.setContent("#request-action", request.renderSplash(req));
    }
  }

  export const cache = new Cache();
}
