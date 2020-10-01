namespace nav {
  export let enabled = true;

  let handler = (p: string) => {
    console.warn("default nav handler called: " + p);
  }

  export function init(f: (p: string) => void) {
    handler = f;
    window.onpopstate = (event: PopStateEvent) => {
      if (event.state) {
        let s = event.state as string;
        handler(s);
      } else {
        handler("");
      }
    }
    let path = location.pathname;
    navigate(path);
  }

  export function pop() {
    let p = location.pathname.substr(0, location.pathname.lastIndexOf("/"));
    if (p === '/c') {
      p = "";
    }
    navigate(p);
  }

  export function navigate(path: string) {
    if (!enabled) {
      handler(path);
      return "";
    }
    if (path.startsWith("text/html;")) {
      return "";
    }
    if (path.startsWith("/")) {
      path = path.substr(1);
    }
    let locPath = location.pathname;
    if (locPath.startsWith("/")) {
      locPath = locPath.substr(1);
    }
    if (locPath !== path) {
      let final = path;
      history.pushState(final, "", "/" + final);
    }
    handler(path);
  }

  export function navActiveRequest() {
    navigate(`/c/${collection.cache.active}/${request.cache.active}`);
  }

  export interface LinkOpts {
    path: string;
    title: string;
    cls?: string;
    onclk?: string;
    isButton?: boolean;
    icon?: string;
  }

  export function link(o: LinkOpts) {
    let href = o.path;
    if (!href.startsWith("/")) {
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
      if (!o.onclk.endsWith(";")) {
        o.onclk += ";"
      }
    } else {
      o.onclk = "";
    }
    if (!o.isButton) {
      o.cls = style.linkColor + o.cls;
    }
    return <a class={o.cls} href={href} onclick={o.onclk + "nav.navigate('" + o.path + "', '" + o.title + "');return false;"}>{i}{o.title}</a>;
  }
}
