import VueRouter, {RouteConfig} from "vue-router";
import About from "@/views/About.vue";
import Home from "@/views/Home.vue";
import NotFound from "@/views/NotFound.vue";
import ProfileView from "@/user/ProfileView.vue";
import CallResultView from "@/call/CallResultView.vue";
import CollectionDetail from "@/collection/CollectionDetail.vue";
import CollectionIndex from "@/collection/CollectionIndex.vue";
import RequestDetail from "@/request/RequestDetail.vue";
import RequestEditor from "@/request/editor/RequestEditor.vue";
import RequestTransform from "@/transform/RequestTransform.vue";
import SessionIndex from "@/session/SessionIndex.vue";
import SessionDetail from "@/session/SessionDetail.vue";
import {requestResultsRef} from "@/call/state";
import {collectionTransformResultRef, requestTransformResultRef, sessionTransformResultRef} from "@/transform/state";
import CollectionTransform from "@/transform/CollectionTransform.vue";
import Testbed from "@/views/Testbed.vue";
import SessionTransform from "@/transform/SessionTransform.vue";
import Config from "@/views/Config.vue";
import Help from "@/views/Help.vue";
import SearchResults from "@/search/SearchResults.vue";
import ImportForm from "@/import/ImportForm.vue";
import ImportResults from "@/import/ImportResults.vue";

const routes: Array<RouteConfig> = [
  {
    path: "/",
    name: "Home",
    component: Home
  },
  {
    path: "/u",
    name: "Profile",
    component: ProfileView
  },
  {
    path: "/search/:q",
    name: "SearchResults",
    component: SearchResults
  },
  {
    path: "/s",
    name: "SessionIndex",
    component: SessionIndex,
  },
  {
    path: "/s/:sess",
    name: "SessionDetail",
    component: SessionDetail
  },
  {
    path: "/x/:sess",
    name: "SessionTransform",
    component: SessionTransform,
    beforeEnter: (to, from, next): void => {
      sessionTransformResultRef.value = undefined;
      next();
    }
  },
  {
    path: "/x/:coll/:fmt",
    name: "CollectionTransform",
    component: CollectionTransform,
    beforeEnter: (to, from, next): void => {
      collectionTransformResultRef.value = undefined;
      next();
    }
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
          requestResultsRef.value = {id: "", coll: "", req: "", cycles: []};
          next();
        }
      },
      {
        path: "transform/:tx",
        name: "RequestTransform",
        component: RequestTransform,
        beforeEnter: (to, from, next): void => {
          requestTransformResultRef.value = undefined;
          next();
        }
      }
    ]
  },
  {
    path: "/cfg",
    name: "Config",
    component: Config
  },
  {
    path: "/i",
    name: "ImportForm",
    component: ImportForm
  },
  {
    path: "/i/:id",
    name: "ImportResults",
    component: ImportResults
  },
  {
    path: "/about",
    name: "About",
    component: About
  },
  {
    path: "/help",
    name: "Help",
    component: Help
  },
  {
    path: "/testbed",
    name: "Testbed",
    component: Testbed
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
