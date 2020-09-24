namespace ui {
  export function setBreadcrumbs(coll: string | undefined, req: string | undefined, act: string | undefined, extra: string[]) {
    const el = dom.req("#breadcrumbs");
    reset(el);
    if (coll) {
      el.appendChild(sep());
      el.appendChild(bcFor(coll, "c", coll));
      if (req) {
        el.appendChild(sep());
        el.appendChild(bcFor(req, "c", coll, req));
        if (act) {
          el.appendChild(sep());
          el.appendChild(bcFor(act, "c", coll, req, act));
          if (extra && extra.length > 0) {
            for (let i = 0; i < extra.length; i++) {
              el.appendChild(sep());
              const ret = [coll, req, act];
              ret.push(...extra.slice(0, i))
              el.appendChild(bcFor(extra[i], ...ret));
            }
          }
        }
      }
    }
  }

  export function reset(el: HTMLElement) {
    for (let i = el.childElementCount - 1; i >= 0; i--) {
      const e = el.children[i];
      if (e.classList.contains("dynamic")) {
        el.removeChild(e);
      }
    }
  }

  function sep() {
    return <span class="uk-navbar-item dynamic" style="padding: 0 8px;"> / </span>
  }

  function bcForExtra(coll: string, req: string, act: string, extra: string[]) {
    return bcFor(act, "c", coll, req, act);
  }

  function bcFor(title: string, ...parts: string[]) {
    const path = parts.map(s => "/" + s).join("");
    return nav.link(path, title, "uk-navbar-item uk-logo uk-margin-remove uk-padding-remove dynamic");
  }
}
