import {defaultProfile, Profile, profileRef} from "@/user/profile";
import {hostRef, Message, Socket, socketRef} from "@/socket/socket";
import {logDebug} from "@/util/log";

interface InitialData {
  host: string;
  profile: Profile;
}

declare global {
  interface Window {
    initialData: InitialData | undefined;
  }
}

export function initState(onMessage: (m: Message) => void): void {
  let cfg = window.initialData;
  if (!cfg) {
    cfg = { host: "", profile: defaultProfile };
  }
  if (!cfg.host) {
    cfg.host = '';
  }
  if (!cfg.profile) {
    cfg.profile = defaultProfile;
  }

  let profile: Profile | undefined;
  const host = cfg && cfg.host ? cfg.host : "";

  if (cfg && cfg.profile) {
    profile = cfg.profile;
  }

  hostRef.value = host;
  profileRef.value = profile;

  function openF(): void {
    logDebug("websocket open");
  }
  function recvF(m: Message): void {
    onMessage(m);
  }
  function errF(err: string): void {
    logDebug("websocket err: " + err);
  }
  let url = "";
  if(hostRef.value.length > 0) {
    url = `ws://${hostRef.value}/ws`;
  }
  socketRef.value = new Socket(openF, recvF, errF, url);
}

