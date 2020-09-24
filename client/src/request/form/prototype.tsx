namespace request.form {
  export function renderSwitcher(r: request.Request) {
    const key = r.key;
    const p = r.prototype;
    return <div>
      <ul data-uk-tab="">
        <li><a href="#">Details</a></li>
        <li><a href="#">Query</a></li>
        <li><a href="#">Auth</a></li>
        <li><a href="#">Headers</a></li>
        <li><a href="#">Body</a></li>
        <li><a href="#">Options</a></li>
      </ul>
      <ul class="uk-switcher uk-margin">
        {renderDetails(r)}
        {renderQueryParams(key, p.query)}
        {renderAuth(key, p.auth)}
        {renderHeaders(key, p.headers)}
        {renderBody(key, p.body)}
        {renderOptions(key, p.options)}
      </ul>
    </div>
  }

  function renderQueryParams(key: String, qp: request.QueryParam[] | undefined) {
    return <li class="request-queryparams-panel">
      <div class="uk-margin-top">
        <label class="uk-form-label" for={key + "-queryparams"}>Query Params</label>
        <textarea class="uk-textarea" id={key + "-queryparams"} name="queryparams">{json.str(qp)}</textarea>
      </div>
    </li>;
  }

  function renderAuth(key: String, as: auth.Auth[] | undefined) {
    return <li class="request-auth-panel">
      <div class="uk-margin-top">
        <label class="uk-form-label" for={key + "-auth"}>Auth</label>
        <textarea class="uk-textarea" id={key + "-auth"} name="auth">{json.str(as)}</textarea>
      </div>
    </li>;
  }

  function renderHeaders(key: String, hs: request.Header[] | undefined) {
    return <li class="request-headers-panel">
      <div class="uk-margin-top">
        <label class="uk-form-label" for={key + "-headers"}>Headers</label>
        <textarea class="uk-textarea" id={key + "-headers"} name="headers">{json.str(hs)}</textarea>
      </div>
    </li>;
  }

  function renderBody(key: String, b: body.Body | undefined) {
    return <li class="request-body-panel">
      <div class="uk-margin-top">
        <label class="uk-form-label" for={key + "-body"}>Body</label>
        <textarea class="uk-textarea" id={key + "-body"} name="body">{json.str(b)}</textarea>
      </div>
    </li>;
  }

  function renderOptions(key: String, opts: request.Options | undefined) {
    return <li class="request-options-panel">
      <div class="uk-margin-top">
        <label class="uk-form-label" for={key + "-options"}>Options</label>
        <textarea class="uk-textarea" id={key + "-options"} name="options">{json.str(opts)}</textarea>
      </div>
    </li>;
  }
}
