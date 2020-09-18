namespace collection {
  export interface Collection {
    readonly key: string;
    readonly title: string;
    readonly description: string;
  }

  class Cache {
    collections?: Collection[];
    active?: string;

    updateCollection(collection: collection.Collection) {
      // TODO
    }

    setActiveCollection(key: string | undefined) {
      this.active = key;
    }
  }

  export const cache = new Cache();
}
