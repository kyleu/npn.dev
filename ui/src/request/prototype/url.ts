import {Prototype} from "@/request/model";
import {Basic} from "@/auth/basic";
import {profileRef} from "@/user/profile";

export interface Part {
  readonly t: string;
  readonly v: string;
  readonly idx: number;
  readonly color: string;
}

function urlColor(key: string): string {
  const dark = profileRef.value?.settings.mode === "dark";
  switch (key) {
    case "protocol":
      return dark ? "#ce494a" : "#5a0000";
    case "auth":
    case "username":
    case "password":
      return dark ? "#e9af41" : "#7d5d14";
    case "domain":
    case "port":
      return dark ? "#47a569" : "#034400";
    case "path":
      return dark ? "#b35da6" : "#001149";
    case "query":
      return dark ? "#6689c3" : "#001149";
    case "fragment":
      return dark ? "#b35da6" : "#001149";
    default:
      return "";
  }
}

export function prototypeToURLParts(p: Prototype): Part[] {
  const ret: Part[] = [];
  const push = (t: string, v: string): void => {
    ret.push({t: t, v: v, idx: ret.length, color: urlColor(t)});
  };

  push("protocol", p.protocol);
  push("", "://");
  if(p.auth && p.auth.type === "basic") {
    const cfg = p.auth.config as Basic;
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
  if (p.query) {
    const qp = p.query.filter(x => x.k && x.k.trim().length > 0);
    if (qp.length > 0) {
      push("", "?");
      const query = qp.map(k => {
        if (k.v.length > 0) {
          return encodeURIComponent(k.k) + '=' + encodeURIComponent(k.v);
        }
        return encodeURIComponent(k.k);
      }).join('&');
      push("query", query);
    }
  }
  if (p.fragment && p.fragment.length > 0) {
    push("", "#");
    push("fragment", encodeURIComponent(p.fragment));
  }

  return ret;
}

export function prototypeToURL(p: Prototype | undefined): string {
  if (!p) {
    return "...";
  }
  return prototypeToURLParts(p).map(x => x.v).join("");
}
