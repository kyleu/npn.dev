<template>
  <div class="uk-overflow-auto">
    <div v-if="mode === 'text'" class="right">
      <a href="" title="render preview" @click.prevent="mode = 'preview'">preview</a>
    </div>
    <div v-else class="right">
      <a href="" title="render preview" @click.prevent="mode = 'text'">text</a>
    </div>
    <em>HTML</em>

    <div v-if="mode === 'text'">
      <pre class="code-view preview-content mt" style="margin: 0;"><code v-html="highlighted"></code></pre>
    </div>
    <div v-else>
      <div class="mt"><HTMLPreview :url="url" :html="config.content" /></div>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {HTMLConfig} from "@/body/model";
import HTMLPreview from "@/body/HTMLPreview.vue";

declare const hljs: {
  highlight: (l: string, c: string) => {value: string};
};

@Component({ components: { HTMLPreview } })
export default class HTMLBody extends Vue {
  @Prop() url!: string;
  @Prop() config!: HTMLConfig;

  mode = "text";

  get highlighted(): string {
    return hljs.highlight("html", this.config.content).value;
  }
}
</script>
