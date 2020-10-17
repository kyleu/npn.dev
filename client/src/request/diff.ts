namespace diff {
  export interface Diff {
    readonly k: string;
    readonly l: any;
    readonly r: any;
  }

  export function diff(l: request.Request, r: request.Request) {
    const ret: Diff[] = [];
    const p = (k: string, lv: any, rv: any) => ret.push({k: k, l: lv, r: rv});

    if (checkNull("request", l, r, p)) {
      return ret;
    }

    comp("key", l.key, r.key, p);
    comp("title", l.title, r.title, p);
    comp("description", l.description, r.description, p);

    const lp = l.prototype;
    const rp = r.prototype;

    comp("method", lp.method, rp.method, p);
    comp("protocol", lp.protocol, rp.protocol, p);
    comp("domain", lp.domain, rp.domain, p);
    comp("port", lp.port, rp.port, p);
    comp("path", lp.path, rp.path, p);
    compArray("query", lp.query, rp.query, p);
    comp("fragment", lp.fragment, rp.fragment, p);
    compArray("headers", lp.headers, rp.headers, p);
    compArray("auth", lp.auth, rp.auth, p);

    if(!checkNull("body", lp.body, rp.body, p)) {
      if (lp.body && rp.body) {
        comp("body.type", lp.body.type, rp.body.type, p);
        comp("body.config", lp.body.config, rp.body.config, p);
      }
    }

    const lpo = lp.options;
    const rpo = rp.options;
    if (checkNull("options", lpo, rpo, p)) {
      return ret;
    }

    if ((!lpo) || (!rpo)) {
      return ret;
    }

    comp("timeout", lpo.timeout, rpo.timeout, p);
    comp("ignoreRedirects", lpo.ignoreRedirects, rpo.ignoreRedirects, p);
    comp("ignoreReferrer", lpo.ignoreReferrer, rpo.ignoreReferrer, p);
    comp("ignoreCerts", lpo.ignoreCerts, rpo.ignoreCerts, p);
    comp("ignoreCookies", lpo.ignoreCookies, rpo.ignoreCookies, p);
    compArray("excludeDefaultHeaders", lpo.excludeDefaultHeaders, rpo.excludeDefaultHeaders, p);
    comp("readCookieJars", lpo.readCookieJars, rpo.readCookieJars, p);
    comp("writeCookieJar", lpo.writeCookieJar, rpo.writeCookieJar, p);
    comp("sslCert", lpo.sslCert, rpo.sslCert, p);
    comp("userAgentOverride", lpo.userAgentOverride, rpo.userAgentOverride, p);

    return ret;
  }

  export function comp(k: string, lv: any, rv: any, p: (k: string, lv: any, rv: any) => void): boolean {
    if (lv === undefined || lv === null) {
      lv = "";
    }
    if (rv === undefined || rv === null) {
      rv = "";
    }
    if (typeof lv === "object" && typeof rv === "object") {
      for (const f in lv) {
        if (lv.hasOwnProperty(f)) {
          if(comp(k + "." + f, lv[f], rv[f], p)) {
            return true;
          }
        }
      }
      for (const f in rv) {
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
}
