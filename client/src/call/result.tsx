namespace call {
  function renderResponse(rsp: undefined | call.Response, hash: string) {
    if (!rsp) {
      return <div>no response</div>;
    }
    const ct = rsp.contentType || "";
    const cl = (rsp.contentLength && rsp.contentLength > -1) ? `(${rsp.contentLength} bytes)` : ((rsp.body && rsp.body.length! > -1) ? `(${rsp.body.length} bytes)` : "");
    const ret = <div>
      <h3>{rsp ? rsp.status : "Unknown"}</h3>
      <em>{rsp.method} {rsp.url}</em>
      <div class="mt">
        <ul data-uk-tab="">
          {nav.hashLink("result", "Result", hash)}
          {nav.hashLink("request", "Request", hash)}
          {nav.hashLink("headers", "Response", hash)}
          {nav.hashLink("body", "Body", hash)}
          {nav.hashLink("timing", "Timing", hash)}
        </ul>
        <ul class="uk-switcher uk-margin">
          <li>
            <div>{(rsp.timing?.completed || 0) / 1000}ms</div>
            <div>{rsp.proto} <em>{rsp.status}</em><div>{ct} {cl}</div></div>
          </li>
          <li>
            {renderHeaders("Final Request Headers", rsp.requestHeaders)}
          </li>
          <li>{renderHeaders("Response Headers", rsp.headers)}</li>
          <li>{rbody.renderBody(rsp.url, rsp.body)}</li>
          <li>{renderTiming(rsp.timing)}</li>
        </ul>
      </div>
    </div>;
    if (rsp.prior) {
      return <div>
        {renderResponse(rsp.prior, hash)}
        <hr />
        {ret}
      </div>
    }
    return ret;
  }

  export function renderResult(r: Result, hash: string) {
    const ret = [
      <div class="right">
        <a class="theme uk-icon" data-uk-icon="close" href="" onclick="nav.pop();return false;" title="close result" />
      </div>,
      r.error ? <div>
        <div class="red-fg">error: {r.error}</div>
      </div> : <div />,
      renderResponse(r.response, hash)
    ];
    return ret;
  }

  function renderHeaders(title: string, headers: header.Header[] = []) {
    if (headers.length === 0) {
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
      <div class="result-timing-graph">
        <div>
          <div class="timing-start">0ms</div>
          <div class="timing-end">{t.completed / 1000}ms</div>
        </div>
        <object type="image/svg+xml" style={"width: 100%; height: " + (sections.length * 24) + "px"} data={timingGraph(sections)}>SVG not supported</object>
      </div>
      <hr />
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
