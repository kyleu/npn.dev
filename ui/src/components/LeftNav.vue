<template>
  <div class="left-nav">
    <h1>LeftNav</h1>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component
export default class LeftNav extends Vue {}
</script>

<style scoped lang="scss">
.left-nav {
  margin-top: 24px;
}
</style>
