<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right">
          <router-link to="/help"><Icon icon="close" /></router-link>
        </div>
        <h3 class="uk-card-title"><Icon icon="npn" class="nav-icon-h3" /> npn</h3>
        <p>You're using <a href="https://npn.dev" target="_blank">npn</a>, an HTTP client that helps you document and test services</p>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { setBC } from "@/util/vutils";
import Icon from "@/util/Icon.vue";

@Component({ components: { Icon } })
export default class About extends Vue {
  mounted(): void {
    setBC(this, { path: "", title: "about" });
  }
}
</script>
