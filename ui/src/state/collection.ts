import {logWarn} from "@/util/log";
import {serverCommands} from "@/util/command";
import {
  CollectionAddedParams, CollectionData,
  collectionsRef,
  onCollectionAdded,
  onCollectionDeleted,
  onCollectionNotFound,
  onCollectionUpdated,
  setCollectionRequestSummaries
} from "@/collection/state";
import {setCollectionTransformResult} from "@/transform/state";
import {Collection, CollectionCount} from "@/collection/collection";
import {Summary} from "@/request/model";
import {CollectionTransformResult} from "@/transform/result";

export function onCollectionMessage(cmd: string, param: unknown): void {
  switch (cmd) {
    case serverCommands.collections:
      collectionsRef.value = param as CollectionCount[];
      break;
    case serverCommands.collectionAdded:
      onCollectionAdded(param as CollectionAddedParams);
      break;
    case serverCommands.collectionUpdated:
      onCollectionUpdated(param as Collection);
      break;
    case serverCommands.collectionDeleted:
      onCollectionDeleted(param as string);
      break;
    case serverCommands.collectionDetail:
      setCollectionRequestSummaries(param as CollectionData<Summary>);
      break;
    case serverCommands.collectionTransform:
      setCollectionTransformResult(param as CollectionTransformResult);
      break;
    case serverCommands.collectionNotFound:
      onCollectionNotFound(param as string);
      break;
    default:
      logWarn("unhandled collection message [" + cmd + "]", param);
  }
}
