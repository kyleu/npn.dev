namespace request.form {
  function renderAuthFields(key: string, auth: auth.Auth[] | undefined) {
  }

  function renderHeaders(key: string, headers: request.Header[] | undefined) {
    return <div class="uk-margin-top">
      <label class="uk-form-label" for={key + "-headers"}>Headers</label>
      <textarea class="uk-textarea" id={key + "-headers"} name="headers">{headers ? JSON.stringify(headers, null, 2) : "null"}</textarea>
    </div>;
  }

  function renderBody(key: string, body: body.Body | undefined) {
    return <div class="uk-margin-top">
      <label class="uk-form-label" for={key + "-body"}>Body</label>
      <textarea class="uk-textarea" id={key + "-body"} name="body">{body ? JSON.stringify(body, null, 2) : "null"}</textarea>
    </div>;
  }

  function renderOptions(key: string, opts: request.Options | undefined) {
    return <div class="uk-margin-top">
      <label class="uk-form-label" for={key + "-options"}>Options</label>
      <textarea class="uk-textarea" id={key + "-options"} name="options">{opts ? JSON.stringify(opts, null, 2) : "null"}</textarea>
    </div>;
  }
}
