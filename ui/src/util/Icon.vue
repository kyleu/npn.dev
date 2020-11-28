<template>
  <span class="uk-icon" v-html="svg" />
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import { icons } from "@/util/icons";
import {logWarn} from "@/util/log";

@Component
export default class Icon extends Vue {
  @Prop() icon!: string;

  get svg(): string {
    const x: string = icons[this.icon];
    if (!x) {
      logWarn("missing icon [" + this.icon + "]");
      return icons["unknown"];
    }
    return x;
  }
}
</script>
