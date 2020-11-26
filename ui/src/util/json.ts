// @ts-ignore
// eslint-disable-next-line
export function jsonStr(x: any | undefined): string {
  if (x === undefined) {
    return "null";
  }
  return JSON.stringify(x, null, 2);
}

export function jsonParse<T>(s: string): T {
  return JSON.parse(s);
}

export function jsonClone<T>(x: T): T {
  return jsonParse(jsonStr(x));
}
