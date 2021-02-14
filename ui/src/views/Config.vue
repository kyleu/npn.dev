<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right">
          <router-link to="/"><Icon icon="close" /></router-link>
        </div>
        <h3 class="uk-card-title"><Icon icon="cog" class="nav-icon-h3" /> System Config</h3>
        <ul class="uk-list uk-list-divider">
          <li><a href="" onclick="npn.onDebug();return false">Debug</a></li>
          <li><router-link to="/testbed">Testbed</router-link></li>
          <li><a href="" @click.prevent="toggleLog()">Toggle Log</a></li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { setBC } from "@/util/vutils";
import Icon from "@/util/Icon.vue";
import {logToggle} from "@/util/log";

@Component({ components: { Icon } })
export default class Config extends Vue {
  toggleLog(): void {
    logToggle();
  }

  mounted(): void {
    setBC(this, { path: "", title: "config" });
  }
}
</script>
