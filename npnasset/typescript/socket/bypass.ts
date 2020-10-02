declare const npn_handler: any | undefined;

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
    if (npn_handler) {
      npn_handler(JSON.stringify(msg, null, 2));
    } else {
      console.warn("no bypass handler configured");
    }
  }
}
