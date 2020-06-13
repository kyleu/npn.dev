namespace drop {
  export function wire() {
    UIkit.util.on(".drop", "show", onDropOpen);
    UIkit.util.on(".drop", "beforehide", onDropBeforeHide);
    UIkit.util.on(".drop", "hide", onDropHide);

    events.register("export");
  }

  function onDropOpen(e: Event) {
    if (!e.target) {
      return;
    }
    const el = e.target as HTMLElement;
    const key = el.dataset["key"] || "";
    let t = el.dataset["t"] || "";
    const f = events.getOpenEvent(key);
    if (f) {
      f(t);
    } else {
      console.warn(`no drop open handler registered for [${key}]`);
    }
  }

  function onDropHide(e: Event) {
    if (!e.target) {
      return;
    }
    const el = e.target as HTMLElement;
    if (el.classList.contains("uk-open")) {
      const key = el.dataset["key"] || "";
      const t = el.dataset["t"] || "";
      const f = events.getCloseEvent(key);
      if (f) {
        f(t);
      }
    }
  }

  let emojiPicked = false;

  export function onEmojiPicked() {
    emojiPicked = true;
    setTimeout(() => (emojiPicked = false), 200);
  }

  function onDropBeforeHide(e: Event) {
    if (emojiPicked) {
      e.preventDefault();
    }
  }
}
