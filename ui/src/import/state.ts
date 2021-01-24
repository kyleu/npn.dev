import {ref} from "@vue/composition-api";
import {socketRef} from "@/socket/socket";
import {pendingRequestsRef, setPendingRequests} from "@/socket/pending";
import {importService} from "@/util/services";
import {clientCommands} from "@/util/command";
import {ImportResult} from "@/import/model";

export const importResultIDRef = ref<string>("");
export const importResultRef = ref<ImportResult>({key: ""});

export function setActiveImport(id: string): void {
  importResultIDRef.value = id;
  const r = importResultRef.value;
  const req = (!r) || r.key !== id;
  if (req && socketRef.value) {
    if (setPendingRequests(pendingRequestsRef, "import", id)) {
      socketRef.value.send({channel: importService.key, cmd: clientCommands.getImport, param: id});
    }
  }
}
