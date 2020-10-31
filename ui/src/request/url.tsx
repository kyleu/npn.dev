import {MethodGet, Prototype} from "@/request/model";
import {trimPrefix, trimSuffix} from "@/util/string";
import {Basic} from "@/auth/basic";

export function urlToPrototype(url: string): Prototype {
  const u = new URL(url);
  return {
    method: MethodGet.key,
    protocol: trimSuffix(u.protocol, ":"),
    domain: u.hostname,
    port: u.port ? parseInt(u.port, 10) : undefined,
    path: trimPrefix(u.pathname, "/"),
    fragment: trimPrefix(u.hash, "#")
  };
}

interface Part {
  readonly t: string;
  readonly v: string;
}

function prototypeToURLParts(p: Prototype): Part[] {
  const ret: Part[] = []
  const push = (t: string, v: string): void => {
    ret.push({t: t, v: v});
  }

  push("protocol", p.protocol);
  push("", "://");
  if(p.auth && p.auth.type === "basic") {
    const cfg = p.auth.config as Basic
    push("username", cfg.username);
    push("", ":");

    if (cfg.showPassword) {
      push("password", cfg.password);
    } else {
      push("password", "****");
    }
    push("", "@");
  }
  push("domain", p.domain);
  if (p.port) {
    push("", ":");
    push("port", p.port.toString());
  }
  if (p.path && p.path.length > 0) {
    push("", "/");
    push("path", p.path);
  }
  if (p.query && p.query.length > 0) {
    push("", "?");
    const query = p.query.map(k => encodeURIComponent(k.k) + '=' + encodeURIComponent(k.v)).join('&');
    push("query", query);
  }
  if (p.fragment && p.fragment.length > 0) {
    push("", "#");
    push("fragment", encodeURIComponent(p.fragment));
  }

  return ret
}

export function prototypeToURL(p: Prototype | undefined): string {
  if (!p) {
    return "..."
  }
  return prototypeToURLParts(p).map(x => x.v).join("");
}

function urlColor(key: string): string {
  switch (key) {
    case "username":
    case "password":
    case "protocol":
    case "auth":
      return "green-fg";
    case "domain":
    case "port":
      return "blue-fg";
    case "path":
      return "bluegrey-fg";
    case "query":
      return "purple-fg";
    case "fragment":
      return "orange-fg";
    default:
      return "";
  }
}

export function prototypeToHTML(p: Prototype | undefined): JSX.Element {
  if (!p) {
    return <span>...</span>
  }
  return <span>{prototypeToURLParts(p).map(x => <span title={x.t} class={urlColor(x.t)}>{ x.v }</span>)}</span>;
}

export function prototypeBaseURL(p: Prototype | undefined): string {
  if (!p) {
    return "invalid";
  }
  let d = p.domain;
  if (p.port && p.port > 0) {
    d += `:${p.port}`;
  }
  return `${p.protocol}://${d}/`;
}

export function baseURL(s: string): string {
  return prototypeBaseURL(urlToPrototype(s))
}
