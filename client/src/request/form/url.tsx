namespace request.form {
  export function renderURL(r: request.Request) {
    return <div>
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
        <a class="uk-form-icon uk-form-icon-flip" href="" onclick="return false;" uk-icon="icon: refresh" />
        <input class="uk-input" id={r.key + "-url"} name="url" type="text" value={request.prototypeToURL(r.prototype)} data-lpignore="true" />
      </div>
      <div class="clear"/>
    </div>;
  }
}
