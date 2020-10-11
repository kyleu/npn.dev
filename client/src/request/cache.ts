namespace request {
  class Cache {
    summaries: map.Map<string, Summary[]> = new map.Map();
    requests: map.Map<string, Request[]> = new map.Map();
    active?: string;
    action?: string;
    extra: string[] = [];

    setCollectionRequests(key: string, coll: collection.Collection | undefined, summs: request.Summary[]) {
      if (coll) {
        this.summaries.set(key, summs);
        if (key === collection.cache.active) {
          dom.setContent("#collection-panel", collection.renderCollection(coll, summs));
          for (let req of summs) {
            if (this.active === req.key) {
              renderActiveRequest(key);
              if (this.action) {
                renderAction(key, req.key, this.action, this.extra);
              }
            }
          }
        }
      } else {
        this.summaries.del(key);
        if (key === collection.cache.active) {
          dom.setContent("#collection-panel", collection.renderNotFound(key));
        }
      }
    }

    setActiveRequest(coll: string | undefined, key: string | undefined) {
      if (!coll) {
        return;
      }
      if (this.active !== key) {
        this.active = key;
        if (this.active) {
          renderActiveRequest(coll);
        }
        collection.renderCollections(collection.cache.collections!);
      }
    }

    setActiveAction(coll: string | undefined, act: string | undefined, extra: string[]) {
      if (!coll) {
        return;
      }

      const sameExtra = this.extra.length === extra.length && this.extra.every(function(value, index) { return value === extra[index]});
      if (this.active /* && (this.action !== act || !sameExtra) */) {
        this.action = act;
        this.extra = extra;
        renderAction(coll, this.active, this.action, this.extra);
      }
    }

    updateRequest(coll: string, r: request.Request) {
      const curr = this.requests.get(coll);
      const updated = group.update(curr, r, x => x.key);
      this.requests.set(coll, updated);
      let summs = this.summaries.get(coll);
      summs = summs.map(x => x.key == r.key ? toSummary(r, 0) : x);
      this.summaries.set(coll, summs)
      if(collection.cache.active === coll) {
        collection.renderCollection(collection.cache.getActiveCollection()!, summs)
      }
    }

    removeRequest(coll: string, rd: string) {
      const curr = this.requests.get(coll);
      const updated = group.remove(curr, rd, x => x.key);
      this.requests.set(coll, updated);
      let summs = this.summaries.get(coll);
      summs = summs.filter(x => x.key !== rd);
      this.summaries.set(coll, summs)
      if (this.active === rd) {
        cache.setActiveRequest(coll, undefined);
      }
      if(collection.cache.active === coll) {
        collection.renderCollection(collection.cache.getActiveCollection()!, summs)
      }
    }
  }

  export const cache = new Cache();
}
