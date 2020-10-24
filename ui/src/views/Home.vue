<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <h3>Welcome to npn!</h3>
      <img class="mt" alt="logo" src="../assets/logo.png" />
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import LeftNav from "@/components/LeftNav.vue";

@Component({ components: { LeftNav } })
export default class Home extends Vue {}
</script>
