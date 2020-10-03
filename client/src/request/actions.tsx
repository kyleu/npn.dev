namespace request {
  export function renderActionEmpty() {
    return <div/>;
  }

  export function renderActionUnknown(key: string | undefined, extra: string[]) {
    return <div>
      {renderActionClose()}
      unknown action: {key} ({extra})
    </div>;
  }

  export function renderActionCall(coll: string, req: string) {
    return <div id={coll + "--" + req + "-call"}>
      {renderActionClose()}
      <div class="call-title">Loading...</div>
      <div class="call-result" />
    </div>;
  }

  export function renderActionClose() {
    return <div class="right">
      <a class="theme uk-icon" data-uk-icon="close" href="" onclick="nav.navigate(`/c/${collection.cache.active}/${request.cache.active}`);return false;" title="close collection" />
    </div>;
  }
}
