<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right"><router-link :class="'uk-icon ' + profile.linkColor + '-fg'" data-uk-icon="close" to="/"></router-link></div>
        <h3 class="uk-card-title">Collections!</h3>
        <CollectionGallery />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import CollectionGallery from "@/collection/CollectionGallery.vue";
import {setBC} from "@/util/vutils";
import Profile from "@/user/profile";
import {profileRef} from "@/state/state";

@Component({components: {CollectionGallery}})
export default class CollectionIndex extends Vue {
  get profile(): Profile | undefined {
    return profileRef.value;
  }

  created(): void {
    setBC(this, {path: "", title: "collections"});
  }

  updated(): void {
    setBC(this, {path: "", title: "collections"});
  }
}
</script>
