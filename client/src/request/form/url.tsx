namespace request.form {
  export function renderURL(coll: string, r: request.Request) {
    const call = "nav.navigate(`/c/" + coll + "/" + r.key + "/call`);return false;"
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
      <div class="uk-inline right" style="width:calc(100% - 120px);">
        <a class="uk-form-icon uk-form-icon-flip" href="" onclick={call} title="send request" uk-icon="icon: play" />
        <form onsubmit={call}>
          <input class="uk-input" id={r.key + "-url"} name="url" type="text" value={request.prototypeToURL(r.prototype)} data-lpignore="true" />
        </form>
      </div>
    </div>;
  }
}
