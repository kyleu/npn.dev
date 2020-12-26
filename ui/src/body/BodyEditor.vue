<template>
  <div>
    <select v-model="body.type" class="uk-select">
      <option value="">No request body</option>
      <option v-for="t in types()" :key="t.key" :value="t.key">{{ t.title }}</option>;
    </select>

    <div v-if="body.type === ''" class="mt">
      No body!
    </div>
    <div v-else-if="body.type === 'form'" class="mt">
      <FormEditor />
    </div>
    <div v-else-if="body.type === 'html'" class="mt">
      <CodeEditor ref="html" v-model="body.htmlContent" language="htmlmixed" />
    </div>
    <div v-else-if="body.type === 'xml'" class="mt">
      <CodeEditor ref="xml" v-model="body.xmlContent" language="htmlmixed" />
    </div>
    <div v-else-if="body.type === 'json'" class="mt">
      <CodeEditor ref="json" v-model="body.jsonContent" language="javascript" />
    </div>
    <div v-else class="mt">
      Unhandled [{{ body.type }}] body editor
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {AllTypes, BodyType} from "@/body/model";
import {BodyConfig, bodyConfigRef} from "@/body/state";
import FormEditor from "@/body/FormEditor.vue";
import CodeEditor from "@/body/CodeEditor.vue";

@Component({ components: {CodeEditor, FormEditor } })
export default class BodyEditor extends Vue {
  get body(): BodyConfig | undefined {
    return bodyConfigRef.value;
  }

  refresh(): void {
    this.$nextTick(function() {
      if (this.$refs["html"]) {
        (this.$refs["html"] as CodeEditor).refresh();
      }
      if (this.$refs["xml"]) {
        (this.$refs["xml"] as CodeEditor).refresh();
      }
      if (this.$refs["json"]) {
        (this.$refs["json"] as CodeEditor).refresh();
      }
    });
  }

  types(): BodyType[] {
    return AllTypes.filter(t => !t.hidden);
  }
}
</script>
