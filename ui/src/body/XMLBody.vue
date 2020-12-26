<template>
  <div class="uk-overflow-auto">
    <em>XML</em>
    <div ref="content" class="mt code-editor"></div>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {XMLConfig} from "@/body/model";

// @ts-ignore
// eslint-disable-next-line
declare const CodeMirror: any;

@Component
export default class XMLBody extends Vue {
  @Prop() config!: XMLConfig

  // @ts-ignore
  // eslint-disable-next-line
  editor: any

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
      e.setValue(this.config.content);
      e.setSize('100%', '100%');
    }
  }

  mounted(): void {
    const el = this.$refs["content"] as HTMLElement;
    this.editor = CodeMirror(el, {
      lineNumbers: true,
      mode: "htmlmixed",
      value: this.config.content,
      readOnly: "nocursor"
    });
    this.editor.setSize('100%', '100%');
  }
}
</script>
