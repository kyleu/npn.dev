<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right"><router-link class="theme uk-icon" data-uk-icon="close" to="/c"></router-link></div>
        <h3 class="uk-card-title">Collection [{{ $route.params.coll }}]</h3>
        <p>TODO!</p>
      </div>
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
