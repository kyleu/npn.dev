import Vue from "vue";
import Workspace from "./Workspace.vue";
import router from "./router";
import store from "./store";
import UIkit from "uikit";
import "@/assets/styles/styles.scss";

// @ts-ignore
import Icons from "uikit/dist/js/uikit-icons";

// @ts-ignore
(UIkit as any).use(Icons);

// @ts-ignore
(window as any).UIkit = UIkit;

Vue.config.productionTip = false;

const render = (h: Vue.CreateElement): Vue.VNode => h(Workspace);

new Vue({
  router: router,
  store: store,
  render: render
}).$mount("#app");
