namespace request.editor {
  export function updateBasicAuth(cache: Cache, auth: auth.Auth | undefined) {
    dom.setValue(cache.auth, json.str(auth));
  }
}
