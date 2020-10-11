namespace collection {
  interface CollectionDetails {
    readonly key: string;
    readonly collection?: Collection;
    readonly requests?: request.Summary[];
  }

  interface CollectionAdded {
    readonly collections: Collection[];
    readonly active: string;
  }

  export function onCollectionMessage(cmd: string, param: any) {
    switch (cmd) {
      case command.server.collections:
        cache.collections = group.sort(param as Collection[], c => c.key);
        log.info(`processing [${cache.collections.length}] collections`);
        renderCollections(cache.collections!);
        break;
      case command.server.collectionDetail:
        const d = param as CollectionDetails;
        log.info(`processing [${d.requests?.length || 0}] requests for collection [${d.key}]`);
        if (d.collection) {
          cache.updateCollection(d.collection);
        } else {
          cache.collections = cache.collections.filter(x => x.key !== del);
        }
        request.cache.setCollectionRequests(d.key, d.collection, d.requests || []);
        renderCollections(cache.collections!);
        break;
      case command.server.collectionAdded:
        const a = param as CollectionAdded;
        log.info(`processing new collection [${a.active}]`);
        cache.collections = a.collections;
        nav.navigate("/c/" + a.active);
        break;
      case command.server.collectionDeleted:
        const del = param as string;
        log.info(`processing deleted collection [${del}]`);
        cache.collections = cache.collections.filter(x => x.key !== del);
        renderCollections(cache.collections!);
        nav.navigate("/");
        break;
      default:
        console.warn(`unhandled collection command [${cmd}]`);
    }
  }
}
