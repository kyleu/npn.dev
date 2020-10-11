namespace map {
  export class Map<K, V> {
    readonly storage: any = Object.create(null);

    get(key: K) {
      return this.storage[key] as V;
    };

    set(key: K, v: V) {
      return this.storage[key] = v;
    };

    del(key: K) {
      delete this.storage[key];
    };

    constructor() {}
  }
}
