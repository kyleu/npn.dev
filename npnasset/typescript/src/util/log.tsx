namespace log {
  let started = 0;
  let content: HTMLElement | undefined;
  let list: HTMLElement | undefined;

  export function init() {
    started = Date.now();
    l("debug", "npn started");
  }

  export function info(msg: string) {
    l("info", msg);
  }

  export function l(level: string, msg: string) {
    if (started === 0) {
      console.warn("call `log.init()` before attempting to log");
      return
    }
    const n = Date.now() - started
    const el = <li class={color(level)}>
      <div class="right">{n}ms</div>
      {msg}
    </li>;
    if (!list) {
      list = dom.opt("#log-list");
      if (!list) {
        console.warn(`${level}: ${msg}`);
        return;
      }
    }
    list.appendChild(el);
    if (!content) {
      content = dom.req<HTMLDivElement>("#log-content");
    }
    content.scrollTo(0, content.scrollHeight);
  }

  export function toggle() {
    const wsc = dom.req("#workspace-content");
    const lp = dom.req("#log-container");

    const curr = (lp.style.display !== "") && (lp.style.display !== "none");
    if (curr) {
      wsc.classList.remove("log-visible");
    } else {
      wsc.classList.add("log-visible");
    }
    dom.setDisplay(lp, !curr);

    if (!content) {
      content = dom.req<HTMLDivElement>("#log-content");
    }
    content.scrollTo(0, content.scrollHeight);
  }

  function color(level: string): string {
    switch (level) {
      case "debug":
        return "grey-fg";
      case "info":
        return "";
      case "warn":
        return "yellow-fg";
      case "error":
        return "red-fg";
      default:
        return "";
    }
  }
}
