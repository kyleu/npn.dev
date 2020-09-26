namespace socket {
  export function route(p: string) {
    console.log("NAV: " + p);
    let parts = p.split("/");
    parts = parts.filter(x => x.length > 0);
    console.info("nav: " + parts.join(" -> "));

    if (parts.length === 0 || parts[0].length === 0) {
      ui.setPanels(undefined, undefined, undefined, []);
      return; // index
    }
    const svc = parts[0];
    switch (svc) {
      case "c":
        let coll = (parts.length > 1 && parts[1].length > 0) ? parts[1] : undefined;
        let req = (parts.length > 2 && parts[2].length > 0) ? parts[2] : undefined;
        let act = (parts.length > 3 && parts[3].length > 0) ? parts[3] : undefined;
        let extra = (parts.length > 4) ? parts.slice(4) : [];
        if (coll !== collection.cache.active) {
          collection.cache.setActiveCollection(coll);
          socket.send({svc: services.collection.key, cmd: command.client.getCollection, param: coll});
        }
        request.cache.setActiveRequest(req);
        request.cache.setActiveAction(act, extra);
        ui.setPanels(coll, req, act, extra);
        break;
      default:
        console.info("unhandled svc [" + svc + "]");
    }
  }
}
