<template>
  <div class="uk-overflow-auto">
    <em>Error</em>
    <div>{{ config.message }}</div>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {ErrorConfig} from "@/body/model";

declare const Prism: object;

@Component
export default class ErrorBody extends Vue {
  @Prop() config!: ErrorConfig;
}
</script>
