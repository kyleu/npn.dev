import {setRequestDetail} from "@/request/state";
import {logWarn} from "@/util/log";
import {serverCommands} from "@/util/command";
import {onRequestCompleted, onRequestStarted} from "@/call/state";
import {setRequestTransformResult} from "@/transform/state";
import {onRequestAdded, onRequestDeleted, onRequestNotFound, RequestAdded, RequestDeleted} from "@/collection/requestDetails";
import {RequestTransformResult} from "@/transform/result";
import {RequestCompleted, RequestStarted} from "@/call/model";
import {NPNRequest} from "@/request/model";

export function onRequestMessage(cmd: string, param: unknown): void {
  switch (cmd) {
    case serverCommands.requestAdded:
      onRequestAdded((param as { coll: RequestAdded}).coll, (param as { req: NPNRequest}).req);
      break;
    case serverCommands.requestDeleted:
      onRequestDeleted(param as RequestDeleted);
      break;
    case serverCommands.requestDetail:
      setRequestDetail((param as { coll: string}).coll, (param as { origKey: string}).origKey, (param as { req: NPNRequest}).req);
      break;
    case serverCommands.requestNotFound:
      onRequestNotFound((param as { coll: string}).coll);
      break;
    case serverCommands.requestTransform:
      setRequestTransformResult(param as RequestTransformResult);
      break;
    case serverCommands.requestStarted:
      onRequestStarted(param as RequestStarted);
      break;
    case serverCommands.requestCompleted:
      onRequestCompleted(param as RequestCompleted);
      break;
    default:
      logWarn("unhandled request message [" + cmd + "]", param);
  }
}
