<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right">
          <router-link to="/"><Icon icon="close" /></router-link>
        </div>
        <h3 class="uk-card-title"><Icon icon="question" class="nav-icon-h3" /> Help</h3>
        <p><a href="https://npn.dev" target="_blank">npn</a> is <a href="https://github.com/kyleu/npn" target="_blank">open source</a></p>
        <p>You can file issues <a href="https://github.com/KyleU/npn/issues" target="_blank">here</a></p>
        <ul class="uk-list uk-list-divider">
          <li><router-link to="/about">About</router-link></li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { setBC } from "@/util/vutils";
import Icon from "@/util/Icon.vue";

@Component({ components: { Icon } })
export default class Help extends Vue {
  mounted(): void {
    setBC(this, { path: "", title: "help" });
  }
}
</script>
