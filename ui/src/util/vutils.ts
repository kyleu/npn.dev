import {Vue} from "vue-property-decorator";
import {Breadcrumb, breadcrumbsRef} from "@/state/state";

export function setBC(me: Vue, ...bc: Breadcrumb[]): void {
  breadcrumbsRef.value = bc;
}

export function setBCReq(me: Vue, action: string): void {
  breadcrumbsRef.value = [
    {path: "/c/" + me.$route.params.coll, title: me.$route.params.coll},
    {path: "/c/" + me.$route.params.coll + "/" + me.$route.params.req, title: me.$route.params.req},
    {path: "", title: action}
  ];
}
