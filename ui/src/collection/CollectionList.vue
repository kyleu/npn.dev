<template>
  <div class="nav-list">
    <CollectionListItem v-for="cc in collections" :key="cc.coll.key" :cc="cc" />
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import CollectionListItem from "@/collection/CollectionListItem.vue";
import {CollectionCount} from "@/collection/collection";
import {collectionsRef} from "@/collection/state";

@Component({ components: { CollectionListItem } })
export default class CollectionList extends Vue {
  get collections(): CollectionCount[] {
    return collectionsRef.value
  }
}
</script>
