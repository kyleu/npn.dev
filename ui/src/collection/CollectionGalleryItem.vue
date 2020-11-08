<template>
  <li>
    <div class="right">
      <em v-if="cc.coll.description && cc.coll.description.length > 0" class="ml mr">{{ cc.coll.description }}</em>
      <span class="uk-badge" :title="cc.count + ' requests'">{{ cc.count }}</span>
    </div>
    <router-link :class="'collection-link ' + profile.settings.linkColor + '-fg'" :to="'/c/' + cc.coll.key" :title="cc.count + ' requests'">
      {{ label }}
    </router-link>
  </li>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {CollectionCount} from "@/collection/collection";
import {Profile, profileRef} from "@/user/profile";

@Component
export default class CollectionGalleryItem extends Vue {
  @Prop() cc!: CollectionCount;

  get profile(): Profile | undefined {
    return profileRef.value;
  }

  get label(): string {
    return this.cc.coll.title.length === 0 ? this.cc.coll.key : this.cc.coll.title;
  }
}
</script>
