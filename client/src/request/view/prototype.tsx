namespace request.view {
  export function renderPrototype(p: Prototype) {
    return <div class="prototype">
      <div data-uk-grid="">
        <div class="uk-width-1-4">URL</div>
        <div class="uk-width-3-4"><div class="url">{request.prototypeToURL(p)}</div></div>
      </div>
      <hr/>
      <div data-uk-grid="">
        <div class="uk-width-1-4">Query Params</div>
        <div class="uk-width-3-4">{renderQueryParams(p.query)}</div>
      </div>
      <hr/>
      <div data-uk-grid="">
        <div class="uk-width-1-4">Auth</div>
        <div class="uk-width-3-4">{view.renderAuth(p.auth)}</div>
      </div>
      <hr/>
      <div data-uk-grid="">
        <div class="uk-width-1-4">Body</div>
        <div class="uk-width-3-4">{view.renderBody(p.body)}</div>
      </div>
      <hr/>
      <div data-uk-grid="">
        <div class="uk-width-1-4">Options</div>
        <div class="uk-width-3-4">{renderOptions(p.options)}</div>
      </div>
    </div>;
  }

  function renderQueryParams(query: request.QueryParam[] | undefined) {
    if (!query || query.length === 0) {
      return <em>no query params</em>;
    }
    return <div>
      {query.map(qp => <div data-uk-grid="">
        <div class="uk-width-1-4">{qp.k}</div>
        <div class="uk-width-3-4">{qp.v}</div>
      </div>)}
    </div>;
  }

  function renderOptions(o: request.Options | undefined) {
    if (!o || !(
      o.timeout || o.ignoreRedirects || o.ignoreReferrer || o.ignoreCerts ||
      o.excludeDefaultHeaders || o.readCookieJars || o.writeCookieJar || o.sslCert || o.userAgentOverride
    )) {
      return <em>no options</em>;
    }
    const section = function(title: string, v: any | undefined) {
      if(!v) {
        return <div />
      }
      return <div data-uk-grid="">
        <div class="uk-width-1-4">{title}</div>
        <div class="uk-width-3-4">{v}</div>
      </div>
    }

    return <div>
      {section("Timeout", o.timeout)}
      {section("ignoreRedirects", o.ignoreRedirects)}
      {section("ignoreReferrer", o.ignoreReferrer)}
      {section("ignoreCerts", o.ignoreCerts)}
      {section("excludeDefaultHeaders", o.excludeDefaultHeaders)}
      {section("readCookieJars", o.readCookieJars)}
      {section("writeCookieJar", o.writeCookieJar)}
      {section("sslCert", o.sslCert)}
      {section("userAgentOverride", o.userAgentOverride)}
    </div>;
  }
}
