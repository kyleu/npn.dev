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

export function getStateSetBCReq(me: Vue, action: string): State {
  const ret = getState(me);
  ret.breadcrumbs = [
    {path: "/c/" + me.$route.params.coll, title: me.$route.params.coll},
    {path: "/c/" + me.$route.params.coll + "/" + me.$route.params.req, title: me.$route.params.req},
    {path: "", title: action}
  ];
  return ret;
}
