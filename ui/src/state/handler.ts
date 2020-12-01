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
  setCollectionRequestSummaries
} from "@/collection/state";
import {jsonClone} from "@/util/json";
import {onRequestCompleted, onRequestStarted} from "@/call/state";
import {setCollectionTransformResult, setRequestTransformResult} from "@/request/transform/state";
import {onRequestAdded, onRequestDeleted, onRequestNotFound} from "@/collection/requestDetails";
import {collectionService, requestService, sessionService, systemService} from "@/util/services";

// @ts-ignore
// eslint-disable-next-line
function onSystemMessage(cmd: string, param: any): void {
  switch (cmd) {
    case serverCommands.log:
      onLog(param);
      break;
    default:
      logWarn("unhandled system message [" + cmd + "]", param);
  }
}

// @ts-ignore
// eslint-disable-next-line
function onCollectionMessage(cmd: string, param: any): void {
  switch (cmd) {
    case serverCommands.collections:
      collectionsRef.value = param;
      break;
    case serverCommands.collectionAdded:
      onCollectionAdded(param.active, param.collections);
      break;
    case serverCommands.collectionUpdated:
      onCollectionUpdated(param);
      break;
    case serverCommands.collectionDeleted:
      onCollectionDeleted(param);
      break;
    case serverCommands.collectionDetail:
      setCollectionRequestSummaries(param.key, param.requests);
      break;
    case serverCommands.collectionTransform:
      setCollectionTransformResult(param);
      break;
    case serverCommands.collectionNotFound:
      onCollectionNotFound(param);
      break;
    default:
      logWarn("unhandled collection message [" + cmd + "]", param);
  }
}

// @ts-ignore
// eslint-disable-next-line
function onRequestMessage(cmd: string, param: any): void {
  switch (cmd) {
    case serverCommands.requestAdded:
      onRequestAdded(param.coll, param.req);
      break;
    case serverCommands.requestDeleted:
      onRequestDeleted(param);
      break;
    case serverCommands.requestDetail:
      setRequestDetail(param.coll, param.req);
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

// @ts-ignore
// eslint-disable-next-line
function onSessionMessage(cmd: string, param: any): void {
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
    default:
      logWarn("unhandled session message [" + cmd + "]", param);
  }
}

export const messageHandler = (msg: Message): void => {
  if (isDebug()) {
    logDebug("IN: " + msg.cmd, jsonClone(msg.param));
  }

  switch (msg.svc) {
    case systemService.key:
      onSystemMessage(msg.cmd, msg.param);
      break;
    case collectionService.key:
      onCollectionMessage(msg.cmd, msg.param);
      break;
    case requestService.key:
      onRequestMessage(msg.cmd, msg.param);
      break;
    case sessionService.key:
      onSessionMessage(msg.cmd, msg.param);
      break;
    default:
      logWarn("unhandled service [" + msg.svc + "]", msg);
  }
};
