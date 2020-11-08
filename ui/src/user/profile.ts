import {ref} from "@vue/composition-api";

import Vue from "vue";
import VueCompositionAPI from "@vue/composition-api";
import VueRouter from "vue-router";

Vue.use(VueCompositionAPI);
Vue.use(VueRouter);

export interface UserSettings {
  readonly navColor: string;
  readonly linkColor: string;
}

export interface Profile {
  readonly userID: string;
  readonly name: string;
  readonly role: string;
  readonly theme: string;
  readonly settings: UserSettings;
  readonly picture: string;
  readonly locale: string;
}

export const profileRef = ref<Profile>();
