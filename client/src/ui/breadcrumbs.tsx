namespace ui {
  export function setBreadcrumbs(coll: string | undefined, req: string | undefined, act: string | undefined) {
    const el = dom.req("#breadcrumbs");
    reset(el);
    if (coll) {
      el.appendChild(sep());
      el.appendChild(bcForColl(coll));
    }
    if (req) {
      el.appendChild(sep());
      el.appendChild(bcForReq(coll!, req));
    }
    if (act) {
      el.appendChild(sep());
      el.appendChild(bcForAct(coll!, req!, act));
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

  function bcForColl(coll: string) {
    return bcFor(coll, coll);
  }

  function bcForReq(coll: string, req: string) {
    return bcFor(req, coll, req);
  }

  function bcForAct(coll: string, req: string, act: string) {
    return bcFor(act, coll, req, act);
  }

  function bcFor(title: string, coll?: string, req?: string, act?: string) {
    if (act) {
      return nav.link("/c/" + coll + "/" + req + "/" + act, title, "uk-navbar-item uk-logo uk-margin-remove uk-padding-remove dynamic")
    }
    if (req) {
      return nav.link("/c/" + coll + "/" + req, title, "uk-navbar-item uk-logo uk-margin-remove uk-padding-remove dynamic")
    }
    return nav.link("/c/" + coll, title, "uk-navbar-item uk-logo uk-margin-remove uk-padding-remove dynamic")
  }
}
