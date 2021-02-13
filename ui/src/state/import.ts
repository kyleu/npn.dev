import {serverCommands} from "@/util/command";
import {logWarn} from "@/util/log";
import {importResultRef} from "@/import/state";
import {ImportResult} from "@/import/model";

export function onImportMessage(cmd: string, param: unknown): void {
  switch (cmd) {
    case serverCommands.importResult:
      importResultRef.value = param as ImportResult;
      break;
    default:
      logWarn("unhandled import message [" + cmd + "]", param);
  }
}

