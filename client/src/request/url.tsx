namespace request {
  export function prototypeToURL(p: Prototype): string {
    return prototypeToURLParts(p).map(x => x.v).join("");
  }

  export function prototypeToHTML(p: Prototype) {
    return <span>{prototypeToURLParts(p).map(x => <span title={x.t} class={urlColor(x.t)}>{ x.v }</span>)}</span>;
  }

  interface Part {
    readonly t: string;
    readonly v: string;
  }

  function prototypeToURLParts(p: Prototype): Part[] {
    const ret: Part[] = []
    let push = function(t: string, v: string) {
      ret.push({t: t, v: v});
    }

    push("protocol", p.protocol);
    push("", "://");
    if(p.auth) {
      for (let a of p.auth) {
        if (a.type === "basic") {
          const cfg = a.config as auth.Basic
          push("username", cfg.username);
          push("", ":");

          if (cfg.showPassword) {
            push("password", cfg.password);
          } else {
            push("password", "****");
          }
          push("", "@");
          break;
        }
      }
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
      var query = p.query.map(k => encodeURIComponent(k.k) + '=' + encodeURIComponent(k.v)).join('&');
      push("query", query);
    }
    if (p.fragment && p.fragment.length > 0) {
      push("", "#");
      push("fragment", encodeURIComponent(p.fragment));
    }

    return ret
  }

  function urlColor(key: string): string {
    switch (key) {
      case "username":
      case "password":
      case "protocol":
      case "auth":
        return "green-fg"
      case "domain":
      case "port":
        return "blue-fg"
      case "path":
        return "bluegrey-fg"
      case "query":
        return "purple-fg"
      default:
        return ""
    }
  }
}
