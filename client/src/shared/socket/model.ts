namespace socket {
  export interface Message {
    readonly svc: string;
    readonly cmd: string;
    readonly param: any;
  }

  export const debug = true;

  export let appUnloading = false;

  export let currentService = "";
  export let currentID = "";
  export let connectTime: number | undefined;

  export let bypass = false;

  export function setAppUnloading() {
    appUnloading = true;
  }

  export function send(msg: Message) {
    if(bypass) { bypassSend(msg) } else { socketSend(msg) }
  }

  export function recv(msg: Message) {
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
}
