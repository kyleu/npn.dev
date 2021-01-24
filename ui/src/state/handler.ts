import {Message} from "@/socket/socket";
import {isDebug, logDebug, logWarn} from "@/util/log";
import {jsonClone} from "@/util/json";
import {collectionService, importService, requestService, sessionService, systemService} from "@/util/services";
import {onSessionMessage} from "@/state/session";
import {onRequestMessage} from "@/state/request";
import {onCollectionMessage} from "@/state/collection";
import {onSystemMessage} from "@/state/system";
import {onImportMessage} from "@/state/import";

export const messageHandler = (msg: Message): void => {
  if (isDebug()) {
    logDebug(`IN(${msg.channel}): ${msg.cmd}`, jsonClone(msg.param));
  }

  switch (msg.channel) {
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
    case importService.key:
      onImportMessage(msg.cmd, msg.param);
      break;
    default:
      logWarn("unhandled channel [" + msg.channel + "]", msg);
  }
};
