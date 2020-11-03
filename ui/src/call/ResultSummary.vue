<template>
  <div>
    <div>{{ (result.response.timing.completed || 0) / 1000 }}ms</div>
    <div>
      {{ result.response.proto }}
      <em>{{ result.response.status }}</em>
      <div>
        {{ result.response.contentType || 'unknown' }}
        {{ (result.response.contentLength && result.response.contentLength > -1) ? '(' + result.response.contentLength + 'bytes)' : ((result.response.body && result.response.body.length > -1) ? '(' + result.response.body.length + ' bytes)' : "") }}
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {CallResult} from "@/call/model";

@Component
export default class ResultSummary extends Vue {
  @Prop() result!: CallResult
}
</script>
