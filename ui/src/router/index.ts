import Vue from "vue";
import VueRouter, { RouteConfig } from "vue-router";
import Home from "@/views/Home.vue";
import About from "@/views/About.vue";
import CollectionDetail from "@/views/CollectionDetail.vue";

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
  {
    path: "/",
    name: "Home",
    component: Home
  },
  {
    path: "/c/:coll",
    name: "Collection",
    component: CollectionDetail
  },
  {
    path: "/about",
    name: "About",
    component: About
  }
];

export const router = new VueRouter({
  mode: "history",
  base: "/",
  routes
});
