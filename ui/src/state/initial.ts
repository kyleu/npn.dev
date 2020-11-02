import {} from "@/util/vutils"

import {Profile, profileRef} from "@/user/profile";
import {hostRef, Message, Socket, socketRef} from "@/socket/socket";
import {logDebug, logWarn} from "@/util/log";

interface InitialData {
  readonly host: string;
  readonly profile: Profile;
}

export function initState(onMessage: (m: Message) => void): void {
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

  function openF(): void {
    logDebug("websocket open");
  }
  function recvF(m: Message): void {
    onMessage(m);
  }
  function errF(err: string): void {
    logWarn("websocket err: " + err);
  }
  let url = "";
  if(hostRef.value.length > 0) {
    url = `ws://${hostRef.value}/s`
  }
  socketRef.value = new Socket(openF, recvF, errF, url);
}

