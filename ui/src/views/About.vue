<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <h3 class="uk-card-title">About</h3>
        <p>TODO</p>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import CollectionGallery from "@/collection/CollectionGallery.vue";
import {setBC} from "@/util/vutils";

@Component({ components: { CollectionGallery } })
export default class About extends Vue {
  mounted(): void {
    setBC(this, {path: "", title: "about"});
  }
}
</script>
