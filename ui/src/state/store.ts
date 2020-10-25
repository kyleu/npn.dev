import Vue from "vue";
import Vuex, {Store} from "vuex";
import {initialState, State} from "@/state";
import {Message, Socket} from "@/socket/socket";

Vue.use(Vuex);

export function newStore(onMsg: (m: Message) => void): Store<State> {
  const ret = new Vuex.Store({
    state: initialState(),
    mutations: {
      onMessage(state: State, msg: Message): void {
        onMsg(msg);
      }
    },
    modules: {}
  });

  function openF(): void {
    console.log("websocket open");
  }
  function recvF(m: Message): void {
    ret.commit("onMessage", {msg: m})
  }
  function errF(err: string): void {
    console.log("websocket err: " + err);
  }

  //const socket =
  new Socket(openF, recvF, errF, ret.state.url);

  return ret;
}
