import {Collection, MockCollections} from "@/model/collection";
import Profile from "@/model/profile";

export interface State {
  profile: Profile;

  collections: Collection[];
  activeCollection: Collection | undefined;
}

export function initialState(): State {
  return {
    collections: MockCollections,
    activeCollection: undefined,
    profile: {} as Profile
  }
}
