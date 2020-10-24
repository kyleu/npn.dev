namespace request {
  function newPrototype(protocol: string, hostname: string, port: number | undefined, path: string, qp: request.QueryParam[], fragment: string, auth: auth.Auth | undefined): Prototype {
    if (str.endsWith(protocol, ":")) {
      protocol = protocol.substr(0, protocol.length - 1);
    }
    if (str.startsWith(fragment, "#")) {
      fragment = fragment.substr(1);
    }

    return {method: "get", protocol: protocol, domain: hostname, port: port, path: path, query: qp, fragment: fragment, auth: auth};
  }

  export function prototypeFromURL(u: string): Prototype {
    const url = new URL(u);

    const qp: QueryParam[] = []
    url.searchParams.forEach((v, k) => qp.push({k: k, v: v}));

    let auth: auth.Auth | undefined;
    if(url.username.length > 0) {
      auth = {type: "basic", config: {"username": url.username, "password": url.password, "showPassword": true}};
    }

    let port: number | undefined;
    if (url.port && url.port.length > 0) {
      port = parseInt(url.port, 10);
    }

    let path = url.pathname;
    if (path.indexOf("/") === 0) {
      path = path.substr(1);
    }

    return newPrototype(url.protocol, url.hostname, port, path, qp, url.hash, auth);
  }
}
