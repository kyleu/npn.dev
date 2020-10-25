import {Collection, MockCollections} from "@/collection/collection";
import Profile from "@/user/profile";

export interface Breadcrumb {
  readonly title: string;
  readonly path: string;
}

export interface State {
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
  /* eslint-disable @typescript-eslint/no-explicit-any */
  // @ts-ignore
  const cfg = (window as any).initialData as InitialData;
  /* eslint-enable @typescript-eslint/no-explicit-any */

  let profile = {} as Profile;
  if (cfg) {
    profile = cfg.profile;
  }
  return {
    profile: profile,
    breadcrumbs: [
      {
        title: "asdf",
        path: "/asdf"
      }, {
        title: "qwer",
        path: "/asdf/qwer"
      }
    ],

    collections: MockCollections,
    activeCollection: undefined
  };
}
