import {cloneRequest, NPNRequest} from "@/request/model";
import {CallResult} from "@/call/model";
import {socketRef} from "@/socket/socket";
import {TransformResult} from "@/request/transform/transformResult";
import {ref} from "@vue/composition-api";
import {requestService} from "@/util/services";
import {clientCommands} from "@/util/command";
import {clearPendingRequest, pendingRequestsRef, setPendingRequest} from "@/socket/pending";
import {getCollectionRequestDetails, getCollectionRequestSummaries, setCollectionRequestDetails} from "@/collection/state";
import {sessionsRef} from "@/session/session";

export interface ActiveRequest {
  readonly coll: string;
  readonly req: string;
}

export const activeRequestRef = ref<ActiveRequest>();
export const requestOriginalRef = ref<NPNRequest>();
export const requestEditingRef = ref<NPNRequest>();

export const callResultRef = ref<CallResult>();
export const transformResultRef = ref<TransformResult>();

export function setActiveRequest(coll: string, req: string): void {
  activeRequestRef.value = {coll, req};
  for (const r of getCollectionRequestDetails(coll) || []) {
    if (r.key == req) {
      if(!requestOriginalRef.value || (requestOriginalRef.value.key !== req)) {
        requestOriginalRef.value = r;
        requestEditingRef.value = cloneRequest(r);
      }
      return;
    }
  }

  getCollectionRequestSummaries(coll);

  if (req && socketRef.value) {
    if (setPendingRequest(pendingRequestsRef, "request", coll + "::" + req)) {
      socketRef.value.send({svc: requestService.key, cmd: clientCommands.getRequest, param: activeRequestRef.value});
    }
  }
}

function filterRequest(r: NPNRequest): NPNRequest {
  if (!r.prototype) {
    r.prototype = {domain: "", method: "", protocol: ""}
  }
  if(!r.prototype.query) {
    r.prototype.query = [];
  }
  if(!r.prototype.headers) {
    r.prototype.headers = [];
  }
  if(!r.prototype.body) {
    r.prototype.body = {type: "", config: {}};
  }
  return r
}

export function setRequestDetail(coll: string, req: NPNRequest): void {
  req = filterRequest(req);
  clearPendingRequest(pendingRequestsRef, "request", coll + "::" + req)
  const rs = getCollectionRequestDetails(coll) || []
  let matched = false;
  for (const r in rs) {
    if (rs[r].key == req.key) {
      matched = true;
      rs[r] = req;
    }
  }
  if (!matched) {
    rs.push(req);
  }

  setCollectionRequestDetails(coll, rs.map(filterRequest));

  if (activeRequestRef.value && req.key === activeRequestRef.value.req && coll === activeRequestRef.value.coll) {
    requestOriginalRef.value = req;
    requestEditingRef.value = cloneRequest(req);
  }
}

export function getCallResult(coll: string, req: string): CallResult | undefined {
  const v = callResultRef.value;
  if (v && v.collection === coll && v.request === req) {
    return v;
  }
  if (requestEditingRef.value && socketRef.value && req.length > 0 && setPendingRequest(pendingRequestsRef, "call", `${coll}::${req}`)) {
    const param = {coll, req, proto: requestEditingRef.value.prototype};
    socketRef.value.send({svc: requestService.key, cmd: clientCommands.call, param});
  }

  return undefined;
}

export function setCallResult(r: CallResult): void {
  clearPendingRequest(pendingRequestsRef, "call", `${r.collection}::${r.request}`)
  callResultRef.value = r;
}

export function getTransformResult(coll: string, req: string, fmt: string): TransformResult | undefined {
  const v = transformResultRef.value;
  if (v && v.coll === coll && v.req === req && v.fmt === fmt) {
    return v;
  }
  const sess = sessionsRef
  if (requestEditingRef.value && socketRef.value && fmt.length > 0 && setPendingRequest(pendingRequestsRef, "transform", `${coll}::${req}::${fmt}`)) {
    const param = {coll, req, sess, fmt, proto: requestEditingRef.value.prototype};
    socketRef.value.send({svc: requestService.key, cmd: clientCommands.transform, param});
  }

  return undefined;
}

export function setTransformResult(r: TransformResult): void {
  clearPendingRequest(pendingRequestsRef, "transform", `${r.coll}::${r.req}::${r.fmt}`)
  transformResultRef.value = r;
}
