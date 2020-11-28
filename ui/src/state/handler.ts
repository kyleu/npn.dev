import {setRequestDetail} from "@/request/state";
import {Message} from "@/socket/socket";
import {isDebug, logDebug, logWarn, onLog} from "@/util/log";
import {serverCommands} from "@/util/command";
import {onSessionAdded, onSessionDeleted, onSessionNotFound, sessionSummariesRef, setSessionDetail} from "@/session/state";
import {
  collectionsRef,
  onCollectionAdded,
  onCollectionDeleted,
  onCollectionNotFound,
  onCollectionUpdated,
  onRequestAdded,
  onRequestDeleted,
  onRequestNotFound,
  setCollectionRequestSummaries
} from "@/collection/state";
import {jsonClone} from "@/util/json";
import {setCallResult} from "@/call/state";
import {setTransformResult} from "@/request/transform/state";

export const messageHandler = (msg: Message): void => {
  if (isDebug()) {
    logDebug("IN: " + msg.cmd, jsonClone(msg.param));
  }

  switch (msg.cmd) {
    case serverCommands.log:
      onLog(msg.param);
      break;

    // Collections
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

    // Requests
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
      onRequestNotFound(msg.param.coll);
      break;
    case serverCommands.callResult:
      setCallResult(msg.param);
      break;
    case serverCommands.transformResult:
      setTransformResult(msg.param);
      break;

    // Sessions
    case serverCommands.sessions:
      sessionSummariesRef.value = msg.param;
      break;
    case serverCommands.sessionAdded:
      onSessionAdded(msg.param);
      break;
    case serverCommands.sessionDeleted:
      onSessionDeleted(msg.param);
      break;
    case serverCommands.sessionDetail:
      setSessionDetail(msg.param);
      break;
    case serverCommands.sessionNotFound:
      onSessionNotFound();
      break;

    default:
      logWarn("unhandled message [" + msg.cmd + "]", msg);
  }
};
