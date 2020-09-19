namespace request.form {
  export function renderPrototype(key: string, p: request.Prototype) {
    return <div>
      <div class="uk-margin-top">
        <label class="uk-form-label" for={key + "-method"}>Method</label>
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

      <div class="uk-margin-top">
        <label class="uk-form-label" for={key + "-url"}>URL</label>
        <input class="uk-input" id={key + "-url"} name="url" type="text" value={request.prototypeToURL(p)} />
      </div>

      {renderAuthFields(key, p.auth)}
      {renderQueryParams(key, p.query)}
      {renderHeaders(key, p.headers)}
      {renderBody(key, p.body)}
      {renderOptions(key, p.options)}
    </div>;
  }

  function renderAuthFields(key: string, auth: auth.Auth[] | undefined) {
    return <div class="uk-margin-top">
      <label class="uk-form-label" for={key + "-auth"}>Auth</label>
      <textarea class="uk-textarea" id={"key-auth"} name="auth">{auth ? JSON.stringify(auth, null, 2) : "null"}</textarea>
    </div>;
  }

  function renderQueryParams(key: string, query: request.QueryParam[] | undefined) {
    return <div class="uk-margin-top">
      <label class="uk-form-label" for={key + "-queryparams"}>Query Params</label>
      <textarea class="uk-textarea" id={key + "-queryparams"} name="queryparams">{query ? JSON.stringify(query, null, 2) : "null"}</textarea>
    </div>;
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
}
