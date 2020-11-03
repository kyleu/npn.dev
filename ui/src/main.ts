import {initState} from "@/state/initial";
import Workspace from "./layout/Workspace.vue";
import {debug} from "./util/debug";
import {router} from "./state/router";
import UIkit from "uikit";
import "@/assets/styles/styles.scss";

// @ts-ignore
import Icons from "uikit/dist/js/uikit-icons";
import {messageHandler} from "@/state/handler";

import Vue from "vue";
import {initDom, setTheme} from "@/npn";

// @ts-ignore
// eslint-disable-next-line
const w = (window as any)

function init(): void {
  // @ts-ignore
  // eslint-disable-next-line
  (UIkit as any).use(Icons);

  w.UIkit = UIkit;

  w.Prism = w.Prism || {};
  w.Prism.manual = true;

  Vue.config.productionTip = false;

  initState(messageHandler);

  const render = (h: Vue.CreateElement): Vue.VNode => h(Workspace);

  const root = new Vue({router, el: "#npn", render});

  w.npn = {root, router, debug};
}

w.init = init;
w.initDom = initDom;
w.setTheme = setTheme;
