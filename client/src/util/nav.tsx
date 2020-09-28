namespace nav {
  let handler = (p: string) => {
    console.info("default nav handler called: " + p);
  }

  export function init(f: (p: string) => void) {
    handler = f;
    window.onpopstate = (event: PopStateEvent) => {
      f(event.state === null ? "" : (event.state as string));
    }
    let path = location.pathname;
    navigate(path);
  }

  export function navigate(path: string) {
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
      history.pushState(path, "", "/" + path);
    }
    handler(path);
  }

  export function pop() {
    let p = location.pathname.substr(0, location.pathname.lastIndexOf("/"));
    if (p === '/c') {
      p = "";
    }
    navigate(p);
  }

  export function navActiveRequest() {
    navigate(`/c/${collection.cache.active}/${request.cache.active}`);
  }

  export function link(path: string, title: string, cls?: string, onclk?: string, isButton?: boolean, icon?: string) {
    let href = path;
    if (!href.startsWith("/")) {
      href = "/" + href;
    }
    if (cls) {
      cls = " " + cls.trim();
    } else {
      cls = "";
    }
    let i = <span />;
    if (icon) {
      i = <span class="nav-icon" data-uk-icon={`icon: ${icon}`} />;
    }
    if (onclk) {
      if (!onclk.endsWith(";")) {
        onclk += ";"
      }
    } else {
      onclk = "";
    }
    if (!isButton) {
      cls = style.linkColor + cls;
    }
    return <a class={cls} href={href} onclick={onclk + "nav.navigate('" + path + "', '" + title + "');return false;"}>{i}{title}</a>;
  }
}
