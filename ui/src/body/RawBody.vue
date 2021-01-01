<template>
  <div class="uk-overflow-auto">
    <em>{{ config.type || 'Raw Result' }}</em>
    <pre class="raw-content">{{ config.content }}</pre>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {RawConfig} from "@/body/model";

@Component
export default class RawBody extends Vue {
  @Prop() config!: RawConfig;
}
</script>
