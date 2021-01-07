import {setRequestDetail} from "@/request/state";
import {logWarn} from "@/util/log";
import {serverCommands} from "@/util/command";
import {onRequestCompleted, onRequestStarted} from "@/call/state";
import {setRequestTransformResult} from "@/transform/state";
import {onRequestAdded, onRequestDeleted, onRequestNotFound} from "@/collection/requestDetails";

// @ts-ignore
// eslint-disable-next-line
export function onRequestMessage(cmd: string, param: any): void {
  switch (cmd) {
    case serverCommands.requestAdded:
      onRequestAdded(param.coll, param.req);
      break;
    case serverCommands.requestDeleted:
      onRequestDeleted(param);
      break;
    case serverCommands.requestDetail:
      setRequestDetail(param.coll, param.origKey, param.req);
      break;
    case serverCommands.requestNotFound:
      onRequestNotFound(param.coll);
      break;
    case serverCommands.requestTransform:
      setRequestTransformResult(param);
      break;
    case serverCommands.requestStarted:
      onRequestStarted(param);
      break;
    case serverCommands.requestCompleted:
      onRequestCompleted(param);
      break;
    default:
      logWarn("unhandled request message [" + cmd + "]", param);
  }
}
