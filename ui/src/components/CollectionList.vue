<template>
  <div class="nav-section">
    <div class="nav-header">
      Collections
    </div>
    <div class="nav-list collection-list">
      <CollectionListItem v-for="coll in colls" :key="coll.key" :coll="coll" />
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import CollectionListItem from "@/components/CollectionListItem.vue";
import {Collection} from "@/model/collection";

@Component({ components: { CollectionListItem } })
export default class CollectionList extends Vue {
  colls: Collection[] = [
    {
      key: "test1",
      title: "Test 1",
      description: "Test Collection 1"
    },
    {
      key: "test2",
      title: "Test 2",
      description: "Test Collection 2"
    }
  ];
}
</script>
