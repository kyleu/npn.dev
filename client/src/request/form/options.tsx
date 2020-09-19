namespace request.form {
  export function renderOptions(key: string, opts: request.Options | undefined) {
    if (!opts) {
      opts = {};
    }
    return <div>
      <div class="uk-margin-top">
        <label class="uk-form-label" for={key + "-opt-timeout"}>Timeout</label>
        <input class="uk-input" id={key + "-opt-timeout"} name="opt-timeout" type="number" value={opts.timeout} />
      </div>
      <div class="uk-margin-top">
        <label class="uk-form-label" for={key + "-opt-ignoreRedirects"}>Ignore Redirects</label>
        {inputBool(key + "-opt-ignoreRedirects", opts.ignoreRedirects)}
      </div>
      <div class="uk-margin-top">
        <label class="uk-form-label" for={key + "-opt-ignoreReferrer"}>Ignore Referrer</label>
        {inputBool(key + "-opt-ignoreReferrer", opts.ignoreReferrer)}
      </div>
      <div class="uk-margin-top">
        <label class="uk-form-label" for={key + "-opt-ignoreCerts"}>Ignore Certs</label>
        {inputBool(key + "-opt-ignoreCerts", opts.ignoreCerts)}
      </div>
      <div class="uk-margin-top">
        <label class="uk-form-label" for={key + "-opt-excludeDefaultHeaders"}>Exclude Default Headers</label>
        <input class="uk-input" id={key + "-opt-excludeDefaultHeaders"} name="opt-excludeDefaultHeaders" type="text" value={opts.excludeDefaultHeaders} />
      </div>
      <div class="uk-margin-top">
        <label class="uk-form-label" for={key + "-opt-readCookieJars"}>Read Cookie Jars</label>
        <input class="uk-input" id={key + "-opt-readCookieJars"} name="opt-readCookieJars" type="text" value={opts.readCookieJars} />
      </div>
      <div class="uk-margin-top">
        <label class="uk-form-label" for={key + "-opt-writeCookieJar"}>Write Cookie Jar</label>
        <input class="uk-input" id={key + "-opt-writeCookieJar"} name="opt-writeCookieJar" type="text" value={opts.writeCookieJar} />
      </div>
      <div class="uk-margin-top">
        <label class="uk-form-label" for={key + "-opt-sslCert"}>SSL Cert</label>
        <input class="uk-input" id={key + "-opt-sslCert"} name="opt-sslCert" type="text" value={opts.sslCert} />
      </div>
      <div class="uk-margin-top">
        <label class="uk-form-label" for={key + "-opt-userAgentOverride"}>User Agent Override</label>
        <input class="uk-input" id={key + "-opt-userAgentOverride"} name="opt-userAgentOverride" type="text" value={opts.userAgentOverride} />
      </div>
    </div>;
  }

  function inputBool(name: string, v: boolean | undefined) {
    if (v) {
      return <div>
        <label><input class="uk-radio" type="radio" name={name} value="true" checked="checked" /> True</label>
        <label class="uk-margin-left"><input class="uk-radio" type="radio" name={name} value="false" /> False</label>
      </div>;
    } else {
      return <div>
        <label><input class="uk-radio" type="radio" name={name} value="true" /> True</label>
        <label class="uk-margin-left"><input class="uk-radio" type="radio" name={name} value="false" checked="checked" /> False</label>
      </div>;
    }
  }
}
