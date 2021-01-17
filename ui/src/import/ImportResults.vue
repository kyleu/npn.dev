<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right">
          <router-link to="/i"><Icon icon="close" /></router-link>
        </div>
        <h3 class="uk-card-title"><Icon icon="upload" class="nav-icon-h3" /> Import Results</h3>
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
export default class ImportResults extends Vue {
  mounted(): void {
    setBC(this, { path: "/i", title: "import" }, { path: "", title: "results" });
  }
}
</script>
