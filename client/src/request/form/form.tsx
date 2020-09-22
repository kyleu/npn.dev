namespace request.form {
  export function renderForm(coll: string, r: request.Request) {
    const key = r.key;
    const p = r.prototype;

    return <form class="uk-form-stacked">
      <input type="hidden" name="coll" value={coll} />
      <input type="hidden" name="originalKey" value={r.key} />
      <fieldset class="uk-fieldset">
        <legend class="hidden">request form</legend>

        <div>
          <div class="left" style="width:120px;">
            <select class="uk-select" id={key + "-method"} name="method">
              {request.allMethods.map(m => {
                if(m.key === p.method) {
                  return <option selected="selected">{ m.key }</option>;
                } else {
                  return <option>{ m.key }</option>;
                }
              })}
            </select>
          </div>
          <div class="right" style="width:calc(100% - 120px);">
            <input class="uk-input" id={key + "-url"} name="url" type="text" value={request.prototypeToURL(p)} />
          </div>
          <div class="clear"/>
        </div>

        <hr />

        <div class="uk-margin-top">
          <ul data-uk-tab="">
            <li><a href="#">Details</a></li>
            <li><a href="#">Query</a></li>
            <li><a href="#">Auth</a></li>
            <li><a href="#">Headers</a></li>
            <li><a href="#">Body</a></li>
            <li><a href="#">Options</a></li>
          </ul>
          <ul class="uk-switcher uk-margin">
            <li class="request-details-panel">
              <div class="uk-margin-top">
                <label class="uk-form-label" for={key + "-key"}>Key</label>
                <input class="uk-input" id={key + "-key"} name="key" type="text" value={ r.key || "" } />
              </div>

              <div class="uk-margin-top">
                <label class="uk-form-label" for={key + "-title"}>Title</label>
                <input class="uk-input" id={key + "-title"} name="title" type="text" value={ r.title || "" } />
              </div>

              <div class="uk-margin-top">
                <label class="uk-form-label" for={key + "-description"}>Description</label>
                <input class="uk-input" id={key + "-description"} name="description" type="text" value={ r.description || "" } />
              </div>
            </li>
            <li class="request-url-panel">
              <div class="uk-margin-top">
                <label class="uk-form-label" for={key + "-queryparams"}>Query Params</label>
                <textarea class="uk-textarea" id={key + "-queryparams"} name="queryparams">{p.query ? JSON.stringify(p.query, null, 2) : "null"}</textarea>
              </div>
            </li>
            <li class="request-auth-panel">
              <div class="uk-margin-top">
                <label class="uk-form-label" for={key + "-auth"}>Auth</label>
                <textarea class="uk-textarea" id={key + "-auth"} name="auth">{p.auth ? JSON.stringify(p.auth, null, 2) : "null"}</textarea>
              </div>
            </li>
            <li class="request-headers-panel">
              <div class="uk-margin-top">
                <label class="uk-form-label" for={key + "-headers"}>Headers</label>
                <textarea class="uk-textarea" id={key + "-headers"} name="headers">{p.headers ? JSON.stringify(p.headers, null, 2) : "null"}</textarea>
              </div>
            </li>
            <li class="request-body-panel">
              <div class="uk-margin-top">
                <label class="uk-form-label" for={key + "-body"}>Body</label>
                <textarea class="uk-textarea" id={key + "-body"} name="body">{p.body ? JSON.stringify(p.body, null, 2) : "null"}</textarea>
              </div>
            </li>
            <li class="request-options-panel">
              <div class="uk-margin-top">
                <label class="uk-form-label" for={key + "-options"}>Options</label>
                <textarea class="uk-textarea" id={key + "-options"} name="options">{p.options ? JSON.stringify(p.options, null, 2) : "null"}</textarea>
              </div>
            </li>
          </ul>
        </div>

        <div class="uk-margin-top">
          <button class="right uk-button uk-button-default uk-margin-top" type="submit">Save Changes</button>
          {nav.link("/c/" + coll + "/" + r.key, "Cancel", "right uk-button uk-button-default uk-margin-top uk-margin-right")}
        </div>
      </fieldset>
    </form>;
  }
}
