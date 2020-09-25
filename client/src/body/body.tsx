namespace body {
  export function renderBody(b?: body.Body) {
    if (!b) {
      return <div>No body</div>;
    }
    switch (b.type) {
      case "json":
        return renderJSON(b.config as body.JSONConfig);
      case "html":
        return renderHTML(b.config as body.HTMLConfig);
      default:
        return <div>TODO: {b.type}</div>;
    }
  }

  function renderHTML(h: body.HTMLConfig) {
    return <div><em>HTML</em><pre style="overflow: auto; max-height: 720px;">{h.content}</pre></div>;
  }

  function renderJSON(j: body.JSONConfig) {
    return <div><em>JSON</em><pre>{json.str(j.msg)}</pre></div>;
  }
}
