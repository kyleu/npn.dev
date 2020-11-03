<template>
  <div>
    <div>{{ (response.timing.completed || 0) / 1000 }}ms</div>
    <div>
      {{ response.proto }}
      <em>{{ response.status }}</em>
      <div>
        {{ response.contentType || 'unknown' }}
        {{ (response.contentLength && response.contentLength > -1) ? '(' + response.contentLength + 'bytes)' : ((response.body && response.body.length > -1) ? '(' + response.body.length + ' bytes)' : "") }}
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {NPNResponse} from "@/call/model";

@Component
export default class ResponseSummary extends Vue {
  @Prop() response!: NPNResponse
}
</script>
