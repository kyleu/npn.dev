namespace flash {
  export function wire() {
    setTimeout(fadeOut, 4000);
  }

  function fadeOut() {
    let matched = false;
    dom.els(".alert-top").forEach(el => {
      matched = true
      el.classList.add("uk-animation-fade", "uk-animation-reverse");
    });
    if (matched) {
      setTimeout(remove, 1000);
    }
  }

  function remove() {
    dom.els(".alert-top").forEach(el => {
      el.remove();
    });
  }
}
