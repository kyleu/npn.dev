namespace request.editor {
  export function initOptionsEditor(el: HTMLTextAreaElement) {
    const parent = el.parentElement!;
    parent.appendChild(createOptionsEditor(el));
  }

  function createOptionsEditor(el: HTMLTextAreaElement) {
    let opts = json.parse(el.value) as Options;

    if(!opts) {
      opts = {} as Options
    }

    return <div>
      <div class="uk-margin-top">
        <label class="uk-form-label" for={el.id + "-timeout"}>Timeout</label>
        <input class="uk-input" id={el.id + "-timeout"} name="opt-timeout" type="number" value={opts.timeout}/>
      </div>
      <div class="uk-margin-top">
        <label class="uk-form-label">Ignore</label>
        <div>
          {inputCheckbox(el.id, "ignoreRedirects", "Redirects", opts.ignoreRedirects || false)}
          {inputCheckbox(el.id, "ignoreReferrer", "Referrer", opts.ignoreReferrer || false)}
          {inputCheckbox(el.id, "ignoreCerts", "Certs", opts.ignoreCerts || false)}
          {inputCheckbox(el.id, "ignoreCookies", "Cookies", opts.ignoreCookies || false)}
        </div>
      </div>
      <div class="uk-margin-top">
        <label class="uk-form-label" for={el.id + "-excludeDefaultHeaders"}>Exclude Default Headers</label>
        <input class="uk-input" id={el.id + "-excludeDefaultHeaders"} name="opt-excludeDefaultHeaders" type="text" value={opts.excludeDefaultHeaders}/>
      </div>
      <div class="uk-margin-top">
        <label class="uk-form-label" for={el.id + "-readCookieJars"}>Read Cookie Jars</label>
        <input class="uk-input" id={el.id + "-readCookieJars"} name="opt-readCookieJars" type="text" value={opts.readCookieJars}/>
      </div>
      <div class="uk-margin-top">
        <label class="uk-form-label" for={el.id + "writeCookieJar"}>Write Cookie Jar</label>
        <input class="uk-input" id={el.id + "-writeCookieJar"} name="opt-writeCookieJar" type="text" value={opts.writeCookieJar}/>
      </div>
      <div class="uk-margin-top">
        <label class="uk-form-label" for={el.id + "-sslCert"}>SSL Cert</label>
        <input class="uk-input" id={el.id + "-sslCert"} name="opt-sslCert" type="text" value={opts.sslCert}/>
      </div>
      <div class="uk-margin-top">
        <label class="uk-form-label" for={el.id + "-userAgentOverride"}>User Agent Override</label>
        <input class="uk-input" id={el.id + "-userAgentOverride"} name="opt-userAgentOverride" type="text" value={opts.userAgentOverride}/>
      </div>
    </div>;
  }

  function inputCheckbox(key: string, prop: string, title: string, v: boolean) {
    const n = "opt-" + prop;
    const id = key + "-" + prop;
    if (v) {
      return <label class="uk-margin-right"><input type="checkbox" name={n} value="true" checked /> {title}</label>;
    } else {
      return <label class="uk-margin-right"><input type="checkbox" name={n} value="true" /> {title}</label>;
    }
  }
}
