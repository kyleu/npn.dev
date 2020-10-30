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
      <div class="mt">
        <label class="uk-form-label" for={el.id + "-timeout"}>Timeout</label>
        {timeoutInput(el, opts.timeout)}
      </div>
      <div class="mt">
        <label class="uk-form-label">Ignore</label>
        <div>
          {ignoreRedirectsInput(el, opts.ignoreRedirects)}
          {ignoreReferrerInput(el, opts.ignoreReferrer)}
          {ignoreCertsInput(el, opts.ignoreCerts)}
          {ignoreCookiesInput(el, opts.ignoreCookies)}
        </div>
      </div>
      <div class="mt">
        <label class="uk-form-label" for={el.id + "-excludeDefaultHeaders"}>Exclude Default Headers</label>
        {excludeDefaultHeadersInput(el, opts.excludeDefaultHeaders)}
      </div>
      <div class="mt">
        <label class="uk-form-label" for={el.id + "-readCookieJars"}>Read Cookie Jars</label>
        {readCookieJarsInput(el, opts.readCookieJars)}
      </div>
      <div class="mt">
        <label class="uk-form-label" for={el.id + "writeCookieJar"}>Write Cookie Jar</label>
        {writeCookieJarInput(el, opts.writeCookieJar)}
      </div>
      <div class="mt">
        <label class="uk-form-label" for={el.id + "-sslCert"}>SSL Cert</label>
        {sslCertInput(el, opts.sslCert)}
      </div>
      <div class="mt">
        <label class="uk-form-label" for={el.id + "-userAgentOverride"}>User Agent Override</label>
        {userAgentOverrideInput(el, opts.userAgentOverride)}
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

  function wire(ret: HTMLElement, f: (opts: request.Options) => void, el: HTMLTextAreaElement) {
    events(ret, () => {
      let opts = json.parse(el.value) as Options;
      f(opts);
      dom.setValue(el, json.str(opts));
      check();
    });
  }

  function timeoutInput(el: HTMLTextAreaElement, timeout: number | undefined) {
    const ret = <input class="uk-input" id={el.id + "-timeout"} name="opt-timeout" type="number" value={timeout || 0}/> as HTMLInputElement;
    wire(ret, (opts: Options) => opts.timeout = parseInt(ret.value, 10), el);
    return ret;
  }

  function ignoreRedirectsInput(el: HTMLTextAreaElement, ignoreRedirects: boolean | undefined) {
    const ret = inputCheckbox(el.id, "ignoreRedirects", "Redirects", ignoreRedirects || false);
    wire(ret, (opts: Options) => opts.ignoreRedirects = dom.req<HTMLInputElement>("input", ret).checked, el);
    return ret;
  }

  function ignoreReferrerInput(el: HTMLTextAreaElement, ignoreReferrer: boolean | undefined) {
    const ret = inputCheckbox(el.id, "ignoreReferrer", "Referrer", ignoreReferrer || false)
    wire(ret, (opts: Options) => opts.ignoreReferrer = dom.req<HTMLInputElement>("input", ret).checked, el);
    return ret;
  }

  function ignoreCertsInput(el: HTMLTextAreaElement, ignoreCerts: boolean | undefined) {
    const ret = inputCheckbox(el.id, "ignoreCerts", "Certs", ignoreCerts || false);
    wire(ret, (opts: Options) => opts.ignoreCerts = dom.req<HTMLInputElement>("input", ret).checked, el);
    return ret;
  }

  function ignoreCookiesInput(el: HTMLTextAreaElement, ignoreCookies: boolean | undefined) {
    const ret = inputCheckbox(el.id, "ignoreCookies", "Cookies", ignoreCookies || false);
    wire(ret, (opts: Options) => opts.ignoreCookies = dom.req<HTMLInputElement>("input", ret).checked, el);
    return ret;
  }

  function excludeDefaultHeadersInput(el: HTMLTextAreaElement, excludeDefaultHeaders: string[] | undefined) {
    const ret = <input class="uk-input" id={el.id + "-excludeDefaultHeaders"} name="opt-excludeDefaultHeaders" type="text" value={excludeDefaultHeaders}/> as HTMLInputElement;
    wire(ret, (opts: Options) => opts.excludeDefaultHeaders = ret.value.split(",").map(x => x.trim()), el);
    return ret;
  }

  function readCookieJarsInput(el: HTMLTextAreaElement, readCookieJars: string[] | undefined) {
    const ret = <input class="uk-input" id={el.id + "-readCookieJars"} name="opt-readCookieJars" type="text" value={readCookieJars}/> as HTMLInputElement;
    wire(ret, (opts: Options) => opts.readCookieJars = ret.value.split(",").map(x => x.trim()), el);
    return ret;
  }

  function writeCookieJarInput(el: HTMLTextAreaElement, writeCookieJar: string | undefined) {
    const ret = <input class="uk-input" id={el.id + "-writeCookieJar"} name="opt-writeCookieJar" type="text" value={writeCookieJar}/> as HTMLInputElement;
    wire(ret, (opts: Options) => opts.writeCookieJar = ret.value, el);
    return ret;
  }

  function sslCertInput(el: HTMLTextAreaElement, sslCert: string | undefined) {
    const ret = <input class="uk-input" id={el.id + "-sslCert"} name="opt-sslCert" type="text" value={sslCert}/> as HTMLInputElement;
    wire(ret, (opts: Options) => opts.sslCert = ret.value, el);
    return ret;
  }

  function userAgentOverrideInput(el: HTMLTextAreaElement, userAgentOverride: string | undefined) {
    const ret = <input class="uk-input" id={el.id + "-userAgentOverride"} name="opt-userAgentOverride" type="text" value={userAgentOverride || ""}/> as HTMLInputElement;
    wire(ret, (opts: Options) => opts.userAgentOverride = ret.value, el);
    return ret;
  }
}
