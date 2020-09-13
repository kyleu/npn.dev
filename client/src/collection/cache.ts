namespace collection {
  export interface Collection {
    readonly key: string
    readonly title: string
    readonly description: string
  }

  class Cache {
    collections?: [Collection];
  }

  export const cache = new Cache();
}
