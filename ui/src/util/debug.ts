import {activeRequestRef, callResultRef, requestEditingRef, requestOriginalRef, transformResultRef} from "@/request/state";
import {breadcrumbsRef} from "@/layout/breadcrumb";
import {collectionsRef, collectionSummariesRef, requestDetailsRef} from "@/collection/state";
import {pendingRequestsRef} from "@/socket/pending";
import {profileRef} from "@/user/profile";
import {jsonParse, jsonStr} from "@/util/json";
import Vue from "vue";
import {VueRouter} from "vue-router/types/router";

export interface NPNDebug {
  root: Vue;
  router: VueRouter;
  debug: () => void;
}

export function debug(): void {
  const ret = {
    activeRequest: activeRequestRef.value,
    breadcrumbs: breadcrumbsRef.value,
    callResult: callResultRef.value,
    collections: collectionsRef.value,
    collectionSummaries: collectionSummariesRef.value,
    pendingRequests: pendingRequestsRef.value,
    profile: profileRef.value,
    requestDetails: requestDetailsRef.value,
    requestEditing: requestEditingRef.value,
    requestOriginal: requestOriginalRef.value,
    transformResult: transformResultRef.value
  };
  console.log(`debug output at [${new Date().toString()}]`)
  console.log(jsonParse(jsonStr(ret)));
}
