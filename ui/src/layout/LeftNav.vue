<template>
  <div>
    <div class="nav-section">
      <div class="nav-header">Collections</div>
      <CollectionList />
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import CollectionList from "@/collection/CollectionList.vue";

@Component({ components: { CollectionList } })
export default class LeftNav extends Vue {}
</script>

<style lang="scss">
.left-nav {
  margin-top: 24px;
}
</style>
