namespace request.view {
  export function renderAuth(auth: auth.Auth[] | undefined) {
    if (!auth || auth.length === 0) {
      return <em>no authentication</em>;
    }
    return <div>
      {auth.map(a => <div data-uk-grid="">
        <div class="uk-width-1-4">{a.type}</div>
        <div class="uk-width-3-4"><pre>{json.str(a.config)}</pre></div>
      </div>)}
    </div>;
  }
}
