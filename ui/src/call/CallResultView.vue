<template>
  <div class="uk-card uk-card-body uk-card-default mt">
    <div class="right"><router-link :class="'uk-icon ' + profile.linkColor + '-fg'" data-uk-icon="close" :to="'/c/' + this.$route.params.coll + '/' + this.$route.params.req"></router-link></div>
    <h3 class="uk-card-title">{{ result ? result.status : 'Loading...' }}</h3>
    <ResponsePanel v-for="(r, idx) in responses" :key="idx" :response="r" />
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {setBCReq} from "@/util/vutils";
import {CallResult, NPNResponse} from "@/call/model";
import {getCallResult} from "@/request/state";
import {Prototype} from "@/request/model";
import {Profile, profileRef} from "@/user/profile";
import ResponsePanel from "@/call/ResponsePanel.vue";

interface CallParam {
  coll: string;
  req: string;
  proto: Prototype;
}

@Component({ components: { ResponsePanel } })
export default class CallResultView extends Vue {
  private pending: CallParam | undefined;

  get profile(): Profile | undefined {
    return profileRef.value;
  }

  get result(): CallResult | undefined {
    return getCallResult(this.$route.params.coll, this.$route.params.req);
  }

  get responses(): NPNResponse[] {
    const ret: NPNResponse[] = [];
    let r = this.result?.response;
    while(r) {
      ret.push(r);
      r = r.prior
    }
    return ret.reverse();
  }

  created(): void {
    setBCReq(this, "call");
  }

  updated(): void {
    setBCReq(this, "call");
  }
}
</script>
