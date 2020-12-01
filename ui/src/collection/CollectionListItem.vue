<template>
  <div class="nav-link">
    <router-link v-style-menu-link class="collection-link" :to="'/c/' + cc.coll.key" :title="cc.count + ' ' + (cc.count === 1 ? 'request' : 'requests')">
      <Icon :icon="icon" class="nav-icon" />
      {{ label }}
    </router-link>
    <div v-if="this.$route.params.coll === cc.coll.key && requests.length > 0">
      <div v-for="r in requests" :key="r.key" class="nav-request-link">
        <router-link v-style-menu-link :to="'/c/' + cc.coll.key + '/' + r.key">
          <Icon icon="link" class="nav-icon" />
          {{ (!r.title || r.title.length === 0) ? r.key : r.title }}
        </router-link>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {CollectionCount} from "@/collection/collection";
import {Summary} from "@/request/model";
import {collectionSummariesRef} from "@/collection/state";
import Icon from "@/util/Icon.vue";

@Component({ components: {Icon} })
export default class CollectionListItem extends Vue {
  @Prop() cc!: CollectionCount;

  get requests(): Summary[] {
    return collectionSummariesRef.value.find(x => x.key === this.cc.coll.key)?.requests || [];
  }

  get label(): string {
    return ((!this.cc.coll.title) || this.cc.coll.title.length === 0) ? this.cc.coll.key : this.cc.coll.title;
  }

  get icon(): string {
    if (this.$route.params.coll === this.cc.coll.key) {
      return "album";
    }
    return "folder";
  }
}
</script>
