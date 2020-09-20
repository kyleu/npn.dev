namespace request.form {
  export function renderForm(coll: string, r: request.Request) {
    return <form class="uk-form-stacked">
      <input type="hidden" name="coll" value={coll} />
      <input type="hidden" name="originalKey" value={r.key} />
      <fieldset class="uk-fieldset">
        <legend class="hidden">request form</legend>

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

        {renderPrototype(r.key, r.prototype)}

        <div class="uk-margin-top">
          <button class="right uk-button uk-button-default uk-margin-top" type="submit">Save Changes</button>
          {nav.link("/c/" + coll + "/" + r.key, "Cancel", "right uk-button uk-button-default uk-margin-top uk-margin-right")}
        </div>
      </fieldset>
    </form>;
  }
}