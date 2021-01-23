import {normalize, NPNRequest} from "@/request/model";
import {socketRef} from "@/socket/socket";
import {ref} from "@vue/composition-api";
import {requestService} from "@/util/services";
import {clientCommands} from "@/util/command";
import {clearPendingRequests, pendingRequestsRef, setPendingRequests} from "@/socket/pending";
import {getCollectionRequestSummaries} from "@/collection/state";
import {jsonClone} from "@/util/json";
import {authConfigRef, toAuthConfig} from "@/auth/state";
import {bodyConfigRef, toBodyConfig} from "@/body/state";
import {getCollectionRequestDetails, setCollectionRequestDetails} from "@/collection/requestDetails";

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
    if (setPendingRequests(pendingRequestsRef, "request", coll + "::" + req)) {
      socketRef.value.send({channel: requestService.key, cmd: clientCommands.getRequest, param: activeRequestRef.value});
    }
  }
}

export function setRequestDetail(coll: string, origKey: string, req: NPNRequest): void {
  req = normalize(req);
  clearPendingRequests(pendingRequestsRef, "request", coll + "::" + req.key);
  const rs = getCollectionRequestDetails(coll) || [];
  let matched = false;
  for (const r in rs) {
    if (rs[r].key === origKey) {
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
