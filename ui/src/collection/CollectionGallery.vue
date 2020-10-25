<template>
  <div class="nav-list collection-list">
    <CollectionGalleryItem v-for="coll in $store.state.collections" :key="coll.key" :coll="coll" />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import CollectionGalleryItem from "@/collection/CollectionGalleryItem.vue";

@Component({ components: { CollectionGalleryItem } })
export default class CollectionGallery extends Vue {}
</script>
