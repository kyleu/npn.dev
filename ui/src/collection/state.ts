import {Collection} from "@/collection/collection";
import {NPNRequest, Summary} from "@/request/model";
import {socketRef} from "@/socket/socket";
import {ref} from "@vue/composition-api";
import {collectionService} from "@/util/services";
import {clientCommands} from "@/util/command";
import {clearPendingRequest, pendingRequestsRef, setPendingRequest} from "@/socket/pending";
import {VueRouter} from "vue-router/types/router";

interface CollectionData<T> {
  readonly key: string;
  requests: T[];
}

export const collectionsRef = ref<Collection[]>([]);

export const collectionSummariesRef = ref<CollectionData<Summary>[]>([]);
export const requestDetailsRef = ref<CollectionData<NPNRequest>[]>([]);

export function getCollection(key: string): Collection | undefined {
  for (const c of collectionsRef.value) {
    if (c.key === key) {
      return c;
    }
  }
}

export function onCollectionAdded(active: string, colls: Collection[]): void {
  collectionsRef.value = colls;

  // @ts-ignore
  // eslint-disable-next-line
  const router = (window as any).npn.router as VueRouter;
  router.push({name: "CollectionDetail", params: {coll: active}});
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
