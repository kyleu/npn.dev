declare const EmojiButton: any | undefined;

namespace style {
  export function setTheme(theme: string) {
    wireEmoji(theme);

    const card = dom.els(".uk-card");
    switch (theme) {
      case "auto":
        let t = "light";
        if (window.matchMedia && window.matchMedia("(prefers-color-scheme: dark)").matches) {
          t = "dark";
        }
        setTheme(t);
        fetch("/profile/theme/" + t).then(r => r.text()).then(() => {
          // console.log(`Set theme to [${t}]`);
        });
        break;
      case "light":
        document.documentElement.classList.remove("uk-light");
        document.body.classList.remove("uk-light");
        document.documentElement.classList.add("uk-dark");
        document.body.classList.add("uk-dark");
        card.forEach(x => {
          x.classList.add("uk-card-default");
          x.classList.remove("uk-card-secondary");
        });
        break;
      case "dark":
        document.documentElement.classList.add("uk-light");
        document.body.classList.add("uk-light");
        document.documentElement.classList.remove("uk-dark");
        document.body.classList.remove("uk-dark");
        card.forEach(x => {
          x.classList.remove("uk-card-default");
          x.classList.add("uk-card-secondary");
        });
        break;
      default:
        console.warn("invalid theme");
        break;
    }
  }

  export function themeLinks(color: string) {
    dom.els(".theme").forEach(el => {
      el.classList.add(`${color}-fg`);
    });
  }

  function wireEmoji(t: string) {
    if (typeof EmojiButton === "undefined") {
      dom.els(".picker-toggle").forEach(el => dom.setDisplay(el, false));
      return;
    }
    const opts = { position: "bottom-end", theme: t, zIndex: 1021 };
    dom.els(".textarea-emoji").forEach(el => {
      const toggle = dom.req(".picker-toggle", el);
      toggle.addEventListener(
        "click",
        () => {
          const textarea = dom.req<HTMLTextAreaElement>(".uk-textarea", el);
          const picker = new EmojiButton(opts);
          picker.on("emoji", (emoji: string) => {
            drop.onEmojiPicked();
            dom.insertAtCaret(textarea, emoji);
          });
          picker.togglePicker(toggle);
        },
        false
      );
    });
  }
}
