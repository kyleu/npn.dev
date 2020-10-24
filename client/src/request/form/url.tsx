namespace request.form {
  export function renderURL(coll: string, r: request.Request) {
    const call = "nav.navigate(`/c/" + coll + "/" + r.key + "/call`);return false;"
    const url = request.prototypeToURL(r.prototype);
    return <div class="uk-margin-top uk-panel">
      <div class="left" style="width:120px;">
        <select class="uk-select" id={r.key + "-method"} name="method">
          {request.allMethods.map(m => {
            if(m.key === r.prototype.method) {
              return <option selected="selected">{ m.key }</option>;
            } else {
              return <option>{ m.key }</option>;
            }
          })}
        </select>
      </div>
      <div class="url-view uk-inline right" id={r.key + "-link"} style="width:calc(100% - 120px);">
        <a class="uk-form-icon uk-form-icon-flip" href="" onclick={call} title="send request" uk-icon="icon: play" />
        <div onclick={"request.editor.toggleURLEditor('" + r.key + "', true);"}>
          <span id={r.key + "-urlview"} class="url-link">{prototypeToHTML(r.prototype)}</span>
        </div>
      </div>
      <div class="url-input hidden uk-inline right" id={r.key + "-edit"} style="width:calc(100% - 120px);">
        <a class="uk-form-icon uk-form-icon-flip" href="" onclick={"request.editor.toggleURLEditor('" + r.key + "', false);return false;"} title="cancel edit" uk-icon="icon: close" />
        <form onsubmit={call}>
          <input class="uk-input" id={r.key + "-url"} name="url" type="text" value={url} data-lpignore="true" />
        </form>
      </div>
    </div>;
  }
}
