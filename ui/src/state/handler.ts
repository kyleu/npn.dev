import {State} from "@/state/state";
import {Message} from "@/socket/socket";
import {logDebug, logWarn} from "@/util/log";
import {Collection} from "@/collection/collection";
import {cloneRequest} from "@/request/model";
import {CallResult} from "@/call/model";
import {TransformResult} from "@/request/transformResult";

export const messageHandler = (state: State, msg: Message): void => {
  logDebug("IN", msg);
  switch (msg.cmd) {
    case "collections":
      state.collections = msg.param as Collection[];
      break;
    case "collectionDetail":
      state.setCollectionRequestSummaries(msg.param.key, msg.param.requests);
      break;
    case "requestDetail":
      state.setRequestDetail(msg.param.coll, msg.param.req);
      if (msg.param.req.key === state.activeRequest?.req && msg.param.coll === state.activeRequest?.coll) {
        state.requestOriginal = cloneRequest(msg.param.req);
        state.requestEditing = msg.param.req;
      }
      break;
    case "callResult":
      state.setCallResult(msg.param as CallResult);
      break;
    case "transformResult":
      state.setTransformResult(msg.param as TransformResult);
      break;
    default:
      logWarn("unhandled message [" + msg.cmd + "]", msg);
  }
}
