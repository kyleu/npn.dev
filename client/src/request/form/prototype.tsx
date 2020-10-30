namespace request.form {
  export function renderSwitcher(r: request.Request, hash: string) {
    const key = r.key;
    const p = r.prototype;
    return <div>
      <ul data-uk-tab="">
        {nav.hashLink("details", "Details", hash)}
        {nav.hashLink("query", "Query", hash)}
        {nav.hashLink("auth", "Auth", hash)}
        {nav.hashLink("headers", "Headers", hash)}
        {nav.hashLink("body", "Body", hash)}
        {nav.hashLink("options", "Options", hash)}
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
      <div class="mt">
        <textarea class="uk-textarea hidden" id={key + "-queryparams"} name="queryparams">{json.str(qp)}</textarea>
      </div>
    </li>;
  }

  function renderAuth(key: String, as: auth.Auth | undefined) {
    return <li class="request-auth-panel">
      <div class="mt">
        <textarea class="uk-textarea hidden" id={key + "-auth"} name="auth">{json.str(as)}</textarea>
      </div>
    </li>;
  }

  function renderHeaders(key: String, hs: header.Header[] | undefined) {
    return <li class="request-headers-panel">
      <div class="mt">
        <textarea class="uk-textarea hidden" id={key + "-headers"} name="headers">{json.str(hs)}</textarea>
      </div>
    </li>;
  }

  function renderBody(key: String, b: rbody.Body | undefined) {
    return <li class="request-body-panel">
      <div class="mt">
        <textarea class="uk-textarea hidden" id={key + "-body"} name="body">{json.str(b)}</textarea>
      </div>
    </li>;
  }

  function renderOptions(key: String, opts: request.Options | undefined) {
    return <li class="request-options-panel">
      <div class="mt">
        <textarea class="uk-textarea hidden" id={key + "-options"} name="options">{json.str(opts)}</textarea>
      </div>
    </li>;
  }
}
