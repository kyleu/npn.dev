namespace notify {
  export function notify(msg: string, status: boolean | string) {
    UIkit.notification(msg, { status: status ? "success" : "danger", pos: "top-right" });
  }
}
