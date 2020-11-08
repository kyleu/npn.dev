import {Collection, CollectionCount} from "@/collection/collection";
import {NPNRequest, Summary} from "@/request/model";
import {socketRef} from "@/socket/socket";
import {ref} from "@vue/composition-api";
import {collectionService} from "@/util/services";
import {clientCommands} from "@/util/command";
import {clearPendingRequest, pendingRequestsRef, setPendingRequest} from "@/socket/pending";
import {setRequestDetail} from "@/request/state";
import {globalRouter} from "@/util/vutils";

interface CollectionData<T> {
  readonly key: string;
  requests: T[];
}

export const collectionsRef = ref<CollectionCount[]>([]);

export const collectionSummariesRef = ref<CollectionData<Summary>[]>([]);
export const requestDetailsRef = ref<CollectionData<NPNRequest>[]>([]);

export function getCollection(key: string): Collection | undefined {
  for (const c of collectionsRef.value) {
    if (c.coll.key === key) {
      return c.coll;
    }
  }
}

export function onCollectionAdded(active: string, colls: CollectionCount[]): void {
  collectionsRef.value = colls;
  globalRouter().push({name: "CollectionDetail", params: {coll: active}});
}

export function onCollectionDeleted(param: string): void {
  collectionsRef.value = collectionsRef.value.filter(x => x.coll.key !== param);
  collectionSummariesRef.value = collectionSummariesRef.value.filter(x => x.key !== param);
  requestDetailsRef.value = requestDetailsRef.value.filter(x => x.key !== param);

  globalRouter().push({name: "CollectionIndex"});
}

export function getCollectionRequestSummaries(key: string): Summary[] | undefined {
  for (const c of collectionSummariesRef.value) {
    if (c.key === key) {
      return c.requests;
    }
  }
  if (socketRef.value && key.length > 0 && setPendingRequest(pendingRequestsRef, "collection", key)) {
    socketRef.value.send({svc: collectionService.key, cmd: clientCommands.getCollection, param: key});
  }

  return undefined;
}

export function setCollectionRequestSummaries(key: string, reqs: Summary[]): void {
  for (const c of collectionSummariesRef.value) {
    if (c.key === key) {
      c.requests = reqs;
      return;
    }
  }
  collectionSummariesRef.value.push({key: key, requests: reqs});
}

export function getCollectionRequestDetails(key: string): NPNRequest[] | undefined {
  for (const c of requestDetailsRef.value) {
    if (c.key === key) {
      return c.requests;
    }
  }
  return undefined;
}

export function setCollectionRequestDetails(key: string, requests: NPNRequest[]): void {
  clearPendingRequest(pendingRequestsRef, "collection", key);
  for (const c of requestDetailsRef.value) {
    if (c.key === key) {
      c.requests = requests;
      return;
    }
  }
  requestDetailsRef.value.push({key, requests});
}

interface RequestAdded {
  key: string;
  requests: Summary[];
}

export function onRequestAdded(coll: RequestAdded, req: NPNRequest): void {
  setCollectionRequestSummaries(coll.key, coll.requests);
  setRequestDetail(coll.key, req);
  globalRouter().push({name: "RequestDetail", params: {coll: coll.key, req: req.key}});
}

interface RequestDeleted {
  req: string;
  coll: string;
  requests: Summary[];
}

export function onRequestDeleted(rd: RequestDeleted): void {
  setCollectionRequestSummaries(rd.coll, rd.requests);
  globalRouter().push({name: "CollectionDetail", params: {coll: rd.coll}});
}

