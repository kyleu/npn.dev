import {LogMessage, logWarn, onLog} from "@/util/log";
import {serverCommands} from "@/util/command";
import {onSearchResults, SearchResult} from "@/search/state";

export function onSystemMessage(cmd: string, param: unknown): void {
  switch (cmd) {
    case serverCommands.log:
      onLog(param as LogMessage);
      break;
    case serverCommands.searchResults:
      onSearchResults(param as SearchResult[]);
      break;
    default:
      logWarn("unhandled system message [" + cmd + "]", param);
  }
}
