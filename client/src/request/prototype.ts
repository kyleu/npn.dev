namespace request {
  function newPrototype(protocol: string, hostname: string, port: number | undefined, path: string, qp: request.QueryParam[], fragment: string, auth: auth.Auth[]): Prototype {
    if (protocol.endsWith(":")) {
      protocol = protocol.substr(0, protocol.length - 1);
    }
    if (fragment.startsWith("#")) {
      fragment = fragment.substr(1);
    }

    return {method: "get", protocol: protocol, domain: hostname, port: port, path: path, query: qp, fragment: fragment, auth: auth};
  }

  export function prototypeFromURL(u: string): Prototype {
    const url = new URL(u);

    const qp: QueryParam[] = []
    for (const [k, v] of url.searchParams) {
      qp.push({k: k, v: v});
    }

    const auth: auth.Auth[] = [];
    if(url.username.length > 0) {
      auth.push({type: "basic", config: {"username": url.username, "password": url.password, "showPassword": true}})
    }

    let port: number | undefined;
    if (url.port.length > 0) {
      port = parseInt(url.port);
    }

    return newPrototype(url.protocol, url.hostname, port, url.pathname, qp, url.hash, auth);
  }
}
