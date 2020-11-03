<template>
  <div>
    <div class="nav-section">
      <router-link tag="div" class="nav-header" :to="{name: 'CollectionIndex'}">Collections</router-link>
      <CollectionList />
    </div>
    <div class="nav-section mt">
      <div class="nav-header">System</div>
      <div class="nav-list">
        <div><a :class="'collection-link ' + profile.linkColor + '-fg'" href="" onclick="npn.debug();return false">Debug</a></div>
        <div><router-link :class="'collection-link ' + profile.linkColor + '-fg'" to="/about">About</router-link></div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import CollectionList from "@/collection/CollectionList.vue";
import {Profile, profileRef} from "@/user/profile";

@Component({ components: { CollectionList } })
export default class LeftNav extends Vue {
  get profile(): Profile | undefined {
    return profileRef.value;
  }
}
</script>

<style lang="scss">
.left-nav {
  margin-top: 24px;
}
</style>
