<template>
  <div class="uk-overflow-auto">
    <div v-if="body">
      <div v-if="body.type === 'html'">
        <pre><code class="language-html">{{ content }}</code></pre>
      </div>
      <div v-else-if="body.type === 'json'">
        <pre><code class="language-json">{{ content }}</code></pre>
      </div>
      <div v-else>
        Unknown [{{ body.type }}]
      </div>
    </div>
    <div v-else>
      <div>no body</div>
      <p><em>{{ url }}!</em></p>
    </div>
  </div>
</template>

<script lang="ts">
// TODO import * as prism from "prismjs";
import {Component, Prop, Vue} from "vue-property-decorator";
import {RBody} from "@/body/model";
import {jsonStr} from "@/util/json";

@Component
export default class ResultBody extends Vue {
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
}
</script>
