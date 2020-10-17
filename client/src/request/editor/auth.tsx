namespace request.editor {
  export function initAuthEditor(el: HTMLTextAreaElement) {
    const parent = el.parentElement!;
    parent.appendChild(createAuthEditor(el));
  }

  export function setAuth(cache: Cache, auth: auth.Auth[] | undefined) {
    const url = new URL(cache.url.value);
    let u = "";
    let p = "";
    if (auth) {
      for (let a of auth) {
        if (a.type === "basic") {
          const basic = a.config as auth.Basic
          u = encodeURIComponent(basic.username)
          p = encodeURIComponent(basic.password)
        }
      }
    }

    url.username = u;
    url.password = p;
    cache.url.value = url.toString();
  }

  function createAuthEditor(el: HTMLTextAreaElement) {
    let authSet = (json.parse(el.value) as auth.Auth[]) || [];
    const editors = authSet.map(a => {
      switch (a.type) {
        case "basic":
          const b = a.config as auth.Basic;
          return <div>BASIC: {json.str(b)}</div>;
        default:
          return <div>{a.type}</div>;
      }
    });
    return <div class="uk-margin-top">
      <div class="current-auth-editors">{editors}</div>
      <a class={style.linkColor} href="" onclick="return false;">Add New Auth</a>
    </div>;
  }

  function updateFn(t: string, cfg: any, el: HTMLTextAreaElement) {
    const nb: any[] = [];
    el.value = json.str(nb);
    check();
  }
}
