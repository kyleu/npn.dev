<template>
  <div v-if="response">
    <h3 class="uk-card-title">{{ response.status }}</h3>
    <em>{{ response.method }} {{ response.url }}</em>
    <div class="mt">
      <ul data-uk-tab="">
        <li><a href="#result" @click="setTab('result')">Result</a></li>
        <li><a href="#request" @click="setTab('request')">Request</a></li>
        <li><a href="#headers" @click="setTab('headers')">Headers</a></li>
        <li><a href="#body" @click="setTab('body')">Body</a></li>
        <li><a href="#timing" @click="setTab('timing')">Timing</a></li>
      </ul>
      <ul class="uk-switcher uk-margin">
        <li><SummaryResponse :response="response" /></li>
        <li><HeadersResponse title="Final Request Headers" :headers="response.requestHeaders" /></li>
        <li><HeadersResponse title="Response Headers" :headers="response.headers" /></li>
        <li><BodyResponse :url="response.url" :body="response.body" /></li>
        <li><TimingResponse :timing="response.timing" /></li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {NPNResponse} from "@/call/model";
import TimingResponse from "@/call/TimingResponse.vue";
import SummaryResponse from "@/call/SummaryResponse.vue";
import BodyResponse from "@/body/BodyResponse.vue";
import HeadersResponse from "@/header/HeadersResponse.vue";

@Component({ components: {BodyResponse, HeadersResponse, SummaryResponse, TimingResponse } })
export default class ResponsePanel extends Vue {
  activeTab = "";

  @Prop() response!: NPNResponse;

  setTab(s: string): void {
    this.activeTab = s;
  }
}
</script>
