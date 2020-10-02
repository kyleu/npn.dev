namespace request {
  class Cache {
    summaries: map.Map<string, Summary[]> = new map.Map();
    requests: map.Map<string, Request[]> = new map.Map();
    active?: string;
    action?: string;
    extra: string[] = [];

    setCollectionRequests(coll: collection.Collection, summs: request.Summary[]) {
      this.summaries.set(coll.key, summs);
      if (coll.key === collection.cache.active) {
        dom.setContent("#collection-panel", collection.renderCollection(coll, summs));
        for (let req of summs) {
          if (this.active === req.key) {
            renderActiveRequest(collection.cache.active);
            if (this.action) {
              renderAction(collection.cache.active, req.key, this.action, this.extra);
            }
          }
        }
      }
    }

    setActiveRequest(key: string | undefined) {
      if (!collection.cache.active) {
        return;
      }
      if (this.active !== key) {
        this.active = key;
        if (this.active) {
          renderActiveRequest(collection.cache.active);
        }
        collection.renderCollections(collection.cache.collections!);
      }
    }

    setActiveAction(act: string | undefined, extra: string[]) {
      if (!collection.cache.active) {
        return;
      }

      const sameExtra = this.extra.length === extra.length && this.extra.every(function(value, index) { return value === extra[index]});
      if (this.active && (this.action !== act || !sameExtra)) {
        this.action = act;
        this.extra = extra;
        renderAction(collection.cache.active, this.active, this.action, this.extra);
      }
    }

    updateRequest(r: request.Request) {
      if (!collection.cache.active) {
        return;
      }
      const curr = this.requests.get(collection.cache.active);
      const updated = group.update(curr, r, x => x.key);
      this.requests.set(collection.cache.active, updated);
    }
  }

  export const cache = new Cache();
}
