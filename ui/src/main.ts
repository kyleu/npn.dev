import Vue from "vue";
import Workspace from "./layout/Workspace.vue";
import {router} from "./state/router";
import {newStore} from "./state/store";
import UIkit from "uikit";
import "@/assets/styles/styles.scss";

// @ts-ignore
import Icons from "uikit/dist/js/uikit-icons";
import {Message} from "@/socket/socket";

// @ts-ignore
// eslint-disable-next-line
(UIkit as any).use(Icons);

// @ts-ignore
// eslint-disable-next-line
(window as any).UIkit = UIkit;

Vue.config.productionTip = false;

const render = (h: Vue.CreateElement): Vue.VNode => h(Workspace);

const handler = (msg: Message): void => {
  console.log(msg);
}

const store = newStore(handler);

const root = new Vue({router, store, render}).$mount("#npn");

// @ts-ignore
// eslint-disable-next-line
(window as any).npn = { root, router, store };
