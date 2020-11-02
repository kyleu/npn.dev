<template>
  <div class="nav-list">
    <CollectionListItem v-for="coll in collections" :key="coll.key" :coll="coll" />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import CollectionListItem from "@/collection/CollectionListItem.vue";
import {Collection} from "@/collection/collection";
import {collectionsRef} from "@/collection/state";

@Component({ components: { CollectionListItem } })
export default class CollectionList extends Vue {
  get collections(): Collection[] {
    return collectionsRef.value
  }
}
</script>
