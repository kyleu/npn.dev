import {Collection} from "@/model/collection";

export interface State {
  collections: Collection[];
  activeCollection: Collection | undefined;
}
