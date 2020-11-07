<template>
  <div class="nav-link">
    <router-link :class="'collection-link ' + profile.linkColor + '-fg'" :to="'/c/' + coll.key">
      <span class="uk-icon nav-icon" :data-uk-icon="icon"></span>
      {{ label }}
    </router-link>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import {Collection} from "@/collection/collection";
import {Profile, profileRef} from "@/user/profile";

@Component
export default class CollectionListItem extends Vue {
  @Prop() coll!: Collection;

  get profile(): Profile | undefined {
    return profileRef.value;
  }

  get label(): string {
    return ((!this.coll.title) || this.coll.title.length === 0) ? this.coll.key : this.coll.title;
  }

  get icon(): string {
    return "icon: " + (this.$route.params.coll === this.coll.key ? "album" : "folder");
  }
}
</script>

<style lang="scss">
  .collection-link.router-link-active {
    font-weight: bold;
  }
</style>
