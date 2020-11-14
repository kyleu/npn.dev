import VueRouter, {RouteConfig} from "vue-router";
import About from "@/views/About.vue";
import Home from "@/views/Home.vue";
import NotFound from "@/views/NotFound.vue";
import Settings from "@/user/Settings.vue";
import CallResultView from "@/call/CallResultView.vue";
import CollectionDetail from "@/collection/CollectionDetail.vue";
import CollectionIndex from "@/collection/CollectionIndex.vue";
import RequestDetail from "@/request/RequestDetail.vue";
import RequestEditor from "@/request/editor/RequestEditor.vue";
import RequestTransform from "@/request/transform/RequestTransform.vue";
import {callResultRef, transformResultRef} from "@/request/state";

const routes: Array<RouteConfig> = [
  {
    path: "/",
    name: "Home",
    component: Home
  },
  {
    path: "/u",
    name: "Settings",
    component: Settings
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
        component: CallResultView,
        beforeEnter: (to, from, next): void => {
          callResultRef.value = undefined;
          next();
        }
      },
      {
        path: "transform/:tx",
        name: "RequestTransform",
        component: RequestTransform,
        beforeEnter: (to, from, next): void => {
          transformResultRef.value = undefined;
          next();
        }
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
