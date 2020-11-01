<template>
  <div>
    <router-link :class="'collection-link ' + profile.linkColor + '-fg'" :to="'/c/' + coll.key">{{ label }}</router-link>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import {Collection} from "@/collection/collection";
import Profile from "@/user/profile";
import {profileRef} from "@/state/state";

@Component
export default class CollectionListItem extends Vue {
  @Prop() coll!: Collection;

  get profile(): Profile | undefined {
    return profileRef.value;
  }

  get label(): string {
    return this.coll.title.length === 0 ? this.coll.key : this.coll.title;
  }
}
</script>

<style lang="scss">
  .collection-link.router-link-active {
    font-weight: bold;
  }
</style>
