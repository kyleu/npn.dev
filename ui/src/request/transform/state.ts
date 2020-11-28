import {socketRef} from "@/socket/socket";
import {TransformResult} from "@/request/transform/transformResult";
import {ref} from "@vue/composition-api";
import {requestService} from "@/util/services";
import {clientCommands} from "@/util/command";
import {clearPendingRequest, pendingRequestsRef, setPendingRequest} from "@/socket/pending";
import {activeSessionRef} from "@/session/state";
import {requestEditingRef} from "@/request/state";

export const transformResultRef = ref<TransformResult>();

export function getTransformResult(coll: string, req: string, fmt: string): TransformResult | undefined {
  const v = transformResultRef.value;
  if (v && v.coll === coll && v.req === req && v.fmt === fmt) {
    return v;
  }
  if (requestEditingRef.value && socketRef.value && fmt.length > 0 && setPendingRequest(pendingRequestsRef, "transform", `${coll}::${req}::${fmt}`)) {
    const param = {coll, req, sess: activeSessionRef.value, fmt, proto: requestEditingRef.value.prototype};
    socketRef.value.send({svc: requestService.key, cmd: clientCommands.transform, param});
  }

  return undefined;
}

export function setTransformResult(r: TransformResult): void {
  clearPendingRequest(pendingRequestsRef, "transform", `${r.coll}::${r.req}::${r.fmt}`);
  transformResultRef.value = r;
}
