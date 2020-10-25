<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right"><router-link class="theme uk-icon" data-uk-icon="close" to="/"></router-link></div>
        <h3 class="uk-card-title">Collections!</h3>
        <CollectionGallery />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import CollectionGallery from "@/collection/CollectionGallery.vue";
@Component({
  components: {CollectionGallery}
})
export default class CollectionIndex extends Vue {}
</script>
