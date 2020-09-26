namespace request.form {
  export function renderFormPanel(coll: string, r: request.Request) {
    return <div>
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right">
          <a class="theme uk-icon" data-uk-icon="close" href="" onclick={"nav.navigate('/c/" + coll + "');return false;"} title="close request" />
        </div>
        <h3 class="uk-card-title">{r.title ? r.title : r.key}</h3>
        {renderURL(r)}
        {renderSavePanel(r)}
        {renderActions(coll, r)}
      </div>
      <div class="request-editor uk-card uk-card-body uk-card-default uk-margin-top">
        {renderSwitcher(r)}
      </div>
      <div class="request-action uk-card uk-card-body uk-card-default uk-margin-top hidden"/>
    </div>
  }

  export function renderDetails(r: request.Request) {
    return <li class="request-details-panel">
      <div class="uk-margin-top">
        <label class="uk-form-label" for={r.key + "-key"}>Key</label>
        <input class="uk-input" id={r.key + "-key"} name="key" type="text" value={ r.key || "" } data-lpignore="true" />
      </div>

      <div class="uk-margin-top">
        <label class="uk-form-label" for={r.key + "-title"}>Title</label>
        <input class="uk-input" id={r.key + "-title"} name="title" type="text" value={ r.title || "" } data-lpignore="true" />
      </div>

      <div class="uk-margin-top">
        <label class="uk-form-label" for={r.key + "-description"}>Description</label>
        <textarea class="uk-textarea" id={r.key + "-description"} name="description" data-lpignore="true">{ r.description || "" }</textarea>
      </div>
    </li>;
  }

  const transforms: any = {
    "http": "HTTP",
    "json": "JSON",
    "curl": "curl"
  };

  function renderSavePanel(r: request.Request) {
    return <div id="save-panel" class="right hiddenX">
      <button class="uk-button uk-button-default uk-margin-small-right uk-margin-top" onclick="console.log('TODO!');">Reset</button>
      <button class="uk-button uk-button-default uk-margin-top" onclick="console.log(request.form.getRequest());">Save Changes</button>
    </div>;
  }

  function renderActions(coll: string, r: request.Request) {
    const path = "/c/" + coll + "/" + r.key;
    const btnClass = "uk-button uk-button-default uk-margin-small-right uk-margin-top"
    const delWarn = "if (!confirm('Are you sure you want to delete request [" + r.key + "]?')) { return false; }"

    return <div>
      {nav.link(path + "/call", "Call", btnClass, "", true)}
      <div class="uk-inline">
        <button type="button" class={btnClass}>Export</button>
        <div id="export-dropdown" uk-dropdown="mode: click">
          <ul class="uk-list uk-list-divider" style="margin-bottom: 0;">
            {Object.keys(transforms).map(k => <li>{nav.link(path + "/transform/" + k, transforms[k], "", "UIkit.dropdown(dom.req('#export-dropdown')).hide(false);")}</li>)}
          </ul>
        </div>
      </div>
      {nav.link(path + "/delete", "Delete", btnClass, delWarn, true)}
    </div>
  }
}
