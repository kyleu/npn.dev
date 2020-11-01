import Vue from "vue";
import VueRouter, { RouteConfig } from "vue-router";
import Home from "@/views/Home.vue";
import About from "@/views/About.vue";
import NotFound from "@/views/NotFound.vue";
import CollectionDetail from "@/collection/CollectionDetail.vue";
import CollectionIndex from "@/collection/CollectionIndex.vue";
import RequestDetail from "@/request/RequestDetail.vue";
import RequestEditor from "@/request/editor/RequestEditor.vue";
import RequestTransform from "@/request/RequestTransform.vue";
import CallResultPanel from "@/call/CallResultPanel.vue";

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
  {
    path: "/",
    name: "Home",
    component: Home
  },
  {
    path: "/c",
    name: "CollectionIndex",
    component: CollectionIndex,
  },
  {
    path: "/c/:coll",
    name: "CollectionDetail",
    component: CollectionDetail
  },
  {
    path: "/c/:coll/:req",
    component: RequestDetail,
    children: [
      {
        path: "",
        name: "RequestDetail",
        component: RequestEditor
      },
      {
        path: "call",
        name: "CallResult",
        component: CallResultPanel
      },
      {
        path: "transform/:tx",
        name: "RequestTransform",
        component: RequestTransform
      }
    ]
  },
  {
    path: "/about",
    name: "About",
    component: About
  },
  {
    path: "*",
    name: "NotFound",
    component: NotFound
  }
];

export const router = new VueRouter({
  mode: "history",
  base: "/",
  routes
});
