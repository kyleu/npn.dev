import {NPNRequest, Summary} from "@/request/model";
import {ref} from "@vue/composition-api";
import {clearPendingRequests, pendingRequestsRef} from "@/socket/pending";
import {setRequestDetail} from "@/request/state";
import {globalRouter} from "@/util/vutils";
import {setCollectionRequestSummaries} from "@/collection/state";

interface CollectionData<T> {
  readonly key: string;
  requests: T[];
}

export const requestDetailsRef = ref<CollectionData<NPNRequest>[]>([]);

export function getCollectionRequestDetails(key: string): NPNRequest[] | undefined {
  for (const c of requestDetailsRef.value) {
    if (c.key === key) {
      return c.requests;
    }
  }
  return undefined;
}

export function setCollectionRequestDetails(key: string, requests: NPNRequest[]): void {
  clearPendingRequests(pendingRequestsRef, "collection", key);
  for (const c of requestDetailsRef.value) {
    if (c.key === key) {
      c.requests = requests;
      return;
    }
  }
  requestDetailsRef.value.push({key, requests});
}

export interface RequestAdded {
  key: string;
  requests: Summary[];
}

export function onRequestAdded(coll: RequestAdded, req: NPNRequest): void {
  setCollectionRequestSummaries(coll);
  setRequestDetail(coll.key, req.key, req);
  globalRouter().push({name: "RequestDetail", params: {coll: coll.key, req: req.key}});
}

export interface RequestDeleted {
  req: string;
  coll: string;
  requests: Summary[];
}

export function onRequestDeleted(rd: RequestDeleted): void {
  setCollectionRequestSummaries({key: rd.coll, requests: rd.requests});
  globalRouter().push({name: "CollectionDetail", params: {coll: rd.coll}});
}

export function onRequestNotFound(coll: string): void {
  globalRouter().push({name: "CollectionDetail", params: {coll}});
}
