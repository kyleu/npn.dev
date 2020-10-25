import {jsonParse, jsonStr} from "@/util/json";

export interface Message {
  readonly svc: string;
  readonly cmd: string;
  readonly param: any;
}

export default class Socket {
  private url: string
  private sock!: WebSocket

  connected = false
  appUnloading = false
  private connectTime: number | undefined;

  private onOpen: (id: string) => void
  private onMessage: (m: Message) => void
  private onError: (svc: string, err: string) => void

  private pauseSeconds = 0;
  private pendingMessages: Message[] = [];

  constructor(openF: (id: string) => void, recvF: (m: Message) => void, errF: (svc: string, err: string) => void, url?: string) {
    this.url = url ? url : socketUrl();

    this.onOpen = openF;
    this.onMessage = recvF;
    this.onError = errF;

    window.onbeforeunload = () => {
      this.setAppUnloading();
    };

    this.init()
  }

  setAppUnloading() {
    this.appUnloading = true;
  }

  private init() {
    this.sock = new WebSocket(this.url);
    this.sock.onopen = this.onSocketOpen;
    this.sock.onmessage = (event) => this.onMessage(jsonParse(event.data));
    this.sock.onerror = (event) => this.onError("socket", event.type);
    this.sock.onclose = this.onSocketClose;
  }

  private onSocketOpen(): void {
    // log.info("socket connected");
    this.connected = true;
    this.pauseSeconds = 1;
    this.pendingMessages.forEach(this.send);
    this.pendingMessages = [];
    this.onOpen("");
  }

  private socketConnect() {
    if (!this.onMessage) {
      throw "onMessage not initialized";
    }
    this.init();
  }

  private send(msg: Message) {
    // console.debug("out", msg);
    if (this.connected) {
      const m = jsonStr(msg);
      this.sock.send(m);
    } else {
      this.pendingMessages.push(msg);
    }
  }

  private onSocketClose() {
    const self = this;
    function disconnect() {
      self.connected = false;
      const elapsed = Date.now() - self.connectTime!;

      if (elapsed < 2000) {
        self.pauseSeconds = self.pauseSeconds * 2;
        // console.debug(`socket closed immediately, reconnecting in ${pauseSeconds} seconds`);
        setTimeout(() => {
          self.socketConnect();
        }, self.pauseSeconds * 1000);
      } else {
        // log.info("socket closed after [" + elapsed + "ms]");
        self.socketConnect();
      }
    }

    if (!self.appUnloading) {
      disconnect();
    }
  }
}

function socketUrl() {
  const l = document.location;
  let protocol = "ws";
  if (l.protocol === "https:") {
    protocol = "wss";
  }
  return protocol + `://${l.host}/s`;
}
