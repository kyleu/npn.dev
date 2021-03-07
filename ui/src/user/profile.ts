import {ref, watchEffect} from "@vue/composition-api";

import Vue from "vue";
import VueCompositionAPI from "@vue/composition-api";
import VueRouter from "vue-router";
import {Color} from "@/user/colors";

Vue.use(VueCompositionAPI);
Vue.use(VueRouter);

export interface UserSettings {
  mode: string;
  navB: Color;
  navF: Color;
  menuB: Color;
  menuF: Color;
  menuL: Color;
  bodyB: Color;
  bodyL: Color;
}

export const defaultSettings: UserSettings = {
  mode: 'light',
  navB: '#193441',
  navF: '#dddddd',
  menuB: '#3e606f',
  menuF: '#cccccc',
  menuL: '#91aa9d',
  bodyB: '#fcfff5',
  bodyL: '#2f657f'
};

export interface Profile {
  readonly userID: string;
  name: string;
  role: string;
  settings: UserSettings;
  picture: string;
  locale: string;
}

export const defaultProfile: Profile = {
  userID: "00000000-0000-0000-0000-000000000000",
  name: "npn",
  role: "",
  settings: defaultSettings,
  picture: "",
  locale: "en-US"
};

export const profileRef = ref<Profile>();

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
      stylesheet.addRule(".uk-tab > .uk-active > a", "");
    }


    const rules = stylesheet.cssRules || stylesheet.rules;
    stylesheet.removeRule(rules.length - 1);
    stylesheet.removeRule(rules.length - 1);
    stylesheet.removeRule(rules.length - 1);

    stylesheet.addRule("a[href]", `color: ${s.bodyL}`, rules.length);
    stylesheet.addRule("body", `background-color: ${s.bodyB} !important;`, rules.length);
    stylesheet.addRule(".uk-tab > .uk-active > a", `border-color: ${s.bodyL} !important; border-width: 2px;`, rules.length);
  }
});
