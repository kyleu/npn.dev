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
            <span class="uk-icon nav-icon" data-uk-icon="icon: code"></span> Debug
          </a>
        </div>
        <div class="nav-link">
          <router-link v-style-menu-link to="/about">
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

@Component({ components: { CollectionList } })
export default class LeftNav extends Vue {}
</script>
