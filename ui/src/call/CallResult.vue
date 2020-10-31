<template>
  <div v-if="req" class="uk-card uk-card-body uk-card-default mt">
    <h3 v-if="result" class="uk-card-title">{{ result.status }}</h3>
    <em v-if="result">{{ result.response.method }} {{ result.response.url }}</em>
    <h3 v-else class="uk-card-title">Loading...</h3>
    <div v-if="result" class="mt">
      <ul data-uk-tab="">
        <li><a href="#result">Result</a></li>
        <li><a href="#request">Request</a></li>
        <li><a href="#headers">Headers</a></li>
        <li><a href="#body">Body</a></li>
        <li><a href="#timing">Timing</a></li>
      </ul>
      <ul class="uk-switcher uk-margin">
        <li>
          <div>{{ (result.response.timing.completed || 0) / 1000 }}ms</div>
          <div>
            {{ result.response.proto }}
            <em>{{ result.response.status }}</em>
            <div>
              {{ result.response.contentType || 'unknown' }}
              {{ (result.response.contentLength && result.response.contentLength > -1) ? '(' + result.response.contentLength + 'bytes)' : ((result.response.body && result.response.body.length > -1) ? '(' + result.response.body.length + ' bytes)' : "") }}
            </div>
          </div>
        </li>
        <li><ResultHeaders title="Final Request Headers" :headers="result.response.requestHeaders" /></li>
        <li><ResultHeaders title="Response Headers" :headers="result.response.headers" /></li>
        <li><ResultBody :url="result.response.url" :body="result.response.body" /></li>
        <li>{renderTiming(rsp.timing)}</li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {NPNRequest} from "@/request/model";
import {getState, getStateSetBCReq} from "@/util/vutils";
import {Result} from "@/call/model";
import ResultHeaders from "@/call/ResultHeaders.vue";
import ResultBody from "@/call/ResultBody.vue";
import RequestDetail from "@/request/RequestDetail.vue";

@Component({ components: { ResultBody, ResultHeaders } })
export default class CallResult extends Vue {
  get req(): NPNRequest | undefined {
    return (this.$parent as RequestDetail).req;
  }

  get result(): Result | undefined {
    return getState(this).callResult;
  }

  created(): void {
    if (this.req) {
      const param = {coll: this.$route.params.coll, req: this.$route.params.req, proto: this.req?.prototype};
      this.$store.commit("send", {svc: "request", cmd: "call", param: param});
    }
    getStateSetBCReq(this, "call");
  }

  updated(): void {
    if (this.req) {
      const param = {coll: this.$route.params.coll, req: this.$route.params.req, proto: this.req?.prototype};
      this.$store.commit("send", {svc: "request", cmd: "call", param: param});
    }
    getStateSetBCReq(this, "call");
  }
}
</script>
