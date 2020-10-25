<template>
  <div id="npn">
    <NavBar />
    <div id="nav-left" class="uk-visible@m"><LeftNav /></div>
    <div id="workspace-content">
      <router-view />
    </div>
    <div id="nav-offcanvas" data-uk-offcanvas="mode: push; overlay: true">
      <div class="uk-offcanvas-bar"><LeftNav /></div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import NavBar from "@/layout/NavBar.vue";
import LeftNav from "@/layout/LeftNav.vue";

@Component({ components: { LeftNav, NavBar } })
export default class Workspace extends Vue {}
</script>

<style lang="scss"></style>
