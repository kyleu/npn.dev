import {CallResult} from "@/call/model";
import {socketRef} from "@/socket/socket";
import {clearPendingRequest, pendingRequestsRef, setPendingRequest} from "@/socket/pending";
import {activeSessionRef} from "@/session/state";
import {requestService} from "@/util/services";
import {clientCommands} from "@/util/command";
import {requestEditingRef} from "@/request/state";
import {ref} from "@vue/composition-api";

export const callResultRef = ref<CallResult>();

export function getCallResult(coll: string, req: string): CallResult | undefined {
  const v = callResultRef.value;
  if (v && v.collection === coll && v.request === req) {
    return v;
  }
  if (requestEditingRef.value && socketRef.value && req.length > 0 && setPendingRequest(pendingRequestsRef, "call", `${coll}::${req}`)) {
    const param = {coll, req, sess: activeSessionRef.value, proto: requestEditingRef.value.prototype};
    socketRef.value.send({svc: requestService.key, cmd: clientCommands.call, param});
  }

  return undefined;
}

export function setCallResult(r: CallResult): void {
  clearPendingRequest(pendingRequestsRef, "call", `${r.collection}::${r.request}`);
  callResultRef.value = r;
}
