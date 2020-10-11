namespace nav {
  export let enabled = true;

  let handler = (p: string, hash: string) => {
    console.warn("default nav handler called: " + p + ((hash.length > 0) ? ("#" + hash) : ""));
  }

  export function init(f: (p: string, hash: string) => void) {
    handler = f;
    window.onpopstate = (event: PopStateEvent) => {
      if (event.state) {
        let s = event.state as string;
        let [path, hash] = extractHash(s);
        handler(path, hash);
      } else {
        handler("", "");
      }
    }
    let path = location.pathname + ((location.hash.length > 0) ? (location.hash) : "");
    navigate(path);
  }

  export function navigate(s: string) {
    let [path, hash] = extractHash(s);
    if (!enabled) {
      handler(path, hash);
      return "";
    }
    if (str.startsWith(path, "text/html;")) {
      return "";
    }
    if (str.startsWith(path, "/")) {
      path = path.substr(1);
    }
    let locPath = location.pathname;
    if (str.startsWith(locPath, "/")) {
      locPath = locPath.substr(1);
    }
    if (locPath !== path) {
      let final = path;
      history.pushState(final, "", "/" + final);
    }
    handler(path, hash);
  }

  export function pop() {
    let p = location.pathname.substr(0, location.pathname.lastIndexOf("/"));
    if (p === '/c') {
      p = "";
    }
    navigate(p);
  }

  export function hashLink(k: string, title: string, active: string) {
    if (active.indexOf("#") === 0) {
      active = active.substr(1);
    }
    let cls = "";
    if (active === k) {
      cls += "uk-active";
    }
    const ret = <li class={cls}><a href={"#" + k}>{title}</a></li>;
    ret.onclick = () => {
      history.replaceState(history.state, "", `#${k}`);
    };
    return ret;
  }

  export interface LinkOpts {
    path: string;
    hash?: string;
    title: string;
    cls?: string;
    onclk?: string;
    isButton?: boolean;
    icon?: string;
  }

  export function link(o: LinkOpts) {
    let href = o.path;
    if (!str.startsWith(href, "/")) {
      href = "/" + href;
    }
    if (o.cls) {
      o.cls = " " + o.cls.trim();
    } else {
      o.cls = "";
    }
    let i = <span />;
    if (o.icon) {
      i = <span class="nav-icon" data-uk-icon={`icon: ${o.icon}`} />;
    }
    if (o.onclk) {
      if (!str.endsWith(o.onclk, ";")) {
        o.onclk += ";"
      }
    } else {
      o.onclk = "";
    }
    if (!o.isButton) {
      o.cls = style.linkColor + o.cls;
    }
    let p = o.path;
    if (o.hash && o.hash.length > 0) {
      p += "#" + o.hash
    }
    return <a class={o.cls} href={href} onclick={o.onclk + "nav.navigate('" + p + "', '" + o.title + "');return false;"}>{i}{o.title}</a>;
  }

  function extractHash(s: string): [string, string] {
    let path = s;
    let hash = "";
    const hashIdx = path.indexOf("#");
    if (hashIdx > -1) {
      hash = path.substr(hashIdx + 1);
      path = path.substr(0, hashIdx);
    }
    return [path, hash];
  }
}
