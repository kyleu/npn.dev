<template>
  <div>
    <select v-model="body.type" class="uk-select">
      <option value="">No request body</option>
      <option v-for="t in types()" :key="t.key" :value="t.key">{{ t.title }}</option>;
    </select>

    <div v-if="(!body) || body.type === ''" class="mt">
      No body!
    </div>
    <div v-else-if="body.type === 'html'" class="mt">
      <textarea v-model="body.config.content" class="uk-textarea"></textarea>
    </div>
    <div v-else-if="body.type === 'json'" class="mt">
      <textarea :value="JSON.stringify(body.config.msg, null, 2)" class="uk-textarea" @input="body.config.msg = JSON.parse($event.target.value)"></textarea>
    </div>
    <div v-else class="mt">
      Unhandled {{ body.type }} editor
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {AllTypes, BodyType, RBody} from "@/body/model";

@Component
export default class BodyEditor extends Vue {
  @Prop() body: RBody | undefined;

  types(): BodyType[] {
    return AllTypes.filter(t => !t.hidden)
  }
}
</script>
