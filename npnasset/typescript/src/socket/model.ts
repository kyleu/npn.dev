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
}
