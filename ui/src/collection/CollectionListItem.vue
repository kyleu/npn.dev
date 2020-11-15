<template>
  <div class="nav-link">
    <router-link v-style-menu-link class="collection-link" :to="'/c/' + cc.coll.key" :title="cc.count + ' ' + (cc.count === 1 ? 'request' : 'requests')">
      <span class="uk-icon nav-icon" :data-uk-icon="icon"></span>
      {{ label }}
    </router-link>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {CollectionCount} from "@/collection/collection";

@Component
export default class CollectionListItem extends Vue {
  @Prop() cc!: CollectionCount;

  get label(): string {
    return ((!this.cc.coll.title) || this.cc.coll.title.length === 0) ? this.cc.coll.key : this.cc.coll.title;
  }

  get icon(): string {
    return "icon: " + (this.$route.params.coll === this.cc.coll.key ? "album" : "folder");
  }
}
</script>
