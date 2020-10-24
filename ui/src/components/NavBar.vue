<template>
  <div class="navbar">
    <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/about">About</router-link>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component
export default class NavBar extends Vue {}
</script>

<style scoped lang="scss">
a {
  &.router-link-exact-active {
    color: #000;
  }
}
</style>
