import {ref, watchEffect} from "@vue/composition-api";

import Vue from "vue";
import VueCompositionAPI from "@vue/composition-api";
import VueRouter from "vue-router";
import {Color} from "@/user/colors";

Vue.use(VueCompositionAPI);
Vue.use(VueRouter);

export interface UserSettings {
  mode: string;
  bodyB: Color;
  bodyL: Color;
  navB: Color;
  navF: Color;
  menuB: Color;
  menuF: Color;
  menuL: Color;
}

export interface Profile {
  readonly userID: string;
  name: string;
  role: string;
  settings: UserSettings;
  picture: string;
  locale: string;
}

export const profileRef = ref<Profile>();
export const tempThemeRef = ref<string>("");

let stylesheet: CSSStyleSheet | undefined;

watchEffect(() => {
  const s = profileRef.value?.settings;
  if (s) {
    if (s.mode === "dark") {
      document.body.classList.remove("uk-dark");
      document.body.classList.add("uk-light");
    } else {
      document.body.classList.remove("uk-light");
      document.body.classList.add("uk-dark");
    }

    if (!stylesheet) {
      stylesheet = document.styleSheets[0];
      stylesheet.addRule("a[href]", "");
      stylesheet.addRule("body", "");
      stylesheet.addRule(".uk-tab > .uk-active > a", "")
    }

    const ss = stylesheet;

    const rules = ss.cssRules || ss.rules;
    ss.removeRule(rules.length - 1);
    ss.removeRule(rules.length - 1);
    ss.removeRule(rules.length - 1);

    ss.addRule("a[href]", `color: ${s.bodyL}`, rules.length);
    ss.addRule("body", `background-color: ${s.bodyB} !important;`, rules.length);
    ss.addRule(".uk-tab > .uk-active > a", `border-color: ${s.bodyL} !important; border-width: 2px;`, rules.length)
  }
})
