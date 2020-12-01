import {Collection, CollectionCount} from "@/collection/collection";
import {Summary} from "@/request/model";
import {socketRef} from "@/socket/socket";
import {ref} from "@vue/composition-api";
import {collectionService} from "@/util/services";
import {clientCommands} from "@/util/command";
import {pendingRequestsRef, setPendingRequests} from "@/socket/pending";
import {globalRouter} from "@/util/vutils";
import {requestDetailsRef} from "@/collection/requestDetails";

interface CollectionData<T> {
  readonly key: string;
  requests: T[];
}

export const collectionsRef = ref<CollectionCount[]>([]);
export const collectionSummariesRef = ref<CollectionData<Summary>[]>([]);

export function getCollection(key: string): Collection | undefined {
  for (const c of collectionsRef.value) {
    if (c.coll.key === key) {
      return c.coll;
    }
  }
}

export function onCollectionNotFound(key: string): void {
  collectionsRef.value = collectionsRef.value.filter(x => x.coll.key !== key);
  globalRouter().push({name: "CollectionIndex"});
}

export function onCollectionAdded(active: string, colls: CollectionCount[]): void {
  collectionsRef.value = colls;
  globalRouter().push({name: "CollectionDetail", params: {coll: active}});
}

export function onCollectionUpdated(coll: Collection): void {
  collectionsRef.value = collectionsRef.value.map(x => {
    if (x.coll.key === coll.key) {
      x.coll = coll;
    }
    return x;
  });
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
  if (socketRef.value && key.length > 0 && setPendingRequests(pendingRequestsRef, "collection", key)) {
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
