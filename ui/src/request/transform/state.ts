import {socketRef} from "@/socket/socket";
import {CollectionTransformResult, RequestTransformResult} from "@/request/transform/result";
import {ref} from "@vue/composition-api";
import {collectionService, requestService} from "@/util/services";
import {clientCommands} from "@/util/command";
import {clearPendingRequest, pendingRequestsRef, setPendingRequest} from "@/socket/pending";
import {activeSessionRef} from "@/session/state";
import {requestEditingRef} from "@/request/state";

export const requestTransformResultRef = ref<RequestTransformResult>();

export function getRequestTransformResult(coll: string, req: string, fmt: string): RequestTransformResult | undefined {
  const v = requestTransformResultRef.value;
  if (v && v.coll === coll && v.req === req && v.fmt === fmt) {
    return v;
  }
  if (requestEditingRef.value && socketRef.value && fmt.length > 0 && setPendingRequest(pendingRequestsRef, "transform", `${coll}::${req}::${fmt}`)) {
    const param = {coll, req, sess: activeSessionRef.value, fmt, proto: requestEditingRef.value.prototype};
    socketRef.value.send({svc: requestService.key, cmd: clientCommands.transform, param});
  }

  return undefined;
}

export function setRequestTransformResult(r: RequestTransformResult): void {
  clearPendingRequest(pendingRequestsRef, "transform", `${r.coll}::${r.req}::${r.fmt}`);
  requestTransformResultRef.value = r;
}

export const collectionTransformResultRef = ref<CollectionTransformResult>();

export function getCollectionTransformResult(coll: string, fmt: string): CollectionTransformResult | undefined {
  const v = collectionTransformResultRef.value;
  if (v && v.coll === coll && v.fmt === fmt) {
    return v;
  }
  if (socketRef.value && fmt.length > 0 && setPendingRequest(pendingRequestsRef, "export-collection", `${coll}::${fmt}`)) {
    const param = {coll, fmt};
    socketRef.value.send({svc: collectionService.key, cmd: clientCommands.transform, param});
  }

  return undefined;
}

export function setCollectionTransformResult(r: CollectionTransformResult): void {
  clearPendingRequest(pendingRequestsRef, "export-collection", `${r.coll}::${r.fmt}`);
  collectionTransformResultRef.value = r;
}
