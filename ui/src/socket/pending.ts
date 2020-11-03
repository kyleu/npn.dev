import {ref, Ref} from "@vue/composition-api";

export interface PendingRequest {
  readonly t: string;
  readonly k: string;
}

export const pendingRequestsRef = ref<PendingRequest[]>([]);

export function hasPendingRequest(reqs: Ref<PendingRequest[]>, t: string, k: string): boolean {
  for (const pr of reqs.value) {
    if (pr.t === t && pr.k === k) {
      return true;
    }
  }
  return false;
}

export function setPendingRequest(reqs: Ref<PendingRequest[]>, t: string, k: string): boolean {
  if (hasPendingRequest(reqs, t, k)) {
    return false;
  }
  reqs.value.push({t, k});
  return true
}

export function clearPendingRequest(reqs: Ref<PendingRequest[]>, t: string, k: string): boolean {
  if (!hasPendingRequest(reqs, t, k)) {
    return false;
  }
  reqs.value = reqs.value.filter(x => x.t !== t && x.k !== k);
  return true
}
