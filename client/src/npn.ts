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
      socket.setAppUnloading();
    };

    nav.init(function (p: string) {
      let parts = p.split("/");
      parts = parts.filter(x => x.length > 0);
      console.info("nav handler called, check it out: " + parts.join(" -> "));
      if (parts.length === 0) {
        return; // index
      }
      const svc = parts[0];
      switch (svc) {
        case "c":
          const collName = parts[1];
          if (collName !== collection.cache.active) {
            collection.cache.setActiveCollection(collName);
            socket.send({svc: services.collection.key, cmd: command.client.getCollection, param: collName});
          }
          if (parts.length > 2) {
            const reqName = parts[2];
            if (reqName !== request.cache.active) {
              request.cache.setActiveRequest(reqName);
            }
          }
          break;
        default:
          console.info("unhandled svc [" + svc + "]");
      }
    });

    socket.socketConnect(svc, id);
  }
}
