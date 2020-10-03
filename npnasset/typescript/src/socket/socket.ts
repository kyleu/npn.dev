namespace socket {
  let sock: WebSocket;
  export let connected = false;
  let pauseSeconds = 0;
  let pendingMessages: Message[] = [];
  let onOpen: (id: string) => void;
  let onMessage: (m: Message) => void;
  let onError: (svc: string, err: string) => void;

  export function init(open: (id: string) => void, recv: (m: Message) => void, err: (svc: string, err: string) => void) {
    onOpen = open;
    onMessage = recv;
    onError = err;
  }

  function socketUrl() {
    const l = document.location;
    let protocol = "ws";
    if (l.protocol === "https:") {
      protocol = "wss";
    }
    return protocol + `://${l.host}/s`;
  }

  export function initSocket() {
    sock = new WebSocket(socketUrl());
    sock.onopen = onSocketOpen;
    sock.onmessage = (event) => onMessage(json.parse(event.data));
    sock.onerror = (event) => onError("socket", event.type);
    sock.onclose = onSocketClose;
  }

  export function socketConnect(svc: string, id: string, useBypass?: boolean) {
    currentService = svc;
    currentID = id;
    connectTime = Date.now();

    if (!onMessage) {
      throw "onMessage not initialized";
    }

    if(useBypass) {
      initBypass();
    } else {
      initSocket();
    }
  }

  function onSocketOpen() {
    log.info("socket connected");
    connected = true;
    pauseSeconds = 1;
    pendingMessages.forEach(send);
    pendingMessages = [];
    onOpen(currentID);
  }

  function onSocketClose() {
    function disconnect() {
      connected = false;
      const elapsed = Date.now() - connectTime!;

      if (elapsed < 2000) {
        pauseSeconds = pauseSeconds * 2;
        if (debug) {
          console.debug(`socket closed immediately, reconnecting in ${pauseSeconds} seconds`);
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

  export function socketSend(msg: socket.Message) {
    if (debug) {
      console.debug("out", msg);
    }
    if (connected) {
      const m = json.str(msg);
      sock.send(m);
    } else {
      pendingMessages.push(msg);
    }
  }
}
