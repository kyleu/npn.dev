namespace tags {
  export function renderInput(v: string) {
    return <input type="text" class="uk-input" value={v} />;
  }

  export function renderItem() {
    return <span class="item">
      <span class="value" onclick="tags.editTag(this.parentElement);"/>
      <span class="editor"/>
      <span class="close" data-uk-icon="icon: close; ratio: 0.6;" onclick="tags.removeTag(this);"/>
    </span>;
  }

  export function renderTagsView(a: readonly string[]) {
    return <div class="tag-view">
      {a.map(s => <span class="item">{s}</span>)}
      <div class="clear"/>
    </div>;
  }
}
