import Profile from "@/user/profile";
import {State} from "@/state/state";

interface InitialData {
  readonly host: string;
  readonly profile: Profile;
}

export function initialState(): State {
  // @ts-ignore
  // eslint-disable-next-line
  const cfg = (window as any).initialData as InitialData;

  let profile = {} as Profile;
  const host = cfg && cfg.host ? cfg.host : "";

  if (cfg && cfg.profile) {
    profile = cfg.profile;
  }

  return new State(host, profile);
}
