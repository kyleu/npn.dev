namespace request {
  export function renderSummaryPanel(coll: string, r: request.Summary) {
    return <div>
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right">
          <a class="theme uk-icon" data-uk-icon="close" href="" onclick={"nav.navigate('/c/" + coll + "');return false;"} title="close request" />
        </div>
        <h3 class="uk-card-title">{r.title ? r.title : r.key}</h3>
        <p>Loading...</p>
      </div>
    </div>
  }
}
