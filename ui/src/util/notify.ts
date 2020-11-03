import UIkit from "uikit";

export function notify(msg: string, status: boolean): void {
  UIkit.notification(msg, { status: status ? "success" : "danger", pos: "top-right" });
}

export function confirm(msg: string, f: () => void): void {
  UIkit.modal.confirm(msg).then(f);
}

export function modal(key: string): UIkit.UIkitModalElement {
  const m = UIkit.modal(key);
  if (!m) {
    console.warn(`no modal available with key [${key}]`);
  }
  return m;
}
