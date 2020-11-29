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
      <div ref="content" class="mt"></div>
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

// @ts-ignore
// eslint-disable-next-line
declare const CodeMirror: any;

@Component({ components: { HTMLPreview } })
export default class HTMLBody extends Vue {
  @Prop() url!: string;
  @Prop() config!: HTMLConfig;

  mode = "text";

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
