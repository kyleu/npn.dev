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
  }

  export const cache = new Cache();
}
