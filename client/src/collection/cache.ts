namespace collection {
  export interface Collection {
    readonly key: string;
    readonly title: string;
    readonly description: string;
  }

  class Cache {
    collections: Collection[] = [];
    active?: string;

    updateCollection(collection: collection.Collection) {
      this.collections = group.updateAndSort(this.collections, collection, t => t.key);
      renderCollections(this.collections);
    }

    setActiveCollection(key: string | undefined) {
      if (this.active !== key) {
        this.active = key;
        renderCollections(this.collections);
      }
    }

    deleteCollection(key: string) {
      socket.send({svc: services.collection.key, cmd: command.client.deleteCollection, param: key});
    }

    getActiveCollection() {
      return this.active ? this.getCollection(this.active) : undefined;
    }

    getCollection(key: string) {
      for (const x of this.collections) {
        if (x.key == key) {
          return x;
        }
      }
      return undefined;
    }
  }

  export const cache = new Cache();
}
