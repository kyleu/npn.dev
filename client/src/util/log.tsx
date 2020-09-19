namespace log {
  let started = 0;
  let container: HTMLElement | undefined;

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
    if (!container) {
      container = dom.req("#log-panel");
    }
    container.appendChild(el);
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
