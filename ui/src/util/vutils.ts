import {Vue} from "vue-property-decorator";
import {Breadcrumb, State} from "@/state";

export function getState(me: Vue): State {
  return (me.$store.state as State);
}

export function getStateSetBC(me: Vue, ...bc: Breadcrumb[]): State {
  const ret = getState(me);
  ret.breadcrumbs = bc;
  return ret;
}
