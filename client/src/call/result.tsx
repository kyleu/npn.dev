namespace call {
  export function renderResult(r: Result) {
    const statusEl = <div>{r.status}: {(r.timing?.completed || 0) / 1000}ms</div>;
    return [
      <div class="right">
        <a class="theme uk-icon" data-uk-icon="close" href="" onclick="nav.pop();return false;" title="close result" />
      </div>,
      section("Result", statusEl),
      <hr/>,
      <div>{renderHeaders("Final Request Headers", r.requestHeaders)}</div>,
      <hr/>,
      ...renderResponse(r.response),
      <hr/>,
      section("Timing", renderTiming(r.timing))
    ];
  }

  function renderResponse(r?: Response) {
    if (!r) {
      return [<div>No response</div>];
    }
    return [
      section("Status", r.status),
      <hr/>,
      section("Protocol", r.proto),
      <hr/>,
      <div>{renderHeaders("Response Headers", r.headers)}</div>,
      <hr/>,
      section("Content", `${r.contentType} (${r.contentLength} bytes)`),
      <hr/>,
      section("Body", body.renderBody(r.body)),
    ];
  }

  function renderHeaders(title: string, headers?: header.Header[]) {
    if (!headers) {
      return section(title, "No headers");
    }
    return <div class="uk-overflow-auto">
      <h4>{title}</h4>
      <table class="uk-table uk-table-divider uk-text-left uk-table-small uk-table-justify">
        <tbody>
        {headers.map(h => <tr title={h.desc}><td class="uk-text-nowrap">{h.k}</td><td class="uk-text-nowrap">{h.v}</td></tr>)}
        </tbody>
      </table>
    </div>;
  }

  function renderTiming(t?: Timing) {
    if (!t) {
      return <div>No timing</div>;
    }

    const sections = timingSections(t);
    return <div class="timing-panel">
      {sections.map(sc => <div>{sc.key}: {sc.start} - {sc.end}</div>)}
    </div>;
  }

  function section(k: string, v: string | JSX.Element | undefined) {
    if (!v) {
      v = "undefined"
    }
    return <div><h4>{k}</h4> {v}</div>;
  }
}
