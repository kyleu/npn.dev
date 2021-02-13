<template>
  <div class="uk-overflow-auto">
    <div v-if="mode === 'text'" class="right">
      <a href="" title="render preview" @click.prevent="mode = 'preview'">preview</a>
    </div>
    <div v-else class="right">
      <a href="" title="render preview" @click.prevent="mode = 'text'">text</a>
    </div>
    <em>HTML</em>

    <div ref="content" class="mt code-editor" :style="{display: mode === 'text' ? 'block' : 'none'}"></div>
    <div v-if="mode === 'preview'">
      <HTMLPreview :url="url" :html="config.content" />
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {HTMLConfig} from "@/body/model";
import HTMLPreview from "@/body/HTMLPreview.vue";
import {Editor, editorFor} from "@/util/editor";

@Component({ components: { HTMLPreview } })
export default class HTMLBody extends Vue {
  @Prop() url!: string;
  @Prop() config!: HTMLConfig;

  mode = "text";

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
      e.setValue(this.config.content);
      e.setSize('100%', '100%');
    }
  }

  mounted(): void {
    const el = this.$refs["content"] as HTMLElement;
    this.editor = editorFor(el, true, "htmlmixed", this.config.content, true);
    this.editor.setSize('100%', '100%');
  }
}
</script>
