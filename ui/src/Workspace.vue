<template>
  <div id="app">
    <NavBar />
    <LeftNav />
    <router-view />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import NavBar from "@/components/NavBar.vue";
import LeftNav from "@/components/LeftNav.vue";

@Component({
  components: {
    LeftNav,
    NavBar
  }
})
export default class Workspace extends Vue {}
</script>

<style lang="scss"></style>
