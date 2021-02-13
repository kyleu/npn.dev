import {jsonClone, jsonParse, jsonStr} from "@/util/json";
import {isDebug, logDebug} from "@/util/log";
import {ref} from "@vue/composition-api";

export interface Message {
  readonly channel: string;
  readonly cmd: string;
  readonly param: unknown;
}

function socketUrl(): string {
  const l = document.location;
  let protocol = "ws";
  if (l.protocol === "https:") {
    protocol = "wss";
  }
  return protocol + `://${l.host}/ws`;
}

export class Socket {
  private readonly url: string;
  private sock!: WebSocket;

  connected = false;
  appUnloading = false;
  private connectTime: number | undefined;

  private readonly onOpen: (id: string) => void;
  private readonly onMessage: (m: Message) => void;
  private readonly onError: (channel: string, err: string) => void;

  private pauseSeconds = 0;
  private pendingMessages: Message[] = [];

  constructor(openF: (id: string) => void, recvF: (m: Message) => void, errF: (channel: string, err: string) => void, url?: string) {
    this.url = (url && url.length > 0) ? url : socketUrl();
    this.onOpen = openF;
    this.onMessage = recvF;
    this.onError = errF;

    window.onbeforeunload = (): void => {
      this.setAppUnloading();
    };

    this.init();
  }

  setAppUnloading(): void {
    this.appUnloading = true;
  }

  private init(): void {
    this.connectTime = Date.now();
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
    this.onOpen("");
    this.pendingMessages.forEach((msg) => this.send(msg));
    this.pendingMessages = [];
  }

  private socketConnect(): void {
    if (!this.onMessage) {
      throw "onMessage not initialized";
    }
    this.init();
  }

  send(msg: Message): void {
    if (this.connected) {
      if (isDebug()) {
        logDebug(`OUT(${msg.channel}): ${msg.cmd}`, jsonClone(msg.param));
      }
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
      if (this.pauseSeconds < 1) {
        this.pauseSeconds = 1;
      }
      this.pauseSeconds = this.pauseSeconds * 2;
      logDebug(`socket closed immediately, reconnecting in ${this.pauseSeconds} seconds`);
      setTimeout(() => {
        this.socketConnect();
      }, this.pauseSeconds * 1000);
    } else {
      logDebug("socket closed after [" + elapsed + "ms]");
      this.socketConnect();
    }
  }

  private onSocketClose(): void {
    if (!this.appUnloading) {
      this.disconnect();
    }
  }
}

export const hostRef = ref<string>("");
export const socketRef = ref<Socket>();
