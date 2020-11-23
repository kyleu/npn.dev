import {setCallResult, setRequestDetail, setTransformResult} from "@/request/state";
import {Message} from "@/socket/socket";
import {logDebug, logWarn} from "@/util/log";
import {serverCommands} from "@/util/command";
import {
  collectionsRef,
  onCollectionAdded,
  onCollectionDeleted,
  onCollectionNotFound, onCollectionUpdated,
  onRequestAdded,
  onRequestDeleted, onRequestNotFound,
  setCollectionRequestSummaries
} from "@/collection/state";
import {sessionsRef} from "@/session/session";

export const messageHandler = (msg: Message): void => {
  logDebug("IN", msg);
  switch (msg.cmd) {
    case serverCommands.sessions:
      sessionsRef.value = msg.param;
      break;
    case serverCommands.collections:
      collectionsRef.value = msg.param;
      break;
    case serverCommands.collectionAdded:
      onCollectionAdded(msg.param.active, msg.param.collections);
      break;
    case serverCommands.collectionUpdated:
      onCollectionUpdated(msg.param);
      break;
    case serverCommands.collectionDeleted:
      onCollectionDeleted(msg.param);
      break;
    case serverCommands.collectionDetail:
      setCollectionRequestSummaries(msg.param.key, msg.param.requests);
      break;
    case serverCommands.collectionNotFound:
      onCollectionNotFound(msg.param);
      break;
    case serverCommands.requestAdded:
      onRequestAdded(msg.param.coll, msg.param.req);
      break;
    case serverCommands.requestDeleted:
      onRequestDeleted(msg.param);
      break;
    case serverCommands.requestDetail:
      setRequestDetail(msg.param.coll, msg.param.req);
      break;
    case serverCommands.requestNotFound:
      onRequestNotFound();
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
