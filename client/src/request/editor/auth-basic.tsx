namespace request.editor {
  export function updateBasicAuth(cache: Cache, auth: auth.Auth[] | undefined) {
    let currentAuth: auth.Auth[] = [];
    try {
      currentAuth = json.parse(cache.auth.value);
    } catch (e) {
      console.warn("invalid auth JSON [" + cache.auth.value + "]")
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
