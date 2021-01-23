import {RequestCompleted, RequestResults, RequestStarted} from "@/call/model";
import {socketRef} from "@/socket/socket";
import {clearPendingRequests, pendingRequestsRef, setPendingRequests} from "@/socket/pending";
import {activeSessionRef} from "@/session/state";
import {requestService} from "@/util/services";
import {clientCommands} from "@/util/command";
import {requestEditingRef} from "@/request/state";
import {ref} from "@vue/composition-api";
import {logError} from "@/util/log";

export const requestResultsRef = ref<RequestResults>({id: "", coll: "", req: "", cycles: []});

export function getRequestResults(coll: string, req: string): RequestResults | undefined {
  const v = requestResultsRef.value;
  if (v && v.coll === coll && v.req === req) {
    return v;
  }
  if (requestEditingRef.value && socketRef.value && req.length > 0) {
    if (setPendingRequests(pendingRequestsRef, "call", `${coll}::${req}`)) {
      const param = {coll, req, sess: activeSessionRef.value, proto: requestEditingRef.value.prototype};
      socketRef.value.send({channel: requestService.key, cmd: clientCommands.call, param});
    }
  }

  return undefined;
}

export function onRequestStarted(r: RequestStarted): void {
  let v = requestResultsRef.value;
  if(v.id !== r.id) {
    // console.log("update: id was [" + v.id + "], now [" + r.id + "]");
    requestResultsRef.value = {id: r.id, coll: r.coll, req: r.req, cycles: []};
    v = requestResultsRef.value;
  }
  v.cycles.push({
    idx: r.idx,
    method: r.method,
    url: r.url,
    started: r.started,

    status: "pending",
    rsp: null,
    error: "",
    duration: 0
  });
}

export function onRequestCompleted(r: RequestCompleted): void {
  const v = requestResultsRef.value;
  if(v.cycles.length < (r.idx + 1)) {
    logError("request for idx [" + r.idx + "] found length [" + v.cycles.length + "]");
  }
  const c = v.cycles[r.idx || 0];
  c.status = r.status;
  c.rsp = r.rsp;
  c.error = r.error;
  c.duration = r.duration;
  clearPendingRequests(pendingRequestsRef, "call", `${r.coll}::${r.req}`);
}
