namespace dom {
  export function setValue(el: string | HTMLInputElement | HTMLTextAreaElement, text: string): HTMLInputElement | HTMLTextAreaElement {
    if (typeof el === "string") {
      el = req<HTMLInputElement>(el);
    }
    el.value = text;
    return el;
  }

  export function wireTextarea(text: HTMLTextAreaElement) {
    function resize() {
      text.style.height = "auto";
      text.style.height = `${text.scrollHeight < 64 ? 64 : text.scrollHeight + 6}px`;
    }

    function delayedResize() {
      window.setTimeout(resize, 0);
    }

    const x = text.dataset["autoresize"];
    if (!x) {
      text.dataset["autoresize"] = "true";

      text.addEventListener("change", resize, false);
      text.addEventListener("cut", delayedResize, false);
      text.addEventListener("paste", delayedResize, false);
      text.addEventListener("drop", delayedResize, false);
      text.addEventListener("keydown", delayedResize, false);

      text.focus();
      text.select();
    }

    resize();
  }

  export function setOptions(el: string | HTMLSelectElement, categories: readonly string[]) {
    if (typeof el === "string") {
      el = req<HTMLSelectElement>(el);
    }
    dom.clear(el);
    categories.forEach(c => {
      const opt = document.createElement("option");
      opt.value = c;
      dom.setText(opt, c);
      (el as HTMLSelectElement).appendChild(opt);
    });
  }

  export function setSelectOption(el: string | HTMLSelectElement, o: string | undefined) {
    if (typeof el === "string") {
      el = req<HTMLSelectElement>(el);
    }
    for (let i = 0; i < el.children.length; i++) {
      const e = el.children.item(i) as HTMLOptionElement;
      e.selected = e.value === o;
    }
  }

  export function insertAtCaret(e: HTMLTextAreaElement, text: string) {
    if (e.selectionStart || e.selectionStart === 0) {
      let startPos = e.selectionStart;
      let endPos = e.selectionEnd;
      e.value = e.value.substring(0, startPos) + text + e.value.substring(endPos, e.value.length);
      e.selectionStart = startPos + text.length;
      e.selectionEnd = startPos + text.length;
    } else {
      e.value += text;
    }
  }
}
