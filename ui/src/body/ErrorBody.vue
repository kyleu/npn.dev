<template>
  <div class="uk-overflow-auto">
    <em>Error</em>
    <div>{{ config.message }}</div>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {ErrorConfig} from "@/body/model";

// @ts-ignore
// eslint-disable-next-line
declare const Prism: any;

@Component
export default class ErrorBody extends Vue {
  @Prop() config!: ErrorConfig
}
</script>
