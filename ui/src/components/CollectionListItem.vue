<template>
  <div><a href="">{{ coll.title.length === 0 ? coll.key : coll.title }}</a></div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import {Collection} from "@/model/collection";

@Component
export default class CollectionListItem extends Vue {
  @Prop(Object) coll: Collection | undefined
}
</script>
