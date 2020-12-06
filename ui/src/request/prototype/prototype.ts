import {Prototype, QueryParam} from "@/request/model";
import {Auth} from "@/auth/model";
import {endsWith, splitString, startsWith} from "@/util/string";

function newPrototype(protocol: string, hostname: string, port: number | undefined, path: string, qp: QueryParam[], fragment: string, auth: Auth | undefined): Prototype {
  if (endsWith(protocol, ":")) {
    protocol = protocol.substr(0, protocol.length - 1);
  }
  if (startsWith(fragment, "#")) {
    fragment = fragment.substr(1);
  }

  return {method: "get", protocol: protocol, domain: hostname, port: port, path: path, query: qp, fragment: fragment, auth: auth};
}

export function prototypeFromURL(u: string): Prototype {
  const [r1, frag] = splitString(u, "#", true);
  const [r2, query] = splitString(r1, "?", true);

  const qp: QueryParam[] = [];
  const qpSplit = query.split("&");
  for (const q in qpSplit) {
    const [k, v] = splitString(qpSplit[q], "=", true);
    qp.push({k, v});
  }

  let [proto, r3] = splitString(r2, ":", true);
  if (r3.length === 0) {
    r3 = proto;
    proto = "http";
  }
  while(r3.startsWith("/")) {
    r3 = r3.substr(1);
  }
  const [r4, path] = splitString(r3, "/", true);

  let [aut, hostOrig] = splitString(r4, '@', true);
  if (hostOrig === "") {
    hostOrig = aut;
    aut = "";
  }
  const [host, portString] = splitString(hostOrig, ':', true);
  let port = undefined;
  if (portString.length > 0) {
    port = parseInt(portString);
  }

  let at: Auth | undefined;

  if (aut !== "") {
    const [user, pass] = splitString(aut, ':', true);
    at = { "type": "basic", config: {"username": user, "password": pass, "showPassword": false}};
  }

  return newPrototype(proto, host, port, path, qp, frag, at);
}
