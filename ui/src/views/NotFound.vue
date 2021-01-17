<template>
  <div class="about">
    <h1>Not Found</h1>
    <p>Sorry about that</p>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import {setBC} from "@/util/vutils";

@Component
export default class NotFound extends Vue {
  mounted(): void {
    setBC(this, {path: "", title: "not found"});
  }
}
</script>
