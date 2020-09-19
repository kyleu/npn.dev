namespace socket {
  export function route(p: string) {
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
}
