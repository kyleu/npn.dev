<template>
  <div>
    <div class="right">{{ ((response.timing.completed || 0) / 1000).toLocaleString() }}ms</div>
    <div>{{ ((response.contentLength && response.contentLength > -1) ? response.contentLength : ((response.body && response.body.length > -1) ? response.body.length : 0)).toLocaleString() }} bytes</div>
    <div class="right">{{ response.contentType || 'unknown' }}</div>
    <div v-if="response.error">
      ERROR: {{ response.error }}
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {NPNResponse} from "@/call/model";

@Component
export default class ResponseSummary extends Vue {
  @Prop() response!: NPNResponse;
}
</script>
