namespace collection {
  interface CollectionDetails {
    readonly collection: Collection;
    readonly requests: request.Summary[];
    readonly description: string;
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
        log.info(`processing [${d.requests.length}] requests for collection [${d.collection.key}]`);
        cache.updateCollection(d.collection);
        request.cache.setCollectionRequests(d.collection, d.requests);
        renderCollections(cache.collections!);
        break;
      case command.server.collectionAdded:
        const a = param as CollectionAdded;
        log.info(`processing new collection [${a.active}]`);
        cache.collections = a.collections;
        nav.navigate("/c/" + a.active);
        break;
      default:
        console.warn(`unhandled collection command [${cmd}]`);
    }
  }
}
