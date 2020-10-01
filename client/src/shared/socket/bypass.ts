declare const npn_handler: any;

namespace socket {
  export function initBypass() {
    bypass = true;
    connected = true;
    nav.enabled = false
  }

  export function bypassSend(msg: Message) {
    if (debug) {
      console.debug("out", msg);
    }
    npn_handler(JSON.stringify(msg, null, 2));
  }
}
