import Vue from "vue";
import Workspace from "./layout/Workspace.vue";
import {router} from "./state/router";
import {newStore} from "./state/store";
import UIkit from "uikit";
import "@/assets/styles/styles.scss";

// @ts-ignore
import Icons from "uikit/dist/js/uikit-icons";
import {messageHandler} from "@/state/handler";

// @ts-ignore
// eslint-disable-next-line
(UIkit as any).use(Icons);

// @ts-ignore
// eslint-disable-next-line
(window as any).UIkit = UIkit;

// @ts-ignore
// eslint-disable-next-line
(window as any).Prism = (window as any).Prism || {};
// @ts-ignore
// eslint-disable-next-line
(window as any).Prism.manual = true;

Vue.config.productionTip = false;

const render = (h: Vue.CreateElement): Vue.VNode => h(Workspace);

const store = newStore(messageHandler);

const root = new Vue({router, store, el: "#npn", render});

// @ts-ignore
// eslint-disable-next-line
(window as any).npn = {root, router, store};
