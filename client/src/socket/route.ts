namespace socket {
  export function routeOld(p: string) {
    let parts = p.split("/");
    parts = parts.filter(x => x.length > 0);
    console.info("nav (old): " + parts.join(" -> "));
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

          const action = parts[3];
          if (action !== request.cache.action) {
            request.cache.setActiveAction(action);
          }
        }
        break;
      default:
        console.info("unhandled svc [" + svc + "]");
    }
  }

  export function route(p: string) {
    let parts = p.split("/");
    parts = parts.filter(x => x.length > 0);
    console.info("nav: " + parts.join(" -> "));

    if (parts.length === 0 || parts[0].length === 0) {
      ui.setPanels();
      return; // index
    }
    const svc = parts[0];
    switch (svc) {
      case "c":
        let coll = (parts.length > 1 && parts[1].length > 0) ? parts[1] : undefined;
        let req = (parts.length > 2 && parts[2].length > 0) ? parts[2] : undefined;
        let act = (parts.length > 3 && parts[3].length > 0) ? parts[3] : undefined;
        if (coll !== collection.cache.active) {
          collection.cache.setActiveCollection(coll);
          socket.send({svc: services.collection.key, cmd: command.client.getCollection, param: coll});
        }
        if (req !== request.cache.active) {
          request.cache.setActiveRequest(req);
        }
        if (act !== request.cache.action) {
          request.cache.setActiveAction(act);
        }
        ui.setPanels(coll, req, act);
        break;
      default:
        console.info("unhandled svc [" + svc + "]");
    }
  }
}
