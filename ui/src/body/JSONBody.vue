<template>
  <div class="uk-overflow-auto">
    <em>JSON</em>
    <pre class="prism-view preview-content mt" style="margin: 0;"><code v-html="highlighted"></code></pre>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {JSONConfig} from "@/body/model";
import {jsonStr} from "@/util/json";

// @ts-ignore
// eslint-disable-next-line
declare const Prism: any;

@Component
export default class JSONBody extends Vue {
  @Prop() config!: JSONConfig

  get highlighted(): string {
    return Prism.highlight(jsonStr(this.config.msg), Prism.languages.js);
  }
}
</script>
