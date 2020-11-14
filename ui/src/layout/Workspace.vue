<template>
  <div id="npn">
    <NavBar />
    <div id="nav-left" v-style-menu class="uk-visible@m"><LeftNav /></div>
    <div id="workspace-content">
      <router-view />
    </div>
    <div id="log-container" class="hidden">
      <div id="log-content">
        <ul id="log-list" class="uk-list uk-list-divider"></ul>
      </div>
    </div>
    <div id="nav-offcanvas" data-uk-offcanvas="mode: push; overlay: true">
      <div v-style-menu class="uk-offcanvas-bar"><LeftNav /></div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import NavBar from "@/layout/NavBar.vue";
import LeftNav from "@/layout/LeftNav.vue";

@Component({ components: {LeftNav, NavBar } })
export default class Workspace extends Vue {}
</script>

<style lang="scss"></style>
