import {initState} from "@/state/initial";
import Workspace from "./layout/Workspace.vue";
import {debug, NPNDebug} from "./util/debug";
import {router} from "./state/router";
import UIkit from "uikit";
import "@/assets/styles/styles.scss";

// @ts-ignore
import {messageHandler} from "@/state/handler";

import Vue from "vue";

declare global {
  interface Window {
    init: () => void;
    npn: NPNDebug;
    Prism: { manual: boolean };
    // @ts-ignore
    // eslint-disable-next-line
    UIkit: any;
  }
}

function init(): void {
  window.UIkit = UIkit;

  window.Prism = window.Prism || {};
  window.Prism.manual = true;

  Vue.config.productionTip = false;

  initState(messageHandler);

  const render = (h: Vue.CreateElement): Vue.VNode => h(Workspace);

  const root = new Vue({router, el: "#npn", render});

  window.npn = {root, router, debug};
}

window.init = init;
