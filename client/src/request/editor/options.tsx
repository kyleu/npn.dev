namespace request.editor {
  export function initOptionsEditor(el: HTMLTextAreaElement) {
    const parent = el.parentElement!;
    parent.appendChild(createOptionsEditor(el));
  }

  function inputBool(key: string, v: boolean) {
    if (v) {
      return <div>
        <label class="uk-margin-small-right"><input class="uk-radio" type="radio" name={key} value="true" checked/> True</label>
        <label><input class="uk-radio" type="radio" name={key} value="false"/> False</label>
      </div>
    } else {
      return <div>
        <label class="uk-margin-small-right"><input class="uk-radio" type="radio" name={key} value="true"/> True</label>
        <label><input class="uk-radio" type="radio" name={key} value="false" checked/> False</label>
      </div>
    }
  }

  function createOptionsEditor(el: HTMLTextAreaElement) {
    let opts = JSON.parse(el.value) as Options;

    if(!opts) {
      opts = {} as Options
    }

    return <div>
      <div class="uk-margin-top">
        <label class="uk-form-label" for={el.id + "-timeout"}>Timeout</label>
        <input class="uk-input" id={el.id + "-timeout"} name="opt-timeout" type="number" value={opts.timeout}/>
      </div>
      <div class="uk-margin-top">
        <label class="uk-form-label" for={el.id + "-ignoreRedirects"}>Ignore Redirects</label>
        {inputBool(el.id + "-ignoreRedirects", opts.ignoreRedirects || false)}
      </div>
      <div class="uk-margin-top">
        <label class="uk-form-label" for="<%= key %>-opt-ignoreReferrer">Ignore Referrer</label>
        {inputBool(el.id + "ignoreReferrer", opts.ignoreReferrer || false)}
      </div>
      <div class="uk-margin-top">
        <label class="uk-form-label" for="<%= key %>-opt-ignoreCerts">Ignore Certs</label>
        {inputBool(el.id + "ignoreCerts", opts.ignoreCerts || false)}
      </div>
      <div class="uk-margin-top">
        <label class="uk-form-label" for="<%= key %>-opt-ignoreCookies">Ignore Cookies</label>
        {inputBool(el.id + "ignoreCookies", opts.ignoreCookies || false)}
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
}
