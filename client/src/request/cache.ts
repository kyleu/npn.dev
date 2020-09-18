namespace request {
  class Cache {
    requests: Map<string, Request[]> = new Map();
    active?: string;

    setCollectionRequests(key: string, requests: Request[]) {
      this.requests.set(key, requests);
      if (key === collection.cache.active) {
        dom.setContent("#request-list", renderRequests(key, requests));
        for (let req of requests) {
          if (this.active == req.key) {
            renderActiveRequest(key, req);
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
      this.active = key;
      for (let req of reqs) {
        if (req.key == key) {
          renderActiveRequest(coll, req);
        }
      }
    }
  }

  function renderActiveRequest(key: string, req: request.Request) {
    dom.setContent("#active-request", renderRequest(key, req));
  }

  export const cache = new Cache();
}
