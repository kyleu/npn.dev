export function jsonStr(x: unknown): string {
  if (x === undefined) {
    return "null";
  }
  return JSON.stringify(x, null, 2);
}

export function jsonParse<T>(s: string): T {
  return JSON.parse(s);
}

export function jsonParseTry<T>(s: string): T | string {
  try {
    return JSON.parse(s);
  } catch (_) {
    return s;
  }
}

export function jsonClone<T>(x: T): T {
  return jsonParse(jsonStr(x));
}
