<template>
  <span class="uk-icon nav-icon" v-html="svg" />
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {icons} from "@/util/icons";

@Component
export default class Icon extends Vue {
  @Prop() icon!: string

  get svg(): string {
    const x: string = icons[this.icon];
    if (!x) {
      console.warn("missing icon [" + this.icon + "]")
      return icons["unknown"];
    }
    return x;
  }
}
</script>
