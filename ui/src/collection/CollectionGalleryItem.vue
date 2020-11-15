<template>
  <li>
    <router-link class="collection-link" :to="'/c/' + cc.coll.key" :title="cc.count + ' ' + (cc.count === 1 ? 'request' : 'requests')">
      <div class="left">
        {{ label }}
      </div>
      <div class="right">
        <em v-if="cc.coll.description && cc.coll.description.length > 0" class="ml mr">{{ cc.coll.description }}</em>
        <span class="uk-badge" :title="cc.count + ' ' + (cc.count === 1 ? 'request' : 'requests')">{{ cc.count }}</span>
      </div>
    </router-link>
    <div class="clear"></div>
  </li>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {CollectionCount} from "@/collection/collection";

@Component
export default class CollectionGalleryItem extends Vue {
  @Prop() cc!: CollectionCount;

  get label(): string {
    return this.cc.coll.title.length === 0 ? this.cc.coll.key : this.cc.coll.title;
  }
}
</script>
