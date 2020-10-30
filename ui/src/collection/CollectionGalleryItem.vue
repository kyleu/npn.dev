<template>
  <li>
    <em v-if="coll.description && coll.description.length > 0" class="right ml">{{ coll.description }}</em>
    <router-link :class="'collection-link ' + $store.state.profile.linkColor + '-fg'" :to="'/c/' + coll.key">{{ label }}</router-link>
  </li>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import {Collection} from "@/collection/collection";

@Component
export default class CollectionGalleryItem extends Vue {
  @Prop() coll!: Collection;

  get label(): string {
    return this.coll.title.length === 0 ? this.coll.key : this.coll.title;
  }
}
</script>
