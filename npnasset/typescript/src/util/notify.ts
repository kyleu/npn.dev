namespace notify {
  export function notify(msg: string, status: boolean) {
    UIkit.notification(msg, { status: status ? "success" : "danger", pos: "top-right" });
  }

  export function confirm(msg: string, f: () => void) {
    UIkit.modal.confirm(msg).then(f);
  }

  export function modal(key: string) {
    const m = UIkit.modal(key);
    if (!m) {
      console.warn(`no modal available with key [${key}]`);
    }
    return m;
  }
}
