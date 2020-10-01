namespace str {
  export function trimPrefix(s: string, prefix: string) {
    if (s.startsWith(prefix)) {
      return s.slice(prefix.length);
    } else {
      return s;
    }
  }

  export function trimSuffix(s: string, suffix: string) {
    if (s.endsWith(suffix)) {
      return s.substring(0, s.lastIndexOf(suffix));
    } else {
      return s;
    }
  }
}
