<template>
  <div v-if="cycle.rsp">
    <h3 class="uk-card-title">{{ cycle.rsp.status }}</h3>
    <em>{{ cycle.rsp.method }} {{ cycle.rsp.url }}</em>
    <div class="mt">
      <ul data-uk-tab="">
        <li><a id="tab-response-result" href="#result" @click="setTab('result')">Result</a></li>
        <li><a id="tab-response-request" href="#request" @click="setTab('request')">Request</a></li>
        <li><a id="tab-response-headers" href="#headers" @click="setTab('headers')">Headers</a></li>
        <li><a id="tab-response-body" href="#body" @click="setTab('body')">Body</a></li>
        <li><a id="tab-response-timing" href="#timing" @click="setTab('timing')">Timing</a></li>
      </ul>
      <ul class="uk-switcher">
        <li><SummaryResponse :response="cycle.rsp" /></li>
        <li><HeadersResponse title="Final Request Headers" :headers="cycle.rsp.requestHeaders" /></li>
        <li><HeadersResponse title="Response Headers" :headers="cycle.rsp.headers" /></li>
        <li><BodyResponse ref="body" :url="cycle.rsp.url" :body="cycle.rsp.body" /></li>
        <li><TimingResponse :timing="cycle.rsp.timing" /></li>
      </ul>
    </div>
  </div>
  <div v-else-if="cycle.error.length > 0">
    <h3 class="uk-card-title">Error</h3>
    <em>{{ cycle.method }} {{ cycle.url }}</em>
    <p>{{ cycle.error }}</p>
  </div>
  <div v-else>
    <h3 class="uk-card-title">Loading...</h3>
    <em>{{ cycle.method }} {{ cycle.url }}</em>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {RequestResultCycle} from "@/call/model";
import TimingResponse from "@/call/TimingResponse.vue";
import SummaryResponse from "@/call/SummaryResponse.vue";
import BodyResponse from "@/body/BodyResponse.vue";
import HeadersResponse from "@/header/HeadersResponse.vue";

@Component({ components: {BodyResponse, HeadersResponse, SummaryResponse, TimingResponse } })
export default class ResponsePanel extends Vue {
  activeTab = "";

  @Prop() cycle!: RequestResultCycle;

  setTab(s: string): void {
    this.activeTab = s;
    if (s === "body") {
      (this.$refs["body"] as BodyResponse).refresh();
    }
  }
}
</script>
