<template>
  <li>
    <a class="uk-accordion-title" href="#">{{ file.filename }} ({{ file.type }})</a>
    <div class="uk-accordion-content">
      <div v-if="file.error" class="uk-alert-danger" style="margin: 12px 0; padding: 6px;">{{ file.error }}</div>
      <pre>{{ file.value }}</pre>
    </div>
  </li>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import Icon from "@/util/Icon.vue";
import {ImportResultFile} from "@/import/model";

@Component({ components: { Icon } })
export default class ImportResultSection extends Vue {
  @Prop() file!: ImportResultFile;
}
</script>
