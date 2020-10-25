import {Collection, MockCollections} from "@/collection/collection";
import Profile from "@/user/profile";

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
