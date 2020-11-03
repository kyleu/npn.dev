<template>
  <div v-if="response" class="mt">
    <em>{{ response.method }} {{ response.url }}</em>
    <div>
      <ul data-uk-tab="">
        <li><a href="#result">Result</a></li>
        <li><a href="#request">Request</a></li>
        <li><a href="#headers">Headers</a></li>
        <li><a href="#body">Body</a></li>
        <li><a href="#timing">Timing</a></li>
      </ul>
      <ul class="uk-switcher uk-margin">
        <li><ResultSummary :response="response" /></li>
        <li><ResultHeaders title="Final Request Headers" :headers="response.requestHeaders" /></li>
        <li><ResultHeaders title="Response Headers" :headers="response.headers" /></li>
        <li><ResultBody :url="response.url" :body="response.body" /></li>
        <li><ResultTiming :timing="response.timing" /></li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {NPNResponse} from "@/call/model";
import ResultTiming from "@/call/ResultTiming.vue";
import ResultSummary from "@/call/ResultSummary.vue";
import ResultBody from "@/call/ResultBody.vue";
import ResultHeaders from "@/call/ResultHeaders.vue";

@Component({ components: {ResultTiming, ResultBody, ResultSummary, ResultHeaders } })
export default class ResponsePanel extends Vue {
  @Prop() response!: NPNResponse
}
</script>
