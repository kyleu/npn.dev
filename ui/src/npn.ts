export function setTheme(theme: string): void {
  let t = "";
  switch (theme) {
    case "auto":
      t = "light";
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
      break;
    case "dark":
      document.documentElement.classList.add("uk-light");
      document.body.classList.add("uk-light");
      document.documentElement.classList.remove("uk-dark");
      document.body.classList.remove("uk-dark");
      break;
    default:
      console.warn("invalid theme");
      break;
  }
}

function themeLinks(color: string): void {
  const linkColor = `${color}-fg`;
  document.querySelectorAll(".theme").forEach(el => {
    el.classList.add(linkColor);
  });
}

export function initDom(t: string, color: string): void {
  try {
    themeLinks(color);
    setTheme(t);
  } catch (e) {
    console.warn("error setting style", e);
  }
}
