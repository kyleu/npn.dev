<template>
  <div v-if="response">
    <h3 class="uk-card-title">{{ response.status }}</h3>
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
        <li><ResponseSummary :response="response" /></li>
        <li><ResponseHeaders title="Final Request Headers" :headers="response.requestHeaders" /></li>
        <li><ResponseHeaders title="Response Headers" :headers="response.headers" /></li>
        <li><ResponseBody :url="response.url" :body="response.body" /></li>
        <li><ResponseTiming :timing="response.timing" /></li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {NPNResponse} from "@/call/model";
import ResponseTiming from "@/call/ResponseTiming.vue";
import ResponseSummary from "@/call/ResponseSummary.vue";
import ResponseBody from "@/call/ResponseBody.vue";
import ResponseHeaders from "@/call/ResponseHeaders.vue";

@Component({ components: {ResponseTiming, ResponseBody, ResponseSummary, ResponseHeaders } })
export default class ResponsePanel extends Vue {
  @Prop() response!: NPNResponse
}
</script>
