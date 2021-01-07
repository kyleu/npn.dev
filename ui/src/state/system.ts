import {logWarn, onLog} from "@/util/log";
import {serverCommands} from "@/util/command";
import {onSearchResults} from "@/search/state";

// @ts-ignore
// eslint-disable-next-line
export function onSystemMessage(cmd: string, param: any): void {
  switch (cmd) {
    case serverCommands.log:
      onLog(param);
      break;
    case serverCommands.searchResults:
      onSearchResults(param);
      break;
    default:
      logWarn("unhandled system message [" + cmd + "]", param);
  }
}
