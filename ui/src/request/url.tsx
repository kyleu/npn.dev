import {Prototype} from "@/request/model";
import {Basic} from "@/auth/basic";

export interface Part {
  readonly t: string;
  readonly v: string;
  readonly idx: number;
  readonly color: string;
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

export function prototypeToURLParts(p: Prototype): Part[] {
  const ret: Part[] = []
  const push = (t: string, v: string): void => {
    ret.push({t: t, v: v, idx: ret.length, color: urlColor(t)});
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
