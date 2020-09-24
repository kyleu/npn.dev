namespace socket {
  export interface Message {
    readonly svc: string;
    readonly cmd: string;
    readonly param: any;
  }

  const debug = true;

  let sock: WebSocket;
  let connected = false;
  let pauseSeconds = 0;
  let appUnloading = false;
  let pendingMessages: Message[] = [];

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
    sock.onopen = onSocketOpen;
    sock.onmessage = (event) => onSocketMessage(JSON.parse(event.data));
    sock.onerror = (event) => npn.onError("socket", event.type);
    sock.onclose = onSocketClose;
  }

  export function send(msg: Message) {
    if (connected) {
      if (debug) {
        console.debug("out", msg);
      }
      const m = json.str(msg);
      sock.send(m);
    } else {
      pendingMessages.push(msg);
    }
  }

  function onSocketOpen() {
    log.info("socket connected");
    connected = true;
    pauseSeconds = 1;
    pendingMessages.forEach(send);
    pendingMessages = [];
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
      case services.request.key:
        request.onRequestMessage(msg.cmd, msg.param);
        break;
      default:
        console.warn(`unhandled message for service [${msg.svc}]`);
    }
  }

  function onSocketClose() {
    function disconnect() {
      connected = false;
      const elapsed = Date.now() - connectTime!;

      if (elapsed < 2000) {
        pauseSeconds = pauseSeconds * 2;
        if (debug) {
          console.info(`socket closed immediately, reconnecting in ${pauseSeconds} seconds`);
        }
        setTimeout(() => {
          socketConnect(currentService, currentID);
        }, pauseSeconds * 1000);
      } else {
        log.info("socket closed after [" + elapsed + "ms]");
        socketConnect(currentService, currentID);
      }
    }

    if (!appUnloading) {
      disconnect();
    }
  }
}
