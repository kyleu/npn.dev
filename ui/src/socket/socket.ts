import {jsonParse, jsonStr} from "@/util/json";

export interface Message {
  readonly svc: string;
  readonly cmd: string;
  // @ts-ignore
  // eslint-disable-next-line
  readonly param: any;
}

function socketUrl(): string {
  const l = document.location;
  let protocol = "ws";
  if (l.protocol === "https:") {
    protocol = "wss";
  }
  return protocol + `://${l.host}/s`;
}

export class Socket {
  private readonly url: string
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
    this.url = (url && url.length > 0) ? url : socketUrl();
    console.log(this.url);
    this.onOpen = openF;
    this.onMessage = recvF;
    this.onError = errF;

    window.onbeforeunload = (): void => {
      this.setAppUnloading();
    };

    this.init()
  }

  setAppUnloading(): void {
    this.appUnloading = true;
  }

  private init(): void {
    this.sock = new WebSocket(this.url);
    this.sock.onopen = (): void => this.onSocketOpen();
    this.sock.onmessage = (event): void => this.onMessage(jsonParse(event.data));
    this.sock.onerror = (event): void => this.onError("socket", event.type);
    this.sock.onclose = (): void => this.onSocketClose();
  }

  private onSocketOpen(): void {
    // log.info("socket connected");
    this.connected = true;
    this.pauseSeconds = 1;
    this.pendingMessages.forEach(this.send);
    this.pendingMessages = [];
    this.onOpen("");
  }

  private socketConnect(): void {
    if (!this.onMessage) {
      throw "onMessage not initialized";
    }
    this.init();
  }

  private send(msg: Message): void {
    // console.debug("out", msg);
    if (this.connected) {
      const m = jsonStr(msg);
      this.sock.send(m);
    } else {
      this.pendingMessages.push(msg);
    }
  }

  private disconnect(): void {
    this.connected = false;
    const elapsed = Date.now() - (this.connectTime || 0);

    if (elapsed < 2000) {
      this.pauseSeconds = this.pauseSeconds * 2;
      // console.debug(`socket closed immediately, reconnecting in ${pauseSeconds} seconds`);
      setTimeout(() => {
        this.socketConnect();
      }, this.pauseSeconds * 1000);
    } else {
      // log.info("socket closed after [" + elapsed + "ms]");
      this.socketConnect();
    }
  }

  private onSocketClose(): void {
    if (!this.appUnloading) {
      this.disconnect();
    }
  }
}
