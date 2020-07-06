namespace tags {
  export function wire() {
    dom.els(".tag-editor").forEach(el => {
      el.addEventListener("moved", onTagEditorUpdate);
      el.addEventListener("added", onTagEditorUpdate);
      el.addEventListener("removed", onTagEditorUpdate);
    });
  }

  export function removeTag(el: HTMLElement) {
    const itemEl = el.parentElement!;
    const editorEl = itemEl.parentElement!;
    itemEl.remove();
    updateEditor(editorEl);
  }

  export function addTag(el: HTMLElement) {
    const editorEl = el.parentElement;
    if(!editorEl) {
      return;
    }
    const itemEl = renderItem();
    editorEl.insertBefore(itemEl, dom.req(".add-item", editorEl));
    editTag(itemEl);
  }

  export function editTag(el: HTMLElement) {
    const valueEl = dom.req(".value", el);
    const editorEl = dom.req(".editor", el);
    dom.setDisplay(valueEl, false);
    dom.setDisplay(editorEl, true);
    const input = renderInput(valueEl.innerText) as HTMLInputElement;
    input.onblur = function() {
      valueEl.innerText = input.value;
      dom.setDisplay(valueEl, true);
      dom.setDisplay(editorEl, false);
      updateEditor(el.parentElement!);
    };
    input.onkeypress = function(e) {
      if (e.key === "Enter") {
        input.blur();
        return false;
      }
      return true;
    };
    dom.setContent(editorEl, input);
    input.focus();
  }

  function onTagEditorUpdate(e: Event) {
    if (!e.target) {
      console.warn("no event target");
      return;
    }
    const el = e.target as HTMLElement;
    updateEditor(el);
  }

  function updateEditor(el: HTMLElement) {
    const key = el.dataset["key"] || "";
    const f = events.getOpenEvent(key);
    if (f) {
      f();
    } else {
      console.warn(`no tag open handler registered for [${key}]`);
    }
    const ret = dom.els(".item", el).map(el => el.innerText);
    dom.setValue(`#model-${key}-input`, ret.join(","));
  }
}
