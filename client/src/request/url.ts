namespace request {
  export interface Part {
    readonly t: string;
    readonly v: string;
  }

  export function prototypeToURLParts(p: Prototype): request.Part[] {
    const ret: Part[] = []
    let push = function(t: string, v: string) {
      ret.push({t: t, v: v});
    }

    push("protocol", p.protocol);
    push("", "://");
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
      var query = p.query.map(k => encodeURIComponent(k.k) + '=' + encodeURIComponent(k.v)).join('&');
      push("query", query);
    }
    if (p.fragment && p.fragment.length > 0) {
      push("", "#");
      push("fragment", encodeURIComponent(p.fragment));
    }

    return ret
  }

  export function prototypeToURL(p: Prototype): string {
    return prototypeToURLParts(p).map(x => x.v).join("");
  }
}
