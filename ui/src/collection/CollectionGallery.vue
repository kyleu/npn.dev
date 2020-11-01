<template>
  <ul class="uk-list uk-list-divider mt">
    <CollectionGalleryItem v-for="coll in collections" :key="coll.key" :coll="coll" />
  </ul>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import CollectionGalleryItem from "@/collection/CollectionGalleryItem.vue";
import {Collection} from "@/collection/collection";
import {collectionsRef} from "@/state/state";

@Component({ components: { CollectionGalleryItem } })
export default class CollectionGallery extends Vue {
  get collections(): Collection[] {
    return collectionsRef.value
  }
}
</script>
