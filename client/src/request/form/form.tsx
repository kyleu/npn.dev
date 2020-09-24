namespace request.form {
  export function renderFormPanel(coll: string, r: request.Request) {
    return <form class="uk-form-stacked" action={"/browse/" + coll + "/" + r.key + "/save"} method="post" onsubmit="return false;">
      <input type="hidden" name="coll" value={coll} />
      <input type="hidden" name="originalKey" value={r.key} />
      <fieldset class="uk-fieldset">
        <legend class="hidden">request form</legend>
        <div class="uk-card uk-card-body uk-card-default uk-margin-top">
          <div class="right">
            <a class="theme uk-icon" data-uk-icon="close" href="" onclick="nav.pop();return false;" title="close collection" />
          </div>
          <h3 class="uk-card-title">{r.title ? r.title : r.key}</h3>
          {renderURL(r)}
        </div>
        <div class="uk-card uk-card-body uk-card-default uk-margin-top">
          {renderSwitcher(r)}
          <div class="uk-margin-top">
            <button class="right uk-button uk-button-default uk-margin-top" type="submit">Save Changes</button>
            {nav.link("/c/" + coll, "Cancel", "right uk-button uk-button-default uk-margin-top uk-margin-right", undefined, true)}
          </div>
        </div>
      </fieldset>
    </form>
  }

  export function renderDetails(r: request.Request) {
    return <li class="request-details-panel">
      <div class="uk-margin-top">
        <label class="uk-form-label" for={r.key + "-key"}>Key</label>
        <input class="uk-input" id={r.key + "-key"} name="key" type="text" value={ r.key || "" } />
      </div>

      <div class="uk-margin-top">
        <label class="uk-form-label" for={r.key + "-title"}>Title</label>
        <input class="uk-input" id={r.key + "-title"} name="title" type="text" value={ r.title || "" } />
      </div>

      <div class="uk-margin-top">
        <label class="uk-form-label" for={r.key + "-description"}>Description</label>
        <input class="uk-input" id={r.key + "-description"} name="description" type="text" value={ r.description || "" } />
      </div>
    </li>;
  }
}
