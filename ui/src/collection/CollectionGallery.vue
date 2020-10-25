<template>
  <ul class="uk-list uk-list-divider mt">
    <CollectionGalleryItem v-for="coll in $store.state.collections" :key="coll.key" :coll="coll" />
  </ul>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import CollectionGalleryItem from "@/collection/CollectionGalleryItem.vue";

@Component({ components: { CollectionGalleryItem } })
export default class CollectionGallery extends Vue {}
</script>
