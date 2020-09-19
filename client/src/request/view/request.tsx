namespace request.view {
  export function renderRequests(coll: string, rs: Request[]) {
    return <ul class="uk-list uk-list-divider">
      {rs.map(r => renderRequestLink(coll, r))}
    </ul>
  }

  export function renderRequestLink(coll: string, r: Request) {
    let title = r.title;
    if (!title || r.title.length === 0) {
      title = r.key;
    }
    return <li>{nav.link("/c/" + coll + "/" + r.key, title)}</li>;
  }

  export function renderRequestDetail(coll: string, r: request.Request) {
    const path = "/c/" + coll + "/" + r.key;
    return <div class="req" id={"req-" + r.key}>
      <div>
        <div data-uk-grid="">
          <div class="uk-width-1-4">Actions</div>
          <div class="uk-width-3-4">
            {nav.link(path + "/call", "Call", "uk-button uk-button-default uk-margin-right")}
            {nav.link(path + "/transform", "Transform", "uk-button uk-button-default uk-margin-right")}
            {nav.link(path + "/edit", "Edit", "uk-button uk-button-default uk-margin-right")}
            {nav.link(path + "/delete", "Delete", "uk-button uk-button-default uk-margin-right", "if (!confirm('Are you sure you want to delete request [" + r.key + "]?')) { return false; }")}
          </div>
        </div>
        <hr/>
        <div data-uk-grid="">
          <div class="uk-width-1-4">Key</div>
          <div class="uk-width-3-4">{r.key}</div>
        </div>
        <hr/>
        <div data-uk-grid="">
          <div class="uk-width-1-4">Title</div>
          <div class="uk-width-3-4">{r.title || ""}</div>
        </div>
        <hr/>
        <div data-uk-grid="">
          <div class="uk-width-1-4">Description</div>
          <div class="uk-width-3-4">{r.description || ""}</div>
        </div>
        <hr/>
      </div>
      {renderPrototype(r.prototype)}
    </div>
  }
}
