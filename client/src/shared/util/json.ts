namespace json {
  export function str(x: any) {
    if (x === undefined) {
      return "null";
    }
    return JSON.stringify(x, null, 2);
  }

  export function parse<T>(s: string) {
    return JSON.parse(s) as T;
  }
}
