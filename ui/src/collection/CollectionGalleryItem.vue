<template>
  <div>
    <router-link class="theme" :to="'/c/' + coll.key">{{ label }}</router-link>
  </div>
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
