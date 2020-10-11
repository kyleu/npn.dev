namespace request.editor {
  export function initBodyEditor(el: HTMLTextAreaElement) {
    const parent = el.parentElement!;
    parent.appendChild(createBodyEditor(el));
  }

  export function setBody(cache: Cache, body: rbody.Body | undefined) {

  }

  function createBodyEditor(el: HTMLTextAreaElement) {
    const b = json.parse(el.value) as rbody.Body;

    return <div class="uk-margin-top">
      {bodySelect(el, b)}
      <div class="body-editor" data-key="" />
      {rbody.AllTypes.filter(t => !t.hidden).map(t => {
        let cfg = (b && b.type == t.key) ? b.config : null;
        return configEditor(t.key,  cfg, t.key === (b ? b.type : ""), el);
      })}
    </div>;
  }

  function bodySelect(el: HTMLTextAreaElement, b: rbody.Body) {
    const ret = <select class="uk-select">
      <option value="">No body</option>
      {rbody.AllTypes.filter(t => !t.hidden).map(t => {
        if (b && b.type === t.key) {
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
            el.value = "null";
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
    switch (key) {
      case "json":
        return <div class={cls} data-key={key}>{jsonEditor(key, active ? config as rbody.JSONConfig : undefined, el)}</div>;
      case "html":
        return <div class={cls} data-key={key}>{htmlEditor(key, active ? config as rbody.HTMLConfig : undefined, el)}</div>;
      default:
        return <div class={cls} data-key={key}>Unimplemented [{key}] editor</div>;
    }
  }

  function htmlEditor(key: string, h: rbody.HTMLConfig | undefined, el: HTMLTextAreaElement) {
    const ret = <textarea class="uk-textarea">{h ? h.content : ""}</textarea> as HTMLTextAreaElement;
    const orig = h ? h.content : "";
    events(ret, () => {
      const changed = orig !== ret.value;
      if (changed) {
        let msg = ret.value;
        updateFn("html", {content: msg}, el);
      }
    });
    return ret;
  }

  function jsonEditor(key: string, j: rbody.JSONConfig | undefined, el: HTMLTextAreaElement) {
    const ret = <textarea class="uk-textarea">{json.str(j ? j.msg : null)}</textarea> as HTMLTextAreaElement;
    const orig = j ? json.str(j.msg) : "null";
    events(ret, () => {
      const changed = orig !== ret.value;
      if (changed) {
        let msg = ret.value;
        try { msg = json.parse(msg) } catch(e) {}
        updateFn("json", {msg: msg}, el);
      }
    });
    return ret;
  }

  function updateFn(t: string, cfg: any, el: HTMLTextAreaElement) {
    const nb: rbody.Body = {type: t, config: cfg};
    el.value = json.str(nb);
    check();
  }
}
