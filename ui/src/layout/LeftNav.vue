<template>
  <div>
    <div class="nav-section">
      <router-link v-style-menu-section tag="div" class="nav-header" :to="{name: 'CollectionIndex'}">Collections</router-link>
      <CollectionList />
    </div>
    <div class="nav-section mt">
      <div v-style-menu-section class="nav-header">System</div>
      <div class="nav-list">
        <div class="nav-link">
          <a v-style-menu-link href="" onclick="npn.debug();return false">
            <Icon icon="code" /> Debug
          </a>
        </div>
        <div class="nav-link">
          <router-link v-style-menu-link to="/about">
            <Icon icon="question" /> About
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import CollectionList from "@/collection/CollectionList.vue";
import Icon from "@/util/Icon.vue";

@Component({ components: {Icon, CollectionList } })
export default class LeftNav extends Vue {}
</script>
