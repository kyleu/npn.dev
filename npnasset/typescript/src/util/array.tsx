namespace arr {
  export function find<T>(a: T[], predicate: (idx: number, x: T) => boolean) {
    const len = a.length >>> 0;
    let k = 0;
    while (k < len) {
      var kValue = a[k];
      if (predicate(k, kValue)) {
        return kValue;
      }
      k++;
    }
    return undefined;
  }
}
