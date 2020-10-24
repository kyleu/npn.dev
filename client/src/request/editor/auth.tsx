namespace request.editor {
  export function initAuthEditor(el: HTMLTextAreaElement) {
    const parent = el.parentElement!;
    parent.appendChild(createAuthEditor(el));
  }

  export function setAuth(input: HTMLInputElement, view: HTMLSpanElement, a: auth.Auth | undefined) {
    const url = new URL(input.value);
    let u = "";
    let p = "";
    if (a?.type === "basic") {
      const basic = a.config as auth.Basic
      u = encodeURIComponent(basic.username)
      p = encodeURIComponent(basic.password)
    }

    url.username = u;
    url.password = p;
    const us = url.toString();
    dom.setValue(input, us);
    dom.setContent(view, prototypeToHTML(prototypeFromURL(us)));
  }

  function createAuthEditor(el: HTMLTextAreaElement) {
    const a = json.parse(el.value) as auth.Auth;

    return <div class="uk-margin-top">
      {authSelect(el, a)}
      {auth.AllTypes.filter(t => !t.hidden).map(t => {
        let cfg = (a && a.type == t.key) ? a.config : null;
        return configEditor(t.key,  cfg, t.key === a?.type, el);
      })}
    </div>;
  }

  function authSelect(el: HTMLTextAreaElement, a: auth.Auth | undefined) {
    const ret = <select class="uk-select">
      <option value="">No body</option>
      {auth.AllTypes.filter(t => !t.hidden).map(t => {
        if (a && a.type === t.key) {
          return <option value={t.key} selected="selected">{t.title}</option>;
        } else {
          return <option value={t.key}>{t.title}</option>;
        }
      })}
    </select> as HTMLSelectElement;
    events(ret, () => {
      dom.els(".body-editor", ret.parentElement!).forEach(e => {
        const key = e.dataset["key"];
        if (ret.value === key) {
          if (key === "") {
            dom.setValue(el, "null");
            check();
          }
          e.classList.remove("hidden");
        } else {
          e.classList.add("hidden");
        }
      });
    });
    return ret;
  }

  function configEditor(key: string, config: any, active: boolean, el: HTMLTextAreaElement) {
    let cls = "uk-margin-top body-editor";
    if (!active) {
      cls += " hidden";
    }
    let e: JSX.Element;
    switch (key) {
      case "basic":
        e = basicEditor(key, active ? config as auth.Basic : undefined, el);
        break;
      default:
        e = <div>Unimplemented [{key}] auth editor</div>;
    }
    return <div class={cls} data-key={key}>{e}</div>;
  }

  function basicEditor(key: string, b: auth.Basic | undefined, el: HTMLTextAreaElement) {
    const ret = <textarea class="uk-textarea">{json.str(b)}</textarea> as HTMLTextAreaElement;
    const orig = b;
    events(ret, () => {
      let msg = ret.value;
      try { msg = json.parse(msg) } catch(e) {}
      const changed = diff.comp("", orig, msg, (k: string, lv: any, rv: any) => {});
      const n: auth.Auth = {type: "basic", config: msg};
      updateFn("json", n, el);
    });
    return ret;
  }

  function updateFn(t: string, cfg: auth.Auth | undefined, el: HTMLTextAreaElement) {
    const e = dom.req<HTMLInputElement>("#" + el.id.replace("-auth",  "-url"))
    const v = dom.req<HTMLSpanElement>("#" + el.id.replace("-auth",  "-urlview"))
    setAuth(e, v, cfg);
    dom.setValue(el, json.str(cfg));
    check();
  }
}
