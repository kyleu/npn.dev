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
    window.onbeforeunload = () => {
      // socket.setAppUnloading();
    };

    // socket.socketConnect(services.fromKey(svc), id);
  }
}
