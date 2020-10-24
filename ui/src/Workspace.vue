<template>
  <div id="app">
    <div class="uk-dark">
      <NavBar />
      <div id="nav-left" class="uk-visible@m"><LeftNav /></div>
      <div id="workspace-content">
        <router-view />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import NavBar from "@/components/NavBar.vue";
import LeftNav from "@/components/LeftNav.vue";

@Component({
  components: {LeftNav, NavBar}
})
export default class Workspace extends Vue {}
</script>

<style lang="scss"></style>
