import {serverCommands} from "@/util/command";
import {logWarn} from "@/util/log";
import {importResultRef} from "@/import/state";

// @ts-ignore
// eslint-disable-next-line
export function onImportMessage(cmd: string, param: any): void {
  switch (cmd) {
    case serverCommands.importResult:
      importResultRef.value = param;
      break;
    default:
      logWarn("unhandled import message [" + cmd + "]", param);
  }
}

