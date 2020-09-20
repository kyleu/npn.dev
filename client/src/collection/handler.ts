namespace collection {
  interface CollectionDetails {
    readonly collection: Collection;
    readonly requests: request.Request[];
    readonly description: string;
  }

  export function onCollectionMessage(cmd: string, param: any) {
    switch (cmd) {
      case command.server.collections:
        cache.collections = param as Collection[];
        log.info(`processing [${cache.collections.length}] collections`);
        dom.setContent("#collection-list", renderCollections(cache.collections));
        break;
      case command.server.detail:
        const d = param as CollectionDetails;
        log.info(`processing [${d.requests.length}] requests for collection [${d.collection.key}]`);
        cache.updateCollection(d.collection);
        request.cache.setCollectionRequests(d.collection.key, d.requests);
        break;
      default:
        console.warn(`unhandled collection command [${cmd}]`);
    }
  }
}