<template>
  <ul class="uk-list uk-list-divider">
    <li v-if="body">Body: {{ body.type }}</li>
  </ul>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {RBody} from "@/body/model";

@Component
export default class RequestEditorBody extends Vue {
  @Prop() body!: RBody | undefined;
}
</script>
