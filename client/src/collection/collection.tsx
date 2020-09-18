namespace collection {
  export function renderCollections(cs: Collection[]) {
    return <ul class="uk-list uk-list-divider">
      {cs.map(renderCollection)}
    </ul>
  }

  export function renderCollection(c: collection.Collection) {
    let title = c.title;
    if (!title || c.title.length === 0) {
      title = c.key;
    }
    return <li>{nav.link("/c/" + c.key, title)}</li>;
  }
}
