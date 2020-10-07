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

    getActiveCollection() {
      for (const x of this.collections) {
        if (x.key == this.active) {
          return x;
        }
      }
      return undefined;
    }
  }

  export const cache = new Cache();
}
