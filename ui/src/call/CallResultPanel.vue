<template>
  <div class="uk-card uk-card-body uk-card-default mt">
    <div class="right"><router-link :class="'uk-icon ' + profile.linkColor + '-fg'" data-uk-icon="close" :to="'/c/' + this.$route.params.coll + '/' + this.$route.params.req"></router-link></div>
    <h3 v-if="result" class="uk-card-title">{{ result.status }}</h3>
    <h3 v-else class="uk-card-title">Loading...</h3>
    <em v-if="result">{{ result.response.method }} {{ result.response.url }}</em>
    <div v-if="result" class="mt">
      <ul data-uk-tab="">
        <li><a href="#result">Result</a></li>
        <li><a href="#request">Request</a></li>
        <li><a href="#headers">Headers</a></li>
        <li><a href="#body">Body</a></li>
        <li><a href="#timing">Timing</a></li>
      </ul>
      <ul class="uk-switcher uk-margin">
        <li><ResultSummary :result="result" /></li>
        <li><ResultHeaders title="Final Request Headers" :headers="result.response.requestHeaders" /></li>
        <li><ResultHeaders title="Response Headers" :headers="result.response.headers" /></li>
        <li><ResultBody :url="result.response.url" :body="result.response.body" /></li>
        <li><ResultTiming :timing="result.response.timing" /></li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {setBCReq} from "@/util/vutils";
import ResultBody from "@/call/ResultBody.vue";
import ResultHeaders from "@/call/ResultHeaders.vue";
import ResultSummary from "@/call/ResultSummary.vue";
import ResultTiming from "@/call/ResultTiming.vue";
import {CallResult} from "@/call/model";
import {getCallResult} from "@/request/state";
import {Prototype} from "@/request/model";
import {Profile, profileRef} from "@/user/profile";

interface CallParam {
  coll: string;
  req: string;
  proto: Prototype;
}

@Component({ components: {ResultTiming, ResultBody, ResultSummary, ResultHeaders } })
export default class CallResultPanel extends Vue {
  private pending: CallParam | undefined;

  get profile(): Profile | undefined {
    return profileRef.value;
  }

  get result(): CallResult | undefined {
    return getCallResult(this.$route.params.coll, this.$route.params.req);
  }

  created(): void {
    setBCReq(this, "call");
  }

  updated(): void {
    setBCReq(this, "call");
  }
}
</script>
