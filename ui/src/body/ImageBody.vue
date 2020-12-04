<template>
  <div class="uk-overflow-auto">
    <em>Image ({{ config.type }})</em>
    <div class="image-body">
      <img alt="response image" :src="dataURL" />
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {ImageConfig} from "@/body/model";

@Component
export default class ImageBody extends Vue {
  @Prop() config!: ImageConfig;

  get dataURL(): string {
    return `data:${this.config.type};base64,${this.config.content}`;
  }
}
</script>
