namespace log {
  const started = Date.now()

  export function info(msg: string) {
    const el = l("info", msg);
    const container = dom.req("#log-panel");
    container.appendChild(el);
  }

  export function l(level: string, msg: string) {
    const n = Date.now() - started
    return <li>
      <div class="right">{n}ms</div>
      {msg}
    </li>;
  }
}
