import {NPNRequest, Options} from "@/request/model";
import {jsonClone} from "@/util/json";
import {checkNull, comp, compArray, Diff} from "@/util/diff";

const debugDiff = false;

function debug(lpo: Options, rpo: Options, ret: Diff[]): void {
  if (debugDiff) {
    console.debug(jsonClone(lpo), jsonClone(rpo));
    if (ret.length > 0) {
      console.debug(ret);
    }
  }
}

export function diffRequests(l: NPNRequest | undefined, r: NPNRequest | undefined): Diff[] {
  const ret: Diff[] = [];
  const p = (k: string, lv: unknown, rv: unknown): number => ret.push({k: k, l: lv, r: rv});

  if ((!l || !r)) {
    checkNull("request", l, r, p);
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
  comp("auth", lp.auth, rp.auth, p);

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
  compArray("readCookieJars", lpo.readCookieJars, rpo.readCookieJars, p);
  comp("writeCookieJar", lpo.writeCookieJar, rpo.writeCookieJar, p);
  comp("sslCert", lpo.sslCert, rpo.sslCert, p);
  comp("userAgentOverride", lpo.userAgentOverride, rpo.userAgentOverride, p);

  debug(lpo, rpo, ret);

  return ret;
}
