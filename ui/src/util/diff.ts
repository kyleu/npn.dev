export interface Diff {
  readonly k: string;
  // @ts-ignore
  // eslint-disable-next-line
  readonly l: any;
  // @ts-ignore
  // eslint-disable-next-line
  readonly r: any;
}

// @ts-ignore
// eslint-disable-next-line
export function comp(k: string, lv: any, rv: any, p: (k: string, lv: any, rv: any) => void): boolean {
  if (lv === undefined || lv === null) {
    lv = "";
  }
  if (rv === undefined || rv === null) {
    rv = "";
  }
  if (typeof lv === "object" && typeof rv === "object") {
    for (const f in lv) {
      // @ts-ignore
      // eslint-disable-next-line
      if (lv.hasOwnProperty(f)) {
        if(comp(k + "." + f, lv[f], rv[f], p)) {
          return true;
        }
      }
    }
    for (const f in rv) {
      // @ts-ignore
      // eslint-disable-next-line
      if (rv.hasOwnProperty(f) && !lv.hasOwnProperty(f)) {
        if (comp(k + "." + f, lv[f], rv[f], p)) {
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

// @ts-ignore
// eslint-disable-next-line
export function compArray(k: string, lv: any[] | undefined, rv: any[] | undefined, p: (k: string, lv: any, rv: any) => void): boolean {
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

// @ts-ignore
// eslint-disable-next-line
export function checkNull(k: string, lv: any, rv: any, p: (k: string, lv: any, rv: any) => void): boolean {
  if (!lv) {
    if (rv) {
      p(k, null, "(defined)")
    }
    return true;
  }
  if (!rv) {
    p(k, "(defined)", null)
    return true;
  }
  return false;
}
