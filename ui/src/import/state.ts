import {ref} from "@vue/composition-api";
import {getCollectionRequestDetails} from "@/collection/requestDetails";
import {jsonClone} from "@/util/json";
import {authConfigRef, toAuthConfig} from "@/auth/state";
import {bodyConfigRef, toBodyConfig} from "@/body/state";
import {getCollectionRequestSummaries} from "@/collection/state";
import {socketRef} from "@/socket/socket";
import {pendingRequestsRef, setPendingRequests} from "@/socket/pending";
import {requestService, systemService} from "@/util/services";
import {clientCommands} from "@/util/command";
import {activeRequestRef, requestEditingRef, requestOriginalRef} from "@/request/state";
import {ImportResult} from "@/import/model";

export function setActiveImport(id: string): void {
  importResultIDRef.value = id;
  const r = importResultRef.value;
  const req = (!r) || r.id !== id;
  if (req && socketRef.value) {
    if (setPendingRequests(pendingRequestsRef, "import", id)) {
      socketRef.value.send({channel: systemService.key, cmd: clientCommands.getImport, param: id});
    }
  }
}

export const importResultIDRef = ref<string>("");
export const importResultRef = ref<ImportResult>({id: ""});
