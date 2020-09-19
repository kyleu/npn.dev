namespace request.view {
  export function renderBody(b: body.Body | undefined) {
    if (!b) {
      return <em>no body</em>;
    }
    return <div data-uk-grid="">
      <div class="uk-width-1-4">{b?.type || "?"}</div>
      <div class="uk-width-3-4"><pre>{JSON.stringify(b.config, null, 2)}</pre></div>
    </div>;
  }
}
