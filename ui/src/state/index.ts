import {Collection, MockCollections} from "@/collection/collection";
import Profile from "@/user/profile";

export interface Breadcrumb {
  readonly title: string;
  readonly path: string;
}

export interface State {
  readonly url: string;
  profile: Profile;
  breadcrumbs: Breadcrumb[];

  collections: Collection[];
  activeCollection: Collection | undefined;
}

interface InitialData {
  readonly url: string;
  readonly profile: Profile;
}

export function initialState(): State {
  // @ts-ignore
  // eslint-disable-next-line
  const cfg = (window as any).initialData as InitialData;

  let profile = {} as Profile;
  const url = cfg && cfg.url ? cfg.url : "";

  if (cfg && cfg.profile) {
    profile = cfg.profile;
  }

  return {
    url: url,
    profile: profile,
    breadcrumbs: [],

    collections: MockCollections,
    activeCollection: undefined
  };
}
