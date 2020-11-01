import Profile from "@/user/profile";
import {hostRef, profileRef} from "@/state/state";

interface InitialData {
  readonly host: string;
  readonly profile: Profile;
}

export function initState(): void {
  // @ts-ignore
  // eslint-disable-next-line
  const cfg = (window as any).initialData as InitialData;

  let profile = {} as Profile;
  const host = cfg && cfg.host ? cfg.host : "";

  if (cfg && cfg.profile) {
    profile = cfg.profile;
  }

  hostRef.value = host;
  profileRef.value = profile;
}
