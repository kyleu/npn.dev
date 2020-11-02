import {ref} from "@vue/composition-api";

import Vue from "vue";
import VueCompositionAPI from "@vue/composition-api";
import VueRouter from "vue-router";

Vue.use(VueCompositionAPI);
Vue.use(VueRouter);

export interface Profile {
  readonly userID: string;
  readonly name: string;
  readonly role: string;
  readonly theme: string;
  readonly navColor: string;
  readonly linkColor: string;
  readonly picture: string;
  readonly locale: string;
}

export const profileRef = ref<Profile>();
