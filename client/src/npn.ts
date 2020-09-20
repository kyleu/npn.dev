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

  export function debug() {
    const dump = function(k: string, v?: string) {
      console.warn(`${k}: ${v}`);
    }
    dump("Active Collection", collection.cache.active);
    dump("Active Request", request.cache.active);
    dump("Active Action", request.cache.action);
  }
}
