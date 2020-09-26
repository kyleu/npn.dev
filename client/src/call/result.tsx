namespace call {
  export function renderResult(r: Result) {
    return [
      <div class="right">
        <a class="theme uk-icon" data-uk-icon="close" href="" onclick="nav.pop();return false;" title="close result" />
      </div>,
      <div class="clear" />,
      <div>
        <ul data-uk-tab="">
          <li><a href="#result">Result</a></li>
          <li><a href="#headers">Response Headers</a></li>
          <li><a href="#body">Body</a></li>
          <li><a href="#request">Request Headers</a></li>
          <li><a href="#timing">Timing</a></li>
        </ul>
        <ul class="uk-switcher uk-margin">
          <li>
            <div>{r.status}: {(r.timing?.completed || 0) / 1000}ms</div>
            {r.response?.proto || ""} {`${r.response?.contentType || ""} (${r.response?.contentLength || "no"} bytes)`}
            <hr />
            <div class="result-timing-graph uk-inline"/>
          </li>
          <li>{renderHeaders("Response Headers", r.response?.headers)}</li>
          <li>{body.renderBody(r.response?.body)}</li>
          <li>{renderHeaders("Final Request Headers", r.requestHeaders)}</li>
          <li>{renderTiming(r.timing)}</li>
        </ul>
      </div>
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
