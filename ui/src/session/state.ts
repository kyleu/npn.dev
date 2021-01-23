import {socketRef} from "@/socket/socket";
import {ref} from "@vue/composition-api";
import {sessionService} from "@/util/services";
import {clientCommands} from "@/util/command";
import {clearPendingRequests, pendingRequestsRef, setPendingRequests} from "@/socket/pending";
import {jsonClone} from "@/util/json";
import {SessAdded, Session, SessionSummary} from "@/session/model";
import {globalRouter} from "@/util/vutils";

export const activeSessionRef = ref<string>("_");
export const sessionSummariesRef = ref<SessionSummary[]>([]);
export const sessionDetailsRef = ref<Session[]>([]);

export const sessionOriginalRef = ref<Session>();
export const sessionEditingRef = ref<Session>();

export function getSessionDetail(key: string): Session | undefined {
  for (const s of sessionDetailsRef.value) {
    if (s.key === key) {
      return s;
    }
  }
}

export function setActiveSession(key: string): void {
  activeSessionRef.value = key;
  const s = getSessionDetail(key);
  if (s && s.key === key) {
    if(!sessionOriginalRef.value || (sessionOriginalRef.value.key !== key)) {
      sessionOriginalRef.value = s;
      sessionEditingRef.value = jsonClone(s);
    }
    return;
  }

  if (key.length > 0 && socketRef.value) {
    if (setPendingRequests(pendingRequestsRef, "session", key)) {
      socketRef.value.send({channel: sessionService.key, cmd: clientCommands.getSession, param: activeSessionRef.value});
    }
  }
}

export function setSessionDetail(s: Session): void {
  clearPendingRequests(pendingRequestsRef, "session", s.key);

  const rs = sessionDetailsRef.value || [];
  let matched = false;
  for (const r in rs) {
    if (rs[r].key === s.key) {
      matched = true;
      rs[r] = s;
    }
  }
  if (!matched) {
    rs.push(s);
  }

  if (activeSessionRef.value && s.key === activeSessionRef.value) {
    sessionOriginalRef.value = s;
    sessionEditingRef.value = jsonClone(s);
  }
}

export function onSessionAdded(msg: SessAdded): void {
  sessionSummariesRef.value = msg.sessions;
  setSessionDetail(msg.active);
  globalRouter().push({name: "SessionDetail", params: {sess: msg.active.key}});
}

export function onSessionDeleted(param: string): void {
  sessionDetailsRef.value = sessionDetailsRef.value.filter(x => x.key !== param);
  sessionSummariesRef.value = sessionSummariesRef.value.filter(x => x.key !== param);
  if (activeSessionRef.value === param) {
    activeSessionRef.value = "_";
  }
  globalRouter().push({name: "SessionIndex"});
}

export function onSessionNotFound(): void {
  globalRouter().push({name: "SessionIndex"});
}
