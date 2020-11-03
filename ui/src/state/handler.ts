import {setCallResult, setRequestDetail, setTransformResult} from "@/request/state";
import {Message} from "@/socket/socket";
import {logDebug, logWarn} from "@/util/log";
import {serverCommands} from "@/util/command";
import {collectionsRef, onCollectionAdded, setCollectionRequestSummaries} from "@/collection/state";

export const messageHandler = (msg: Message): void => {
  logDebug("IN", msg);
  switch (msg.cmd) {
    case serverCommands.collections:
      collectionsRef.value = msg.param;
      break;
    case serverCommands.collectionAdded:
      onCollectionAdded(msg.param.active, msg.param.collections);
      break;
    case serverCommands.collectionDetail:
      setCollectionRequestSummaries(msg.param.key, msg.param.requests);
      break;
    case serverCommands.requestDetail:
      setRequestDetail(msg.param.coll, msg.param.req);
      break;
    case serverCommands.callResult:
      setCallResult(msg.param);
      break;
    case serverCommands.transformResult:
      setTransformResult(msg.param);
      break;
    default:
      logWarn("unhandled message [" + msg.cmd + "]", msg);
  }
}
