<template>
  <div>
    <div class="nav-section">
      <router-link tag="div" class="nav-header" :to="{name: 'CollectionIndex'}">Collections</router-link>
      <CollectionList />
    </div>
    <div class="nav-section mt">
      <div class="nav-header">System</div>
      <div class="nav-list">
        <div class="nav-link">
          <a :class="profile.settings.linkColor + '-fg'" href="" onclick="npn.debug();return false">
            <span class="uk-icon nav-icon" data-uk-icon="icon: code"></span> Debug
          </a>
        </div>
        <div class="nav-link">
          <router-link :class="profile.settings.linkColor + '-fg'" to="/about">
            <span class="uk-icon nav-icon" data-uk-icon="icon: question"></span> About
          </router-link>
        </div>
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
