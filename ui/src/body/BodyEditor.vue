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
      <textarea v-model="body.htmlContent" rows="8" class="uk-textarea"></textarea>
    </div>
    <div v-else-if="body.type === 'json'" class="mt">
      <textarea v-model="body.jsonContent" rows="8" class="uk-textarea"></textarea>
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
@Component({
  components: {FormEditor}
})
export default class BodyEditor extends Vue {
  get body(): BodyConfig | undefined {
    return bodyConfigRef.value;
  }

  types(): BodyType[] {
    return AllTypes.filter(t => !t.hidden);
  }
}
</script>
