<template>
  <div class="uk-overflow-auto">
    <em>JSON</em>
    <pre class="code-view preview-content mt" style="margin: 0;"><code v-html="highlighted"></code></pre>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {JSONConfig} from "@/body/model";
import {jsonStr} from "@/util/json";

declare const hljs: {
  highlight: (l: string, c: string) => {value: string};
};

@Component
export default class JSONBody extends Vue {
  @Prop() config!: JSONConfig

  get highlighted(): string {
    return hljs.highlight("json", jsonStr(this.config.msg)).value;
  }
}
</script>
