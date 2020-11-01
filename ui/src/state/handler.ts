import {
  activeRequestRef,
  callResultRef,
  collectionsRef,
  requestEditingRef,
  requestOriginalRef,
  setCollectionRequestSummaries, setRequestDetail,
  transformResultRef
} from "@/state/state";
import {Message} from "@/socket/socket";
import {logDebug, logWarn} from "@/util/log";
import {Collection} from "@/collection/collection";
import {cloneRequest} from "@/request/model";

export const messageHandler = (_: object, msg: Message): void => {
  logDebug("IN", msg);
  switch (msg.cmd) {
    case "collections":
      collectionsRef.value = msg.param as Collection[];
      break;
    case "collectionDetail":
      setCollectionRequestSummaries(msg.param.key, msg.param.requests);
      break;
    case "requestDetail":
      setRequestDetail(msg.param.coll, msg.param.req);
      if (activeRequestRef.value && msg.param.req.key === activeRequestRef.value.req && msg.param.coll === activeRequestRef.value.coll) {
        requestOriginalRef.value = cloneRequest(msg.param.req);
        requestEditingRef.value = msg.param.req;
      }
      break;
    case "callResult":
      callResultRef.value = msg.param;
      break;
    case "transformResult":
      transformResultRef.value = msg.param;
      break;
    default:
      logWarn("unhandled message [" + msg.cmd + "]", msg);
  }
}
