import Vue from "vue";
import Vuex from "vuex";
import {initialState} from "@/store/state";

Vue.use(Vuex);

export const store = new Vuex.Store({
  state: initialState(),
  mutations: {},
  actions: {},
  modules: {}
});
