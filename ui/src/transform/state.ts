import {socketRef} from "@/socket/socket";
import {CollectionTransformResult, RequestTransformResult, SessionTransformResult} from "@/transform/result";
import {ref} from "@vue/composition-api";
import {collectionService, requestService, sessionService} from "@/util/services";
import {clientCommands} from "@/util/command";
import {clearPendingRequests, pendingRequestsRef, setPendingRequests} from "@/socket/pending";
import {activeSessionRef} from "@/session/state";
import {requestEditingRef} from "@/request/state";

export const requestTransformResultRef = ref<RequestTransformResult>();

export function getRequestTransformResult(coll: string, req: string, fmt: string): RequestTransformResult | undefined {
  const v = requestTransformResultRef.value;
  if (v && v.coll === coll && v.req === req && v.fmt === fmt) {
    return v;
  }
  if (requestEditingRef.value && socketRef.value && fmt.length > 0 && setPendingRequests(pendingRequestsRef, "transform", `${coll}::${req}::${fmt}`)) {
    const param = {coll, req, sess: activeSessionRef.value, fmt, proto: requestEditingRef.value.prototype};
    socketRef.value.send({channel: requestService.key, cmd: clientCommands.transform, param});
  }

  return undefined;
}

export function setRequestTransformResult(r: RequestTransformResult): void {
  clearPendingRequests(pendingRequestsRef, "transform", `${r.coll}::${r.req}::${r.fmt}`);
  requestTransformResultRef.value = r;
}

export const collectionTransformResultRef = ref<CollectionTransformResult>();

export function getCollectionTransformResult(coll: string, fmt: string): CollectionTransformResult | undefined {
  const v = collectionTransformResultRef.value;
  if (v && v.coll === coll && v.fmt === fmt) {
    return v;
  }
  if (socketRef.value && fmt.length > 0 && setPendingRequests(pendingRequestsRef, "export-collection", `${coll}::${fmt}`)) {
    const param = {coll, fmt};
    socketRef.value.send({channel: collectionService.key, cmd: clientCommands.transform, param});
  }

  return undefined;
}

export function setCollectionTransformResult(r: CollectionTransformResult): void {
  clearPendingRequests(pendingRequestsRef, "export-collection", `${r.coll}::${r.fmt}`);
  collectionTransformResultRef.value = r;
}

export const sessionTransformResultRef = ref<SessionTransformResult>();

export function getSessionTransformResult(sess: string): SessionTransformResult | undefined {
  const v = sessionTransformResultRef.value;
  if (v && v.key === sess) {
    return v;
  }
  if (socketRef.value && sess.length > 0 && setPendingRequests(pendingRequestsRef, "export-session", sess)) {
    socketRef.value.send({channel: sessionService.key, cmd: clientCommands.transform, param: sess});
  }

  return undefined;
}

export function setSessionTransformResult(r: SessionTransformResult): void {
  clearPendingRequests(pendingRequestsRef, "export-session", r.key);
  sessionTransformResultRef.value = r;
}
