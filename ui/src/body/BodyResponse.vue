<template>
  <div class="uk-overflow-auto">
    <div v-if="body">
      <em>{{ body.type }}</em>
      <pre><code :class="'language-' + body.type">{{ content }}</code></pre>
    </div>
    <div v-else>
      <div>no body</div>
      <p><em>{{ url }}!</em></p>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {RBody} from "@/body/model";
import {jsonStr} from "@/util/json";

@Component
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
}
</script>
