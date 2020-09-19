namespace npn {
  export function onError(svc: string, err: string) {
    console.error(`${svc}: ${err}`);
    const idx = err.lastIndexOf(":");
    if (idx > -1) {
      err = err.substr(idx + 1);
    }
    notify.notify(`${svc} error: ${err}`, false);
  }

  export function init(svc: string, id: string) {
    log.init();

    window.onbeforeunload = () => {
      socket.setAppUnloading();
    };

    nav.init(socket.route);

    socket.socketConnect(svc, id);
  }
}
