<template>
  <div class="about">
    <h1>Not Found</h1>
    <p>Sorry about that</p>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import CollectionGallery from "@/collection/CollectionGallery.vue";
import {setBC} from "@/util/vutils";

@Component({ components: { CollectionGallery } })
export default class NotFound extends Vue {
  updated(): void {
    setBC(this, {path: "", title: "not found"});
  }
}
</script>
