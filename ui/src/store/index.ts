import Vue from "vue";
import Vuex from "vuex";
import {State} from "@/store/state";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    collections: [
      {
        key: "test1",
        title: "Test 1",
        description: "Test Collection 1"
      },
      {
        key: "test2",
        title: "Test 2",
        description: "Test Collection 2"
      }
    ],
    activeCollection: undefined
  } as State,
  mutations: {},
  actions: {},
  modules: {}
});
