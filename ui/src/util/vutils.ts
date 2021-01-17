import Vue from "vue";
import VueRouter from "vue-router";

import { Breadcrumb, breadcrumbsRef } from "@/layout/breadcrumb";
import { profileRef } from "@/user/profile";

Vue.directive("style-menu", function(el) {
  el.style.backgroundColor = profileRef.value?.settings.menuB || "";
  el.style.color = profileRef.value?.settings.menuF || "";
});

Vue.directive("style-menu-link", function(el) {
  el.style.color = profileRef.value?.settings.menuL || "";
});

Vue.directive("style-menu-section", function(el) {
  el.classList.add("nav-header");
  el.style.color = profileRef.value?.settings.menuF || "";
  el.style.borderBottom = "1px solid " + profileRef.value?.settings.menuF || "";
});

Vue.directive("style-nav", function(el) {
  el.style.backgroundColor = profileRef.value?.settings.navB || "";
  el.style.color = profileRef.value?.settings.navF || "";
});

Vue.directive("style-nav-link", function(el) {
  el.style.borderColor = profileRef.value?.settings.navF || "";
  el.style.color = profileRef.value?.settings.navF || "";
});

Vue.directive("style-button", function(el) {
  el.style.borderColor = profileRef.value?.settings.bodyL || "";
  el.style.color = profileRef.value?.settings.bodyL || "";
});

Vue.directive("style-link", function(el) {
  el.style.color = profileRef.value?.settings.bodyL || "";
});

export function setBC(me: Vue, ...bc: Breadcrumb[]): void {
  breadcrumbsRef.value = bc;
}

export function setBCReq(me: Vue, action: string): void {
  breadcrumbsRef.value = [
    { path: "/c/" + me.$route.params.coll, title: me.$route.params.coll },
    {
      path: "/c/" + me.$route.params.coll + "/" + me.$route.params.req,
      title: me.$route.params.req
    },
    { path: "", title: action }
  ];
}

export function globalRouter(): VueRouter {
  return window.npn.router;
}
