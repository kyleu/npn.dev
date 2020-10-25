import Vue from "vue";
import Vuex from "vuex";
import {initialState, State} from "@/state/state";

Vue.use(Vuex);

export const store = new Vuex.Store({
  state: initialState(),
  mutations: {
    onMessage(state: State): void {
      console.log(state.activeCollection?.key || "no active collection");
    }
  },
  modules: {}
});
