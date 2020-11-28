import {initState} from "@/state/initial";
import Workspace from "./layout/Workspace.vue";
import {onDebug, NPNDebug} from "./util/debug";
import {router} from "./state/router";
import UIkit from "uikit";
import "@/assets/styles/styles.scss";

// @ts-ignore
import {messageHandler} from "@/state/handler";

import Vue from "vue";
import {logInfo} from "@/util/log";

declare global {
  interface Window {
    init: () => void;
    npn: NPNDebug;
    UIkit: object;
  }
}

function init(): void {
  window.UIkit = UIkit;

  Vue.config.productionTip = false;

  initState(messageHandler);

  const render = (h: Vue.CreateElement): Vue.VNode => h(Workspace);

  const root = new Vue({router, el: "#npn", render});

  logInfo("npn has started");

  window.npn = {root, router, onDebug};
}

window.init = init;
