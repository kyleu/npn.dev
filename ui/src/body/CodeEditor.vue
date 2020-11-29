<template>
  <div class="code-editor">
    <input :value="value" name="code" type="hidden" data-lpignore="true" @input="$emit('input', $event.target.value)" />
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import FormEditor from "@/body/FormEditor.vue";

// @ts-ignore
// eslint-disable-next-line
declare const CodeMirror: any;

@Component({ components: { FormEditor } })
export default class CodeEditor extends Vue {
  @Prop() language!: string
  @Prop() value!: string

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
      const pos = e.getCursor();
      e.setValue(this.value);
      e.setCursor(pos);
      e.setSize('100%', '100%');
    }
  }

  mounted(): void {
    this.editor = CodeMirror(this.$el as HTMLElement, {
      lineNumbers: true,
      mode: this.language,
      value: this.value
    });
    this.editor.setSize('100%', '100%');
    this.editor.on("change", () => {
      const v = this.editor?.getValue();
      if (v !== this.value) {
        this.$emit('input', v);
      }
    });
  }
}
</script>
