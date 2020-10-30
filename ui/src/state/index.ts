import {Collection} from "@/collection/collection";
import Profile from "@/user/profile";
import {NPNRequest, Summary} from "@/request/model";

export interface Breadcrumb {
  readonly title: string;
  readonly path: string;
}

interface CollectionData<T> {
  readonly key: string;
  requests: T[];
}

export class State {
  readonly url: string;
  profile: Profile;
  breadcrumbs: Breadcrumb[] = [];

  collections: Collection[] = [];
  collectionSummaries: CollectionData<Summary>[] = [];
  requestDetails: CollectionData<NPNRequest>[] = [];

  constructor(url: string, profile: Profile) {
    this.url = url;
    this.profile = profile;
  }

  getCollection(key: string): Collection | undefined {
    for (const c of this.collections) {
      if (c.key === key) {
        return c;
      }
    }
  }

  getCollectionRequestSummaries(key: string): Summary[] | undefined {
    for (const c of this.collectionSummaries) {
      if (c.key === key) {
        return c.requests;
      }
    }
    return undefined;
  }

  setCollectionRequestSummaries(key: string, reqs: Summary[]): void {
    for (const c of this.collectionSummaries) {
      if (c.key === key) {
        c.requests = reqs;
        return;
      }
    }
    this.collectionSummaries.push({key: key, requests: reqs});
  }

  getRequestSummary(coll: string, req: string): Summary | undefined {
    for (const r of this.getCollectionRequestSummaries(coll) || []) {
      if (r.key == req) {
        return r;
      }
    }
    return undefined;
  }

  getCollectionRequestDetails(key: string): NPNRequest[] | undefined {
    for (const c of this.requestDetails) {
      if (c.key === key) {
        return c.requests;
      }
    }
    return undefined;
  }

  setCollectionRequestDetails(key: string, requests: NPNRequest[]): void {
    for (const c of this.requestDetails) {
      if (c.key === key) {
        c.requests = requests;
        return;
      }
    }
    this.requestDetails.push({key, requests})
  }

  getRequestDetail(coll: string, req: string): NPNRequest | undefined {
    for (const r of this.getCollectionRequestDetails(coll) || []) {
      if (r.key == req) {
        return r;
      }
    }
    return undefined;
  }

  setRequestDetail(coll: string, req: NPNRequest): void {
    const rs = this.getCollectionRequestDetails(coll) || []
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
    this.setCollectionRequestDetails(coll, rs);
  }
}

interface InitialData {
  readonly url: string;
  readonly profile: Profile;
}

export function initialState(): State {
  // @ts-ignore
  // eslint-disable-next-line
  const cfg = (window as any).initialData as InitialData;

  let profile = {} as Profile;
  const url = cfg && cfg.url ? cfg.url : "";

  if (cfg && cfg.profile) {
    profile = cfg.profile;
  }

  return new State(url, profile);
}
