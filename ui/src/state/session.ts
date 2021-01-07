import {logWarn} from "@/util/log";
import {serverCommands} from "@/util/command";
import {onSessionAdded, onSessionDeleted, onSessionNotFound, sessionSummariesRef, setSessionDetail} from "@/session/state";
import {setSessionTransformResult} from "@/transform/state";

// @ts-ignore
// eslint-disable-next-line
export function onSessionMessage(cmd: string, param: any): void {
  switch (cmd) {
    case serverCommands.sessions:
      sessionSummariesRef.value = param;
      break;
    case serverCommands.sessionAdded:
      onSessionAdded(param);
      break;
    case serverCommands.sessionDeleted:
      onSessionDeleted(param);
      break;
    case serverCommands.sessionDetail:
      setSessionDetail(param);
      break;
    case serverCommands.sessionNotFound:
      onSessionNotFound();
      break;
    case serverCommands.sessionTransform:
      setSessionTransformResult(param);
      break;
    default:
      logWarn("unhandled session message [" + cmd + "]", param);
  }
}
