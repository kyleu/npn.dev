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
    if (inIframe()) {
      document.body.innerHTML = "";
      document.body.appendChild(rbody.iframeError());
      return;
    }

    log.init();

    window.onbeforeunload = () => {
      socket.setAppUnloading();
    };

    nav.init(routing.route);
    socket.socketConnect(svc, id, svc === "wasm");
  }

  export function debug() {
    const dump = (k: string, v: string = "") => {
      console.warn(`${k}: ${v}`);
    }
    dump("Active Collection", collection.cache.active);
    dump("Active Request", request.cache.active);
    dump("Active Action", `${request.cache.action} [${request.cache.extra}]`);
  }

  export function testbed() {
    log.info("Testbed!");
  }

  function inIframe () {
    try {
      return window.self !== window.top;
    } catch (e) {
      return true;
    }
  }
}
