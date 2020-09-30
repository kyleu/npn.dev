namespace request {
  export interface Diff {
    readonly k: string;
    readonly l: any;
    readonly r: any;
  }

  export function diff(l: Request, r: Request) {
    const ret: Diff[] = [];
    const p = (k: string, lv: any, rv: any) => ret.push({k: k, l: lv, r: rv});
    const comp = (k: string, lv: any, rv: any) => {
      if (lv !== rv) {
        p(k, lv, rv);
      }
    }
    const checkNull = (k: string, lv: any, rv: any) => {
      if (!l) {
        if (r) {
          p(k, null, "(defined)")
        }
        return true;
      }
      if (!r) {
        p(k, "(defined)", null)
        return true;
      }
      return false;
    }

    if (checkNull("request", l, r)) {
      return ret;
    }

    comp("key", l.key, r.key);
    comp("title", l.title, r.title);
    comp("description", l.description, r.description);

    const lp = l.prototype;
    const rp = r.prototype;
    comp("method", lp.method, rp.method);
    comp("protocol", lp.protocol, rp.protocol);
    comp("domain", lp.domain, rp.domain);
    comp("port", lp.port, rp.port);
    comp("path", lp.path, rp.path);

    if(!checkNull("query", lp.query, rp.query)) {
      if ((!lp.query) || (!rp.query)) {
        return ret;
      }
      comp("query.length", lp.query.length, rp.query.length);
    }

    comp("fragment", lp.fragment, rp.fragment);

    if(!checkNull("headers", lp.headers, rp.headers)) {
      if ((!lp.headers) || (!rp.headers)) {
        return ret;
      }
      comp("headers.length", lp.headers.length, rp.headers.length);
    }

    if(!checkNull("auth", lp.auth, rp.auth)) {
      if ((!lp.auth) || (!rp.auth)) {
        return ret;
      }
      comp("auth.length", lp.auth.length, rp.auth.length);
    }

    if(!checkNull("body", lp.body, rp.body)) {
      if ((!lp.body) || (!rp.body)) {
        return ret;
      }
      comp("body.type", lp.body.type, rp.body.type);
      comp("body.config", lp.body.config, rp.body.config);
    }

    const lpo = lp.options;
    const rpo = rp.options;
    if (checkNull("options", lpo, rpo)) {
      return ret;
    }

    if ((!lpo) || (!rpo)) {
      return ret;
    }

    comp("timeout", lpo.timeout, rpo.timeout);
    comp("ignoreRedirects", lpo.ignoreRedirects, rpo.ignoreRedirects);
    comp("ignoreReferrer", lpo.ignoreReferrer, rpo.ignoreReferrer);
    comp("ignoreCerts", lpo.ignoreCerts, rpo.ignoreCerts);
    comp("ignoreCookies", lpo.ignoreCookies, rpo.ignoreCookies);
    comp("excludeDefaultHeaders", lpo.excludeDefaultHeaders, rpo.excludeDefaultHeaders);
    comp("readCookieJars", lpo.readCookieJars, rpo.readCookieJars);
    comp("writeCookieJar", lpo.writeCookieJar, rpo.writeCookieJar);
    comp("sslCert", lpo.sslCert, rpo.sslCert);
    comp("userAgentOverride", lpo.userAgentOverride, rpo.userAgentOverride);

    return ret;
  }
}
