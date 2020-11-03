<template>
  <div class="uk-overflow-auto">
    <div v-if="body">
      <em>{{ body.type }}</em>
      <div class="prism-view mt"><pre style="margin: 0;"><code v-html="highlighted"></code></pre></div>
    </div>
    <div v-else>
      <div>no body</div>
      <p><em>{{ url }}!</em></p>
    </div>
  </div>
</template>

<script lang="ts">
// @ts-ignore
// eslint-disable-next-line
declare const Prism: any;

import {Component, Prop, Vue} from "vue-property-decorator";
import {RBody} from "@/body/model";
import {jsonStr} from "@/util/json";
import {PrismEditor} from "vue-prism-editor";

@Component({ components: { PrismEditor } })
export default class BodyResponse extends Vue {
  @Prop() url!: string
  @Prop() body!: RBody

  get content(): string {
    if (!this.body) {
      return "no body";
    }
    if (this.body.type === "html") {
      return this.body.config.content;
    }
    if (this.body.type === "json") {
      return jsonStr(this.body.config.msg);
    }
    return "unknown body [" + this.body.type + "]";
  }

  get highlighted(): string {
    return Prism.highlight(this.content, this.getLang(this.body.type));
  }

  private getLang(t: string) {
    switch (t) {
      case "json":
        return Prism.languages.js;
      default:
        return Prism.languages[t];
    }
  }
}
</script>
