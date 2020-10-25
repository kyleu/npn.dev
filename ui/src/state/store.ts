import Vue from "vue";
import Vuex, {Store} from "vuex";
import {initialState, State} from "@/state";
import {Message, Socket} from "@/socket/socket";
import {jsonStr} from "@/util/json";

Vue.use(Vuex);

function openF(): void {
  console.log("websocket open");
}
function recvF(m: Message): void {
  console.log("websocket message: " + jsonStr(m));
}
function errF(err: string): void {
  console.log("websocket err: " + err);
}

export const store = function(): Store<State> {
  const state = initialState();
  const socket = new Socket(openF, recvF, errF, state.url);

  return new Vuex.Store({
    state: state,
    mutations: {
      onMessage(state: State): void {
        console.log(state.activeCollection?.key || "no active collection");
      }
    },
    modules: {}
  })
}();
