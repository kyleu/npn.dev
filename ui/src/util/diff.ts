export interface Diff {
  readonly k: string;
  readonly l: unknown;
  readonly r: unknown;
}

export function comp(k: string, lv: unknown, rv: unknown, p: (k: string, lv: unknown, rv: unknown) => void): boolean {
  if (lv === undefined || lv === null) {
    lv = "";
  }
  if (rv === undefined || rv === null) {
    rv = "";
  }
  if (typeof lv === "object" && typeof rv === "object") {
    const lvo = lv as { [name: string]: unknown };
    const rvo = rv as { [name: string]: unknown };
    for (const f in lvo) {
      if (Object.prototype.hasOwnProperty.call(lvo, f)) {
        if (comp(k + "." + f, lvo[f], rvo[f], p)) {
          return true;
        }
      }
    }
    for (const f in rv) {
      if (Object.prototype.hasOwnProperty.call(rv, f) && !Object.prototype.hasOwnProperty.call(lv, f)) {
        if (comp(k + "." + f, lvo[f], rvo[f], p)) {
          return true;
        }
      }
    }
  } else {
    if (lv !== rv) {
      p(k, lv, rv);
      return true;
    }
  }
  return false;
}

export function compArray(k: string, lv: unknown[] | undefined, rv: unknown[] | undefined, p: (k: string, lv: unknown, rv: unknown) => void): boolean {
  if (lv === undefined || lv === null) {
    lv = [];
  }
  if (rv === undefined || rv === null) {
    rv = [];
  }
  if (lv.length !== rv.length) {
    p(k + ".length", lv.length, rv.length);
    return true;
  }
  for (let i = 0; i < lv.length; i++) {
    if (comp(k + "[" + i + "]", lv[i], rv[i], p)) {
      return true;
    }
  }
  return false;
}

export function checkNull(k: string, lv: unknown, rv: unknown, p: (k: string, lv: unknown, rv: unknown) => void): boolean {
  if (!lv) {
    if (rv) {
      p(k, null, "(defined)");
    }
    return true;
  }
  if (!rv) {
    p(k, "(defined)", null);
    return true;
  }
  return false;
}
