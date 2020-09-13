namespace socket {
  export interface Message {
    readonly svc: string;
    readonly cmd: string;
    readonly param: any;
  }

  const debug = true;

  let sock: WebSocket;
  let appUnloading = false;
  let currentService = "";
  let currentID = "";
  let connectTime: number | undefined;

  function socketUrl() {
    const l = document.location;
    let protocol = "ws";
    if (l.protocol === "https:") {
      protocol = "wss";
    }
    return protocol + `://${l.host}/s`;
  }

  export function setAppUnloading() {
    appUnloading = true;
  }

  export function socketConnect(svc: string, id: string) {
    currentService = svc;
    currentID = id;
    connectTime = Date.now();

    sock = new WebSocket(socketUrl());
    sock.onopen = () => {
      send({ svc: services.system.key, cmd: command.client.connect, param: id });
    };
    sock.onmessage = (event) => {
      const msg = JSON.parse(event.data);
      onSocketMessage(msg);
    };
    sock.onerror = (event) => {
      npn.onError("socket", event.type);
    };
    sock.onclose = () => {
      onSocketClose();
    };
  }

  export function send(msg: Message) {
    if (debug) {
      console.debug("out", msg);
    }
    sock.send(JSON.stringify(msg));
  }

  export function onSocketMessage(msg: Message) {
    if (debug) {
      console.debug("in", msg);
    }

    switch (msg.svc) {
      case services.system.key:
        system.onSystemMessage(msg.cmd, msg.param);
        break;
      case services.collection.key:
        collection.onCollectionMessage(msg.cmd, msg.param);
        break;
      default:
        console.warn(`unhandled message for service [${msg.svc}]`);
    }
  }

  function onSocketClose() {
    function disconnect(seconds: number) {
      if (debug) {
        console.info(`socket closed, reconnecting in ${seconds} seconds`);
      }
      setTimeout(() => {
        socketConnect(currentService, currentID);
      }, seconds * 1000);
    }

    if (!appUnloading) {
      disconnect(10);
    }
  }
}
