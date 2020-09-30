namespace call {
  export function renderResult(r: Result) {
    let rspDetail = <div>no result</div>

    if (r.response) {
      const ct = r.response.contentType || "";
      const cl = (r.response.contentLength && r.response.contentLength > -1) ? `(${r.response.contentLength} bytes)` : "";
      rspDetail = <div>{r.response.proto} <em>{r.response.status}</em><div>{ct} {cl}</div></div>;
    }

    return [
      <div class="right">
        <a class="theme uk-icon" data-uk-icon="close" href="" onclick="nav.pop();return false;" title="close result" />
      </div>,
      <div class="clear" />,
      <div>
        <ul data-uk-tab="">
          <li><a href="#result">Result</a></li>
          <li><a href="#request">Request</a></li>
          <li><a href="#headers">Response</a></li>
          <li><a href="#body">Body</a></li>
          <li><a href="#timing">Timing</a></li>
        </ul>
        <ul class="uk-switcher uk-margin">
          <li>
            <div>{r.status}: {(r.timing?.completed || 0) / 1000}ms</div>
            {rspDetail}
          </li>
          <li>
            <h3 class="uk-margin-small-bottom">{r.url}</h3>
            {renderHeaders("Final Request Headers", r.requestHeaders)}
          </li>
          <li>{renderHeaders("Response Headers", r.response?.headers)}</li>
          <li>{rbody.renderBody(r.url, r.response?.body)}</li>
          <li>{renderTiming(r.timing)}</li>
        </ul>
      </div>
    ];
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
