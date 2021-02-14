<template>
  <div class="nav-list">
    <div v-if="(!collections) || (collections.length === 0)" class="nav-link">No collections</div>
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
    return collectionsRef.value.filter(x => x.coll.key !== "_");
  }
}
</script>
