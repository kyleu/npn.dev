export function jsonStr(x: any | undefined) {
  if (x === undefined) {
    return "null";
  }
  return JSON.stringify(x, null, 2);
}

export function jsonParse<T>(s: string) {
  return JSON.parse(s) as T;
}
