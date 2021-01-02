<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right">
          <router-link to="/"><Icon icon="close" /></router-link>
        </div>
        <h3 class="uk-card-title"><Icon icon="search" class="nav-icon-h3" /> Search Results</h3>
        <p>TODO</p>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { setBC } from "@/util/vutils";
import Icon from "@/util/Icon.vue";

@Component({ components: { Icon } })
export default class SearchResults extends Vue {
  mounted(): void {
    setBC(this, { path: "", title: "search" });
  }
}
</script>
