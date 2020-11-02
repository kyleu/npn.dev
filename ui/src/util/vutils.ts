import Vue from "vue";
import VueCompositionAPI from "@vue/composition-api";
import VueRouter from "vue-router";

import {Breadcrumb, breadcrumbsRef} from "@/layout/breadcrumb";

Vue.use(VueCompositionAPI);
Vue.use(VueRouter);

export function setBC(me: Vue, ...bc: Breadcrumb[]): void {
  breadcrumbsRef.value = bc;
}

export function setBCReq(me: Vue, action: string): void {
  breadcrumbsRef.value = [
    {path: "/c/" + me.$route.params.coll, title: me.$route.params.coll},
    {path: "/c/" + me.$route.params.coll + "/" + me.$route.params.req, title: me.$route.params.req},
    {path: "", title: action}
  ];
}
