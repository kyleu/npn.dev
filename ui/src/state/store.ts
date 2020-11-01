import Vue from "vue";
import Vuex, {Store} from "vuex";
import {State} from "@/state/state";
import {Message, Socket} from "@/socket/socket";
import {logDebug, logError, logWarn} from "@/util/log";
import {cloneRequest} from "@/request/model";
import {initialState} from "@/state/initial";

Vue.use(Vuex);

export interface ActiveRequest {
  readonly coll: string;
  readonly req: string;
}

function send(s: State, msg: Message): void {
  if (!s.socket) {
    logError("no socket available");
  } else {
    s.socket.send(msg);
  }
}

function setActiveRequest(s: State, x: ActiveRequest): void {
  s.activeRequest = x;
  const rd = s.getRequestDetail(x.coll, x.req);
  s.requestEditing = rd;
  if (rd && ((!s.requestOriginal) || s.requestOriginal.key !== x.req)) {
    console.log("1")
    s.requestOriginal = cloneRequest(s.requestEditing);
  }
}

export function newStore(onMessage: (s: State, m: Message) => void): Store<State> {
  let ret: Store<State> | undefined = undefined;

  const state = initialState();
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
  if(state.host && state.host.length > 0) {
    url = `ws://${state.host}/s`
  }
  state.socket = new Socket(openF, recvF, errF, url);

  ret = new Vuex.Store({
    state: state,
    mutations: { onMessage, send, setActiveRequest },
    modules: {}
  })

  return ret;
}
