<template>
  <div class="uk-card uk-card-body uk-card-default mt">
    <div class="right"><router-link :class="'uk-icon ' + profile.linkColor + '-fg'" data-uk-icon="close" :to="'/c/' + this.$route.params.coll + '/' + this.$route.params.req"></router-link></div>
    <h3 v-if="result" class="uk-card-title">{{ result.req }}: {{ result.fmt }}</h3>
    <h3 v-else class="uk-card-title">Loading...</h3>
    <div v-if="result" class="mt">
      <pre>{{ result.out }}</pre>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {setBCReq} from "@/util/vutils";
import {TransformResult} from "@/request/transformResult";
import {profileRef, requestEditingRef, transformResultRef} from "@/state/state";
import {Prototype} from "@/request/model";
import Profile from "@/user/profile";

interface TransformParam {
  coll: string;
  req: string;
  fmt: string;
  proto: Prototype;
}

@Component
export default class RequestTransform extends Vue {
  private pending: TransformParam | undefined;

  get profile(): Profile | undefined {
    return profileRef.value;
  }

  get result(): TransformResult | undefined {
    return transformResultRef.value;
  }

  created(): void {
    transformResultRef.value = undefined;
    const re = requestEditingRef.value;
    if (re) {
      const param = {coll: this.$route.params.coll, req: this.$route.params.req, fmt: this.$route.params.tx, proto: re.prototype};
      this.$store.commit("send", {svc: "request", cmd: "transform", param: param});
    }
    setBCReq(this, this.$route.params.tx);
  }

  updated(): void {
    const re = requestEditingRef.value;
    if (re) {
      const param: TransformParam = {coll: this.$route.params.coll, req: this.$route.params.req, fmt: this.$route.params.tx, proto: re.prototype};
      if ((this.pending) && (this.pending.coll === param.coll && this.pending.req === param.req && this.pending.fmt === param.fmt)) {
        // console.log("?");
      } else {
        transformResultRef.value = undefined;
        this.$store.commit("send", {svc: "request", cmd: "transform", param: param});
        this.pending = param;
      }
    }
    setBCReq(this, this.$route.params.tx);
  }
}
</script>
