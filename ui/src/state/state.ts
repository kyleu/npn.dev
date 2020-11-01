import {Collection} from "@/collection/collection";
import Profile from "@/user/profile";
import {NPNRequest, Summary} from "@/request/model";
import {ActiveRequest} from "@/state/store";
import {CallResult} from "@/call/model";
import {Socket} from "@/socket/socket";
import {TransformResult} from "@/request/transformResult";
import VueCompositionAPI, {ref} from "@vue/composition-api";
import Vue from "vue";

Vue.use(VueCompositionAPI);

export interface Breadcrumb {
  readonly title: string;
  readonly path: string;
}

interface CollectionData<T> {
  readonly key: string;
  requests: T[];
}

export const hostRef = ref<string>("");
export const breadcrumbsRef = ref<Breadcrumb[]>();
export const profileRef = ref<Profile>();

export const collectionsRef = ref<Collection[]>([]);

export const collectionSummariesRef = ref<CollectionData<Summary>[]>([]);
export const requestDetailsRef = ref<CollectionData<NPNRequest>[]>([]);

export const activeRequestRef = ref<ActiveRequest>();
export const requestOriginalRef = ref<NPNRequest>();
export const requestEditingRef = ref<NPNRequest>();

export const callResultRef = ref<CallResult>();
export const transformResultRef = ref<TransformResult>();

export const socketRef = ref<Socket>();

export function getCollection(key: string): Collection | undefined {
  for (const c of collectionsRef.value) {
    if (c.key === key) {
      return c;
    }
  }
}

export function getCollectionRequestSummaries(key: string): Summary[] | undefined {
  for (const c of collectionSummariesRef.value) {
    if (c.key === key) {
      return c.requests;
    }
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

function getCollectionRequestDetails(key: string): NPNRequest[] | undefined {
  for (const c of requestDetailsRef.value) {
    if (c.key === key) {
      return c.requests;
    }
  }
  return undefined;
}

function setCollectionRequestDetails(key: string, requests: NPNRequest[]): void {
  for (const c of requestDetailsRef.value) {
    if (c.key === key) {
      c.requests = requests;
      return;
    }
  }
  requestDetailsRef.value.push({key, requests})
}

export function getRequestDetail(coll: string, req: string): NPNRequest | undefined {
  for (const r of getCollectionRequestDetails(coll) || []) {
    if (r.key == req) {
      return r;
    }
  }
  return undefined;
}

export function setRequestDetail(coll: string, req: NPNRequest): void {
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
  setCollectionRequestDetails(coll, rs);
}
