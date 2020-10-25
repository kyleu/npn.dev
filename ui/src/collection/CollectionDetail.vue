<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <h3>Collection [{{ $route.params.coll }}]</h3>
      <p>TODO!</p>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {Collection} from "@/collection/collection";

@Component
export default class CollectionDetail extends Vue {
  @Prop() coll!: Collection
}
</script>
