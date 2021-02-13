<template>
  <div class="uk-overflow-auto">
    <em>JSON</em>
    <div ref="content" class="mt code-editor"></div>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {JSONConfig} from "@/body/model";
import {jsonStr} from "@/util/json";
import {Editor, editorFor} from "@/util/editor";

@Component
export default class JSONBody extends Vue {
  @Prop() config!: JSONConfig

  editor?: Editor

  refresh(): void {
    const e = this.editor;
    if (e) {
      e.setSize('100%', '100%');
      setTimeout(function() { e.refresh(); }, 10);
    }
  }

  updated(): void {
    const e = this.editor;
    if(e) {
      e.setValue(jsonStr(this.config.msg));
      e.setSize('100%', '100%');
    }
  }

  mounted(): void {
    const el = this.$refs["content"] as HTMLElement;
    this.editor = editorFor(el, true, "javascript", jsonStr(this.config.msg), true);
    this.editor.setSize('100%', '100%');
  }
}
</script>
