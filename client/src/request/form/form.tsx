namespace request.form {
  export function renderFormPanel(coll: string, r: request.Request) {
    return <div>
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right">
          <a class="theme uk-icon" data-uk-icon="close" href="" onclick={"nav.navigate('/c/" + coll + "');return false;"} title="close request" />
        </div>
        <h3 class="uk-card-title">{r.title ? r.title : r.key}</h3>
        {renderURL(r)}
        {renderActions(coll, r)}
      </div>
      <div class="request-editor uk-card uk-card-body uk-card-default uk-margin-top">
        {renderSwitcher(r)}
        <div class="uk-margin-top hidden">
          <button class="right uk-button uk-button-default uk-margin-top" type="submit">Save Changes</button>
        </div>
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
        <input class="uk-input" id={r.key + "-description"} name="description" type="text" value={ r.description || "" } data-lpignore="true" />
      </div>
    </li>;
  }

  const transforms: any = {
    "http": "HTTP",
    "json": "JSON",
    "curl": "curl"
  };

  function renderActions(coll: string, r: request.Request) {
    const path = "/c/" + coll + "/" + r.key;
    return <div class="uk-margin-top">
      {nav.link(path + "/call", "Call", "uk-button uk-button-default uk-margin-small-right", "", true)}
      <div class="uk-inline">
        <button type="button" class="uk-button uk-button-default uk-margin-small-right">Export</button>
        <div id="export-dropdown" uk-dropdown="mode: click">
          <ul class="uk-list uk-list-divider" style="margin-bottom: 0;">
            {Object.keys(transforms).map(k => <li>{nav.link(path + "/transform/" + k, transforms[k], "", "UIkit.dropdown(dom.req('#export-dropdown')).hide(false);")}</li>)}
          </ul>
        </div>
      </div>
      {nav.link(path + "/delete", "Delete", "uk-button uk-button-default uk-margin-small-right", "if (!confirm('Are you sure you want to delete request [" + r.key + "]?')) { return false; }", true)}
    </div>
  }
}
