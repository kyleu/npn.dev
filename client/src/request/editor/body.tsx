namespace request.editor {
  export function initBodyEditor(el: HTMLTextAreaElement) {
    const parent = el.parentElement!;
    parent.appendChild(createBodyEditor(el));
  }

  function createBodyEditor(el: HTMLTextAreaElement) {
    const b = json.parse(el.value) as rbody.Body;

    return <div class="uk-margin-top">
      <select class="uk-select">
        <option value="">No body</option>
        {rbody.AllTypes.filter(t => !t.hidden).map(t => {
          if (b && b.type === t.key) {
            return <option value={t.key} selected="selected">{t.title}</option>;
          } else {
            return <option value={t.key}>{t.title}</option>;
          }
        })}˙
      </select>
      {rbody.AllTypes.filter(t => !t.hidden).map(t => {
        let cfg = (b && b.type == t.key) ? b.config : null;
        return configEditor(t.key,  cfg, t.key === (b ? b.type : ""));
      })}
    </div>;
  }

  function configEditor(key: string, config: any, active: boolean) {
    let cls = "uk-margin-top body-editor-" + key;
    if (!active) {
      cls += " hidden";
    }
    switch (key) {
      case "json":
        const j = config as rbody.JSONConfig
        return <div class={cls}><textarea class="uk-textarea">{json.str(j ? j.msg : null)}</textarea></div>;
      default:
        return <div class={cls}>Unimplemented [{key}] editor</div>;
    }
  }

  export function setBody(cache: Cache, body: rbody.Body | undefined) {

  }
}
