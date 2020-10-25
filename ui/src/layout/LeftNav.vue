<template>
  <div>
    <div class="nav-section">
      <router-link tag="div" class="nav-header" :to="{name: 'CollectionIndex'}">Collections</router-link>
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
