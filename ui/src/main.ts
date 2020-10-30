import Vue from "vue";
import Workspace from "./layout/Workspace.vue";
import {router} from "./state/router";
import {newStore} from "./state/store";
import UIkit from "uikit";
import "@/assets/styles/styles.scss";

// @ts-ignore
import Icons from "uikit/dist/js/uikit-icons";
import {Message} from "@/socket/socket";
import {State} from "@/state";
import {Collection} from "@/collection/collection";
import {logDebug, logWarn} from "@/util/log";
import {NPNRequest} from "@/request/model";

// @ts-ignore
// eslint-disable-next-line
(UIkit as any).use(Icons);

// @ts-ignore
// eslint-disable-next-line
(window as any).UIkit = UIkit;

Vue.config.productionTip = false;

const render = (h: Vue.CreateElement): Vue.VNode => h(Workspace);

const handler = (state: State, msg: Message): void => {
  logDebug("IN", msg);
  switch (msg.cmd) {
    case "collections":
      state.collections = msg.param as Collection[];
      break;
    case "collectionDetail":
      state.setCollectionRequestSummaries(msg.param.key, msg.param.requests);
      break;
    case "requestDetail":
      state.setRequestDetail(msg.param.coll, msg.param.req as NPNRequest);
      break;
    default:
      logWarn("unhandled message [" + msg.cmd + "]", msg);
  }
}

const store = newStore(handler);

const root = new Vue({router, store, el: "#npn", render});

// @ts-ignore
// eslint-disable-next-line
(window as any).npn = { root, router, store };
