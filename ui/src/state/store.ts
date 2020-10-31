import Vue from "vue";
import Vuex, {Store} from "vuex";
import {initialState, State} from "@/state";
import {Message, Socket} from "@/socket/socket";
import {logDebug, logError, logWarn} from "@/util/log";
import {cloneRequest} from "@/request/model";

Vue.use(Vuex);

export interface ActiveRequest {
  readonly coll: string;
  readonly req: string;
}

export function newStore(onMsg: (s: State, m: Message) => void): Store<State> {
  const state = initialState();
  let socket: Socket | undefined = undefined;

  function send(s: State, msg: Message): void {
    if (!socket) {
      logError("no socket available");
    } else {
      socket.send(msg);
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

  const ret = new Vuex.Store({
    state: state,
    mutations: {
      onMessage: onMsg,
      send: send,
      setActiveRequest: setActiveRequest
    },
    modules: {}
  });

  function openF(): void {
    logDebug("websocket open");
  }
  function recvF(m: Message): void {
    ret.commit("onMessage", m)
  }
  function errF(err: string): void {
    logWarn("websocket err: " + err);
  }

  socket = new Socket(openF, recvF, errF, state.url);

  return ret;
}
