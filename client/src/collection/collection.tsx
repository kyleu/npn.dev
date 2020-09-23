namespace collection {
  export function renderCollections(cs: Collection[]) {
    return cs.map(renderCollectionLink);
  }

  function renderCollectionLink(c: collection.Collection) {
    let title = c.title;
    if (!title || c.title.length === 0) {
      title = c.key;
    }
    return <div class="nav-item">{nav.link("/c/" + c.key, title)}</div>;
  }

  export function renderCollection(coll: collection.Collection, requests: request.Request[]) {
    const cn = coll.title ? coll.title : coll.key;
    return <div class="uk-card uk-card-body uk-card-default uk-margin-top">
      <div class="right">
        <a class="theme uk-icon" data-uk-icon="close" href="" onclick="nav.pop();return false;" title="close collection" />
      </div>
      <h3 class="uk-card-title">{cn}</h3>
      <div id="request-list" class="uk-margin-top">
        {request.view.renderRequests(cn, requests)}
      </div>
    </div>
  }
}
