namespace socket {
  export interface Message {
    readonly svc: string;
    readonly cmd: string;
    readonly param: any;
  }

  const debug = true;

  let socket: WebSocket;
  let appUnloading = false;
  let currentService = "";
  let currentID = "";

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
    // system.cache.currentService = svc;
    // system.cache.currentID = id;
    // system.cache.connectTime = Date.now();

    socket = new WebSocket(socketUrl());
    socket.onopen = () => {
      send({ svc: svc, cmd: "connect", param: id });
    };
    socket.onmessage = (event) => {
      const msg = JSON.parse(event.data);
      onSocketMessage(msg);
    };
    socket.onerror = (event) => {
      // rituals.onError(services.system, event.type);
    };
    socket.onclose = () => {
      onSocketClose();
    };
  }

  export function send(msg: Message) {
    if (debug) {
      console.debug("out", msg);
    }
    socket.send(JSON.stringify(msg));
  }

  export function onSocketMessage(msg: Message) {
    if (debug) {
      console.debug("in", msg);
    }
    switch (msg.svc) {
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
