<template>
  <div v-if="timing" class="timing-panel">
    <div class="result-timing-graph">
      <div>
        <div class="timing-start">0ms</div>
        <div class="timing-end">{{ timing.completed / 1000 }}ms</div>
      </div>
      <object type="image/svg+xml" :style="'width: 100%; height: ' + (sections.length * 24) + 'px'" :data="graphMarkup">SVG not supported</object>
    </div>
  </div>
  <div v-else>no timing</div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {Timing, timingGraph, TimingSection, timingSections} from "@/call/timing";
import {hostRef} from "@/socket/socket";

@Component
export default class TimingResponse extends Vue {
  @Prop() timing: Timing | undefined;

  get sections(): TimingSection[] {
    return this.timing ? timingSections(this.timing) : [];
  }

  get graphMarkup(): string {
    let url = "";
    if (hostRef.value.length > 0) {
      url = `http://${hostRef.value}`;
    }
    return timingGraph(url, this.sections);
  }
}
</script>
