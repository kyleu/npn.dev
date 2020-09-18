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
    if (path.startsWith("/w")) {
      path = path.substr(2)
    }
    navigate(path);
  }

  export function navigate(path: string) {
    if (path.startsWith("/")) {
      path = path.substr(1);
    }
    let fullpath = "/w";
    if (path.length > 0) {
      fullpath = fullpath + "/" + path;
    }
    if (location.pathname !== fullpath) {
      history.pushState(path, "", fullpath);
    }
    handler(path);
  }

  export function link(path: string, title: string) {
    return <a class={style.linkColor} href={path} onclick={"nav.navigate('" + path + "', '" + title + "');return false;"}>{title}</a>
  }
}
