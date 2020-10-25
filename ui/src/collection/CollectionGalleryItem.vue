<template>
  <li>
    <router-link class="theme collection-link" :to="'/c/' + coll.key">{{ label }}</router-link>
    <em v-if="coll.description.length > 0" class="ml">{{ coll.description }}</em>
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
