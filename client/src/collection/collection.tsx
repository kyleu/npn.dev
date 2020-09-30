namespace collection {
  export function renderCollections(cs: Collection[]) {
    return dom.els(".collection-list").forEach(el => {
      dom.setContent(el, cs.map(c => renderCollectionLink(c)));
    });
  }

  function renderCollectionLink(c: collection.Collection) {
    let title = c.title;
    if (!title || title.length === 0) {
      title = c.key;
    }
    let link = nav.link({path: "/c/" + c.key, title: title, icon: "folder"});
    if (cache.active === c.key) {
      const summs = request.cache.summaries.get(c.key);
      if (summs) {
        link = <div>
          {nav.link({path: "/", title: title, icon: "album"})}
          {summs.map(s => <div class="uk-margin-small-left">
            {nav.link({path: "/c/" + c.key + "/" + s.key, title: (s.title && s.title.length > 0) ? s.title : s.key, icon: "link"})}
          </div>)}
        </div>;
      }
    }
    return <div class={"nav-item collection-link collection-link-" + c.key}>{link}</div>;
  }

  export function renderCollection(coll: collection.Collection, requests: request.Summary[]) {
    const cn = coll.title ? coll.title : coll.key;
    return <div>
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right">
          <a class="theme uk-icon" data-uk-icon="close" href="" onclick="nav.pop();return false;" title="close collection"/>
        </div>
        <h3 class="uk-card-title"><span class="nav-icon-h3" data-uk-icon="icon: album"/>{cn}</h3>
        <p>{coll.description || ""}</p>
      </div>
      <div class="uk-card uk-card-body uk-card-default uk-margin-top">
        <h3 class="uk-card-title">Requests</h3>
        <form onsubmit="collection.addRequestURL();return false;">
          <div class="uk-margin-top uk-inline uk-width-expand">
            <button class="uk-form-icon uk-form-icon-flip" type="submit" title="add request" uk-icon="icon: plus"/>
            <input id="coll-request-add-url" class="uk-input" placeholder="add a request by url" data-lpignore="true"/>
          </div>
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
      const param = {"coll": collection.cache.active, "url": url};
      socket.send({svc: services.collection.key, cmd: command.client.addURL, param: param});
    }
    console.log("adding request [" + url + "]");
  }

  function renderRequests(coll: string, rs: request.Summary[]) {
    rs = group.sort(rs, x => x.order);
    return <ul class="uk-list uk-list-divider">
      {rs.map(r => renderRequestLink(coll, r))}
    </ul>
  }

  function renderRequestLink(coll: string, r: request.Summary) {
    let title = r.title;
    if (!title || r.title.length === 0) {
      title = r.key;
    }
    return <li>
      {nav.link({path: "/c/" + coll + "/" + r.key, title: title})}
      {r.description && r.description.length ? <div><em>{r.description}</em></div> : <span/>}
    </li>;
  }
}
