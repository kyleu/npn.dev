<template>
  <div>
    <router-link class="theme" :to="'/c/' + coll.key">{{ coll.title.length === 0 ? coll.key : coll.title }}</router-link>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import {Collection} from "@/model/collection";

@Component
export default class CollectionListItem extends Vue {
  @Prop() coll!: Collection;
}
</script>
