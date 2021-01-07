import {logWarn} from "@/util/log";
import {serverCommands} from "@/util/command";
import {
  collectionsRef,
  onCollectionAdded,
  onCollectionDeleted,
  onCollectionNotFound,
  onCollectionUpdated,
  setCollectionRequestSummaries
} from "@/collection/state";
import {setCollectionTransformResult} from "@/transform/state";

// @ts-ignore
// eslint-disable-next-line
export function onCollectionMessage(cmd: string, param: any): void {
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
