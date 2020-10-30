<template>
  <div>
    <div class="nav-section">
      <router-link tag="div" class="nav-header" :to="{name: 'CollectionIndex'}">Collections</router-link>
      <CollectionList />
    </div>
    <div class="nav-section">
      <div class="nav-header">System</div>
      <div class="nav-list collection-list">
        <router-link :class="'collection-link ' + $store.state.profile.linkColor + '-fg'" to="/about">About</router-link>
      </div>
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
