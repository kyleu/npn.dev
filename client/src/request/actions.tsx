namespace request {
  export function renderActionEmpty(r: request.Request) {
    return <div/>;
  }

  export function renderActionUnknown(key: string | undefined, extra: string[], r: request.Request) {
    return <div>
      {renderClose(r)}
      unknown action: {key} ({extra})
    </div>;
  }

  export function renderActionCall(coll: string, r: request.Request) {
    return <div id={coll + "--" + r.key + "-call"}>
      {renderClose(r)}
      <div class="call-title">Loading...</div>
      <div class="call-result"/>
    </div>;
  }

  function renderClose(r: request.Request) {
    return <div class="right">
      <a class="theme uk-icon" data-uk-icon="close" href="" onclick="nav.navigate(`/c/${collection.cache.active}/${request.cache.active}`);return false;" title="close collection" />
    </div>;
  }
}
