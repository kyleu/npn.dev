namespace group {
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
      const ret = arr.find(this.groups, (_, x) => x.key === key);
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

  export function findGroup<K, V>(groups: Group<K, V>[], key: K): readonly V[] {
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

  export function sort<T, S>(a: T[] | undefined, matchFn: (t: T) => S): T[] {
    if (!a) {
      return [];
    }
    a.sort((l, r) => {
      const lv = matchFn(l);
      const rv = matchFn(r);
      if (lv > rv) {
        return 1;
      }
      if (lv < rv) {
        return -1;
      }
      return 0;
    })
    return a;
  }

  export function update<T, S>(a: T[] | undefined, v: T, matchFn: (t: T) => S): T[] {
    if (!a) {
      return [v];
    }
    let matched = false;
    const key = matchFn(v);
    for (const idx in a) {
      const c = a[idx]
      if (matchFn(c) == key) {
        matched = true;
        a[idx] = v;
      }
    }
    if (!matched) {
      a.push(v);
    }
    return a;
  }

  export function updateAndSort<T, S>(a: T[] | undefined, v: T, matchFn: (t: T) => S): T[] {
    return sort(update(a, v, matchFn), matchFn);
  }

  export function remove<T, S>(a: T[] | undefined, key: S, matchFn: (t: T) => S): T[] {
    if (!a) {
      return [];
    }
    for (const idx in a) {
      const c = a[idx]
      if (matchFn(c) == key) {
        delete a[idx];
      }
    }
    return a;
  }
}
