namespace drop {
  export function wire() {
    dom.els(".drop").forEach(el => {
      el.addEventListener("show", onDropOpen);
      el.addEventListener("beforehide", onDropBeforeHide);
      el.addEventListener("hide", onDropHide);
    });
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
