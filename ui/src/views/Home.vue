<template>
  <div class="home">
    <img alt="logo" src="../assets/logo.png" />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import LeftNav from "@/components/LeftNav.vue";

@Component({
  components: {
    LeftNav
  }
})
export default class Home extends Vue {}
</script>
