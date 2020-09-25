namespace request.editor {
  export function initAuthEditor(el: HTMLTextAreaElement) {

  }

  export function setAuth(cache: Cache, auth: auth.Auth[] | undefined) {
    const url = new URL(cache.url.value);
    let u = "";
    let p = "";
    if (auth) {
      for (let a of auth) {
        if (a.type === "basic") {
          const basic = a.config as auth.Basic
          u = encodeURIComponent(basic.username)
          p = encodeURIComponent(basic.password)
        }
      }
    }

    url.username = u;
    url.password = p;
    cache.url.value = url.toString();
  }

  export function updateBasicAuth(cache: Cache, auth: auth.Auth[] | undefined) {
    let currentAuth: auth.Auth[] = [];
    try {
      currentAuth = json.parse(cache.auth.value);
    } catch (e) {
      console.log("invalid auth JSON [" + cache.auth.value + "]")
    }
    let matched = -1;
    if (!currentAuth) {
      currentAuth = [];
    }
    for (let i = 0; i < currentAuth.length; i++) {
      const x = currentAuth[i];
      if (x.type === "basic") {
        matched = i;
      }
    }

    let basic: auth.Basic | undefined;
    if (auth) {
      for (let i = 0; i < auth.length; i++) {
        const x = auth[i];
        if (x.type === "basic") {
          basic = x.config as auth.Basic;
        }
      }
    }

    if (matched === -1) {
      if (basic) {
        currentAuth.push({type: "basic", config: basic});
      }
    } else {
      if (basic) {
        let curr = currentAuth[matched].config as auth.Basic;
        if (curr) {
          curr = {
            username: basic.username,
            password: basic.password,
            showPassword: curr.showPassword
          }
        } else {
          curr = basic
        }
        currentAuth[matched] = {type: "basic", config: curr};
      } else {
        currentAuth.splice(matched, 1);
      }
    }

    cache.auth.value = json.str(currentAuth);
  }
}
