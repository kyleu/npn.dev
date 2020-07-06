namespace collection {
  export class Group<K, V> {
    readonly key: K;
    readonly members: V[] = [];

    constructor(key: K) {
      this.key = key;
    }
  }

  export class GroupSet<K, V> {
    readonly groups: Group<K, V>[] = [];

    findOrInsert(key: K): Group<K, V> {
      const ret = this.groups.find(x => x.key === key);
      if (ret) {
        return ret;
      }
      const n = new Group<K, V>(key);
      this.groups.push(n);
      return n;
    }
  }

  export function groupBy<K, V>(list: V[] | null, func: (x: V) => K): GroupSet<K, V> {
    const res = new GroupSet<K, V>();
    if (list) {
      list.forEach(o => {
        const group = res.findOrInsert(func(o));
        group.members.push(o);
      });
    }
    return res;
  }

  export function findGroup<K, V>(groups: collection.Group<K, V>[], key: K): readonly V[] {
    for (const g of groups) {
      if (g.key === key) {
        return g.members;
      }
    }
    return [];
  }

  export function flatten<T>(a: readonly T[][]): ReadonlyArray<T> {
    const ret: T[] = [];
    a.forEach(v => ret.push(...v));
    return ret;
  }
}
