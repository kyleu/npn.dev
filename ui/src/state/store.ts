import Vue from "vue";
import Vuex, {Store} from "vuex";
import {activeRequestRef, getRequestDetail, hostRef, requestEditingRef, requestOriginalRef, socketRef} from "@/state/state";
import {Message, Socket} from "@/socket/socket";
import {logDebug, logError, logWarn} from "@/util/log";
import {cloneRequest} from "@/request/model";
import {initState} from "@/state/initial";

Vue.use(Vuex);

export interface NPNState {}

export interface ActiveRequest {
  readonly coll: string;
  readonly req: string;
}

function send(s: NPNState, msg: Message): void {
  if (!socketRef.value) {
    logError("no socket available");
  } else {
    socketRef.value.send(msg);
  }
}

function setActiveRequest(s: NPNState, x: ActiveRequest): void {
  activeRequestRef.value = x;
  const rd = getRequestDetail(x.coll, x.req);
  requestEditingRef.value = rd;
  if (rd && ((!requestOriginalRef.value) || requestOriginalRef.value.key !== x.req)) {
    requestOriginalRef.value = cloneRequest(requestEditingRef.value);
  }
}

export function newStore(onMessage: (s: NPNState, m: Message) => void): Store<NPNState> {
  let ret: Store<NPNState> | undefined = undefined;

  initState();
  function openF(): void {
    logDebug("websocket open");
  }
  function recvF(m: Message): void {
    if (ret) {
      ret.commit("onMessage", m);
    }
  }
  function errF(err: string): void {
    logWarn("websocket err: " + err);
  }
  let url = "";
  if(hostRef.value.length > 0) {
    url = `ws://${hostRef.value}/s`
  }
  socketRef.value = new Socket(openF, recvF, errF, url);

  ret = new Vuex.Store({
    mutations: { onMessage, send, setActiveRequest },
    modules: {}
  })

  return ret;
}
