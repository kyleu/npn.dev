namespace request {
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
    return <div>
      {nav.link("/c/" + coll + "/" + r.key, title)}
    </div>;
  }

  export function renderRequest(coll: string, r: request.Request) {
    return renderPrototype(r.prototype);
  }

  function renderPrototype(p: Prototype) {
    return <div>{request.prototypeToURL(p)}</div>;
  }
}
