<template>
  <div class="nav-link">
    <router-link :class="'collection-link ' + profile.settings.linkColor + '-fg'" :to="'/c/' + cc.coll.key" :title="cc.count + ' requests'">
      <span class="uk-icon nav-icon" :data-uk-icon="icon"></span>
      {{ label }}
    </router-link>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {CollectionCount} from "@/collection/collection";
import {Profile, profileRef} from "@/user/profile";

@Component
export default class CollectionListItem extends Vue {
  @Prop() cc!: CollectionCount;

  get profile(): Profile | undefined {
    return profileRef.value;
  }

  get label(): string {
    return ((!this.cc.coll.title) || this.cc.coll.title.length === 0) ? this.cc.coll.key : this.cc.coll.title;
  }

  get icon(): string {
    return "icon: " + (this.$route.params.coll === this.cc.coll.key ? "album" : "folder");
  }
}
</script>

<style lang="scss">
  .collection-link.router-link-active {
    font-weight: bold;
  }
</style>
