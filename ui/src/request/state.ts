import {NPNRequest} from "@/request/model";
import {socketRef} from "@/socket/socket";
import {ref} from "@vue/composition-api";
import {requestService} from "@/util/services";
import {clientCommands} from "@/util/command";
import {clearPendingRequest, pendingRequestsRef, setPendingRequest} from "@/socket/pending";
import {getCollectionRequestDetails, getCollectionRequestSummaries, setCollectionRequestDetails} from "@/collection/state";
import {jsonClone} from "@/util/json";
import {authConfigRef, toAuthConfig} from "@/auth/state";
import {bodyConfigRef, toBodyConfig} from "@/body/state";

export interface ActiveRequest {
  readonly coll: string;
  readonly req: string;
}

export const activeRequestRef = ref<ActiveRequest>();
export const requestOriginalRef = ref<NPNRequest>();
export const requestEditingRef = ref<NPNRequest>();

export function setActiveRequest(coll: string, req: string): void {
  activeRequestRef.value = {coll, req};
  for (const r of getCollectionRequestDetails(coll) || []) {
    if (r.key === req) {
      if(!requestOriginalRef.value || (requestOriginalRef.value.key !== req)) {
        requestOriginalRef.value = r;
        requestEditingRef.value = jsonClone(r);
        authConfigRef.value = toAuthConfig(r.prototype.auth);
        bodyConfigRef.value = toBodyConfig(r.prototype.body);
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

function normalize(r: NPNRequest): NPNRequest {
  if (!r.prototype) {
    r.prototype = {domain: "", method: "", protocol: ""};
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
  return r;
}

export function setRequestDetail(coll: string, req: NPNRequest): void {
  req = normalize(req);
  clearPendingRequest(pendingRequestsRef, "request", coll + "::" + req.key);
  const rs = getCollectionRequestDetails(coll) || [];
  let matched = false;
  for (const r in rs) {
    if (rs[r].key === req.key) {
      matched = true;
      rs[r] = req;
    }
  }
  if (!matched) {
    rs.push(req);
  }

  setCollectionRequestDetails(coll, rs.map(normalize));

  if (activeRequestRef.value && req.key === activeRequestRef.value.req && coll === activeRequestRef.value.coll) {
    requestOriginalRef.value = req;
    requestEditingRef.value = jsonClone(req);
    authConfigRef.value = toAuthConfig(requestEditingRef.value?.prototype.auth);
    bodyConfigRef.value = toBodyConfig(requestEditingRef.value?.prototype.body);
  }
}
