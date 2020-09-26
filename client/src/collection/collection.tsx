namespace collection {
  export function renderCollections(cs: Collection[]) {
    return cs.map(renderCollectionLink);
  }

  function renderCollectionLink(c: collection.Collection) {
    let title = c.title;
    if (!title || c.title.length === 0) {
      title = c.key;
    }
    return <div class="nav-item">{nav.link("/c/" + c.key, title, "", "", false, "folder")}</div>;
  }

  export function renderCollection(coll: collection.Collection, requests: request.Request[]) {
    const cn = coll.title ? coll.title : coll.key;
    return <div>
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right">
          <a class="theme uk-icon" data-uk-icon="close" href="" onclick="nav.pop();return false;" title="close collection" />
        </div>
        <h3 class="uk-card-title"><span class="nav-icon-h3" data-uk-icon="icon: folder" />{cn}</h3>
        <p>{coll.description || ""}</p>
      </div>
      <div class="uk-card uk-card-body uk-card-default uk-margin-top">
        <h3 class="uk-card-title">Requests</h3>
        <form onsubmit="collection.addRequestURL();return false;">
          <input id="coll-request-add-url" class="uk-input" placeholder="add a request by url" />
        </form>
        <div id="request-list" class="uk-margin-top">
          {renderRequests(coll.key, requests)}
        </div>
      </div>
    </div>
  }

  export function addRequestURL() {
    const input = dom.req<HTMLInputElement>("#coll-request-add-url");
    const url = input.value.trim();
    if (url && url.length > 0) {
      input.value = "";
      console.log(url);
    }
  }

  function renderRequests(coll: string, rs: request.Request[]) {
    return <ul class="uk-list uk-list-divider">
      {rs.map(r => renderRequestLink(coll, r))}
    </ul>
  }

  function renderRequestLink(coll: string, r: request.Request) {
    let title = r.title;
    if (!title || r.title.length === 0) {
      title = r.key;
    }
    return <li>
      {nav.link("/c/" + coll + "/" + r.key, title)}
      {r.description && r.description.length ? <div><em>{r.description}</em></div> : <span />}
    </li>;
  }
}
