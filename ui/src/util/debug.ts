import {activeRequestRef, requestEditingRef, requestOriginalRef} from "@/request/state";
import {breadcrumbsRef} from "@/layout/breadcrumb";
import {collectionsRef, collectionSummariesRef} from "@/collection/state";
import {pendingRequestsRef} from "@/socket/pending";
import {profileRef} from "@/user/profile";
import {jsonClone} from "@/util/json";
import Vue from "vue";
import {VueRouter} from "vue-router/types/router";
import {sessionDetailsRef, sessionEditingRef, sessionOriginalRef, sessionSummariesRef} from "@/session/state";
import {requestTransformResultRef} from "@/transform/state";
import {bodyConfigRef} from "@/body/state";
import {requestDetailsRef} from "@/collection/requestDetails";
import {requestResultsRef} from "@/call/state";
import {hostRef} from "@/socket/socket";

export interface NPNDebug {
  root: Vue;
  router: VueRouter;
  onDebug: () => void;
}

export function onDebug(): void {
  const ret = {
    host: hostRef.value,
    profile: profileRef.value,
    breadcrumbs: breadcrumbsRef.value,
    session: {
      sessions: sessionDetailsRef.value,
      sessionSummaries: sessionSummariesRef.value,
      sessionEditing: sessionEditingRef.value,
      sessionOriginal: sessionOriginalRef.value
    },
    collection: {
      collections: collectionsRef.value,
      collectionSummaries: collectionSummariesRef.value
    },
    request: {
      activeRequest: activeRequestRef.value,
      pendingRequests: pendingRequestsRef.value,
      requestDetails: requestDetailsRef.value,
      requestEditing: requestEditingRef.value,
      requestOriginal: requestOriginalRef.value,
      requestResultsRef: requestResultsRef.value,
      transformResult: requestTransformResultRef.value
    },
    body: bodyConfigRef.value
  };

  console.debug(`debug output at [${new Date().toString()}]`);
  console.debug(jsonClone(ret));
}
