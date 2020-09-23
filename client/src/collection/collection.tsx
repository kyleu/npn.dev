namespace collection {
  export function renderCollections(cs: Collection[]) {
    return cs.map(renderCollection);
  }

  export function renderCollection(c: collection.Collection) {
    let title = c.title;
    if (!title || c.title.length === 0) {
      title = c.key;
    }
    return <div class="nav-item">{nav.link("/c/" + c.key, title)}</div>;
  }
}
