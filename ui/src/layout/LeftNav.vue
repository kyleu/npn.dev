<template>
  <div>
    <SessionSelector />
    <div class="nav-section mt">
      <router-link v-style-menu-section class="nav-header" tag="div" :to="{name: 'CollectionIndex'}">Collections</router-link>
      <CollectionList />
    </div>
    <div class="nav-section mt">
      <div v-style-menu-section>System</div>
      <div class="nav-list">
        <div class="nav-link"><a v-style-menu-link href="" onclick="npn.debug();return false"><Icon icon="code" /> Debug</a></div>
        <div class="nav-link"><router-link v-style-menu-link to="/about"><Icon icon="question" /> About</router-link></div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import CollectionList from "@/collection/CollectionList.vue";
import Icon from "@/util/Icon.vue";
import SessionSelector from "@/session/SessionSelector.vue";

@Component({ components: { Icon, CollectionList, SessionSelector } })
export default class LeftNav extends Vue {}
</script>
