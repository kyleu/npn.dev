namespace request {
  export function urlToPrototype(url: string): Prototype {
    const u = new URL(url);
    return {
      method: MethodGet.key,
      protocol: str.trimSuffix(u.protocol, ":"),
      domain: u.hostname,
      port: u.port ? parseInt(u.port, 10) : undefined,
      path: str.trimPrefix(u.pathname, "/"),
      fragment: str.trimPrefix(u.hash, "#")
    };
  }

  export function prototypeToURL(p: Prototype): string {
    return prototypeToURLParts(p).map(x => x.v).join("");
  }

  export function prototypeToHTML(p: Prototype) {
    return <span>{prototypeToURLParts(p).map(x => <span title={x.t} class={urlColor(x.t)}>{ x.v }</span>)}</span>;
  }

  export function baseURL(s: string) {
    return prototypeBaseURL(urlToPrototype(s))
  }

  export function prototypeBaseURL(p: Prototype | undefined) {
    if (!p) {
      return "invalid";
    }
    let d = p.domain;
    if (p.port && p.port > 0) {
      d += `:${p.port}`;
    }
    return `${p.protocol}://${d}/`;
  }

  interface Part {
    readonly t: string;
    readonly v: string;
  }

  function prototypeToURLParts(p: Prototype): Part[] {
    const ret: Part[] = []
    let push = (t: string, v: string) => {
      ret.push({t: t, v: v});
    }

    push("protocol", p.protocol);
    push("", "://");
    if(p.auth && p.auth.type === "basic") {
      const cfg = p.auth.config as auth.Basic
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
}
