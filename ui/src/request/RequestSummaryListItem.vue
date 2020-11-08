<template>
  <li>
    <router-link :class="'req-link ' + profile.settings.linkColor + '-fg'" :to="'/c/' + coll + '/' + req.key">
      <div class="right"><em>{{ req.url }}</em></div>
      {{ label }}
    </router-link>
  </li>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import {Summary} from "@/request/model";
import {Profile, profileRef} from "@/user/profile";

@Component
export default class RequestSummaryListItem extends Vue {
  @Prop() coll!: string;
  @Prop() req!: Summary;

  get profile(): Profile | undefined {
    return profileRef.value;
  }

  get label(): string {
    return (!this.req.title || this.req.title.length === 0) ? this.req.key : this.req.title;
  }
}
</script>

<style lang="scss">
  .req-link.router-link-active {
    font-weight: bold;
  }
</style>
