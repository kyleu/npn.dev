import {initState} from "@/state/initial";
import Workspace from "./layout/Workspace.vue";
import {onDebug, NPNDebug} from "./util/debug";
import {router} from "./state/router";
import UIkit from "uikit";
import "@/assets/styles/styles.scss";

import {messageHandler} from "@/state/handler";

import Vue from "vue";
import {logInfo, setDebug, setPublic} from "@/util/log";

declare global {
  interface Window {
    init: (debug: boolean) => void;
    npn: NPNDebug;
    UIkit: object;
  }
}

function init(debug?: boolean, pub?: boolean): void {
  setDebug(debug || false);
  setPublic(pub || false);
  window.UIkit = UIkit;

  Vue.config.productionTip = false;

  initState(messageHandler);

  const render = (h: Vue.CreateElement): Vue.VNode => h(Workspace);

  const root = new Vue({router, el: "#npn", render});

  let msg = "[npn] has started";
  if(debug) {
    msg = `${msg} (debug)`;
  }
  logInfo(msg);

  window.npn = {root, router, onDebug};
}

window.init = init;
