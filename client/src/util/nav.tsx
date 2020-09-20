namespace nav {
  let handler = function (p: string) {
    console.info("default nav handler called: " + p);
  }

  export function init(f: (p: string) => void) {
    handler = f;
    window.onpopstate = function (event: PopStateEvent) {
      f(event.state === null ? "" : (event.state as string));
    }
    let path = location.pathname;
    navigate(path);
  }

  export function navigate(path: string) {
    if (path.startsWith("/")) {
      path = path.substr(1);
    }
    if (location.pathname !== path) {
      history.pushState(path, "", "/" + path);
    }
    handler(path);
  }

  export function pop() {
    let p = location.pathname.substr(0, location.pathname.lastIndexOf("/"));
    if (p === '/c') {
      p = ""
    }
    navigate(p)
  }

  export function link(path: string, title: string, cls?: string, onclk?: string) {
    let href = path;
    if (!href.startsWith("/")) {
      href = "/" + href;
    }
    if (cls) {
      cls = " " + cls.trim();
    } else {
      cls = "";
    }
    if (onclk) {
      if (!onclk.endsWith(";")) {
        onclk += ";"
      }
    } else {
      onclk = "";
    }
    return <a class={style.linkColor + cls} href={href} onclick={onclk + "nav.navigate('" + path + "', '" + title + "');return false;"}>{title}</a>;
  }
}
