export function startsWith(str: string, search: string, pos?: number): boolean {
  pos = (pos && pos > 0) ? pos | 0 : 0;
  return str.substring(pos, pos + search.length) === search;
}

export function endsWith(str: string, search: string, pos?: number): boolean {
  pos = (pos && pos < str.length) ? pos : str.length;
  return str.substring(pos - search.length, pos) === search;
}

export function trimPrefix(s: string, prefix: string): string {
  if (startsWith(s, prefix)) {
    return s.slice(prefix.length);
  } else {
    return s;
  }
}

export function trimSuffix(s: string, suffix: string): string {
  if (endsWith(s, suffix)) {
    return s.substring(0, s.lastIndexOf(suffix));
  } else {
    return s;
  }
}
