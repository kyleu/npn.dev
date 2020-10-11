namespace routing {
  export function recv(msg: socket.Message) {
    if (socket.debug) {
      console.debug("in", msg);
    }
    switch (msg.svc) {
      case services.system.key:
        system.onSystemMessage(msg.cmd, msg.param);
        break;
      case services.collection.key:
        collection.onCollectionMessage(msg.cmd, msg.param);
        break;
      case services.request.key:
        request.onRequestMessage(msg.cmd, msg.param);
        break;
      default:
        console.warn(`unhandled message for service [${msg.svc}]`);
    }
  }

  export function route(p: string) {
    let parts = p.split("/");
    parts = parts.filter(x => x.length > 0);
    console.debug("nav: " + parts.join(" -> "));

    const svc = (parts.length > 0) ? parts[0] : "c";
    switch (svc) {
      case "c":
        const coll = (parts.length > 1 && parts[1].length > 0) ? parts[1] : undefined;
        const req = (parts.length > 2 && parts[2].length > 0) ? parts[2] : undefined;
        const act = (parts.length > 3 && parts[3].length > 0) ? parts[3] : undefined;
        const extra = (parts.length > 4) ? parts.slice(4) : [];
        const currColl = collection.cache.active;
        collection.cache.setActiveCollection(coll);
        if (coll !== currColl && coll) {
          socket.send({svc: services.collection.key, cmd: command.client.getCollection, param: coll});
        }
        request.cache.setActiveRequest(currColl, req);
        request.cache.setActiveAction(currColl, act, extra);
        ui.setPanels(coll, req, act, extra);
        break;
      default:
        console.warn("unhandled svc [" + svc + "]");
    }
  }
}
