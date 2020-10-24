<template>
  <div class="nav-list collection-list">
    <CollectionListItem v-for="coll in $store.state.collections" :key="coll.key" :coll="coll" />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import CollectionListItem from "@/components/CollectionListItem.vue";

@Component({ components: { CollectionListItem } })
export default class CollectionList extends Vue {}
</script>
