import {logWarn} from "@/util/log";
import {serverCommands} from "@/util/command";
import {onSessionAdded, onSessionDeleted, onSessionNotFound, sessionSummariesRef, setSessionDetail} from "@/session/state";
import {setSessionTransformResult} from "@/transform/state";
import {SessAdded, Session, SessionSummary} from "@/session/model";
import {SessionTransformResult} from "@/transform/result";

export function onSessionMessage(cmd: string, param: unknown): void {
  switch (cmd) {
    case serverCommands.sessions:
      sessionSummariesRef.value = param as SessionSummary[];
      break;
    case serverCommands.sessionAdded:
      onSessionAdded(param as SessAdded);
      break;
    case serverCommands.sessionDeleted:
      onSessionDeleted(param as string);
      break;
    case serverCommands.sessionDetail:
      setSessionDetail(param as Session);
      break;
    case serverCommands.sessionNotFound:
      onSessionNotFound();
      break;
    case serverCommands.sessionTransform:
      setSessionTransformResult(param as SessionTransformResult);
      break;
    default:
      logWarn("unhandled session message [" + cmd + "]", param);
  }
}
