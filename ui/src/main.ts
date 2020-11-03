import {initState} from "@/state/initial";
import Workspace from "./layout/Workspace.vue";
import {debug, NPNDebug} from "./util/debug";
import {router} from "./state/router";
import UIkit from "uikit";
import "@/assets/styles/styles.scss";

// @ts-ignore
import Icons from "uikit/dist/js/uikit-icons";
import {messageHandler} from "@/state/handler";

import Vue from "vue";
import {initDom, setTheme} from "@/npn";

declare global {
  interface Window {
    init: () => void;
    initDom: (t: string, color: string) => void;
    setTheme: (s: string) => void;
    npn: NPNDebug;
    Prism: { manual: boolean };
  }
}

function init(): void {
  // @ts-ignore
  // eslint-disable-next-line
  UIkit.use(Icons);

  window.Prism = window.Prism || {};
  window.Prism.manual = true;

  Vue.config.productionTip = false;

  initState(messageHandler);

  const render = (h: Vue.CreateElement): Vue.VNode => h(Workspace);

  const root = new Vue({router, el: "#npn", render});

  window.npn = {root, router, debug};
}

window.init = init;
// Legacy
window.initDom = initDom;
window.setTheme = setTheme;
