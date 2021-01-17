<template>
  <div id="breadcrumbs" class="uk-navbar-left">
    <router-link title="npn" class="uk-navbar-item uk-logo" to="/" exact>
      <svg width="32px" height="32px" viewBox="-0 0 68 68" xmlns="http://www.w3.org/2000/svg">
        <g fill="none">
          <path id="logo-symbol" :style="{ strokeWidth: '5px', paintOrder: 'fill', stroke: profile.settings.navF }" d="M 50.655 0 L 50.611 12.31 L 30.603 26.31 M 30.603 42.31 L 50.611 56.31 L 50.611 68.048 M 2 34.371 L 28.902 34.31 M 28.902 17.31 L 28.902 51.31 M 30.303 51.31 L 30.303 17.31 M 40.607 52.31 L 43.609 48.31 L 47.61 54.31 L 40.607 52.31 Z M 8.594 33.31 C 9.364 11.769 33.173 -0.86 51.451 10.577 C 69.728 22.014 68.766 48.94 49.718 59.043 C 31.449 68.734 9.332 55.971 8.594 35.31" />
        </g>
      </svg>
    </router-link>
    <router-link v-style-nav-link title="npn" class="uk-navbar-item uk-logo uk-margin-remove uk-padding-remove dynamic" to="/" exact>npn</router-link>

    <span v-for="bc in breadcrumbs" :key="bc.path" class="uk-navbar-item trim dynamic">
      <span v-style-nav-link :style="{ padding: '0 8px' }"> / </span>
      <router-link v-style-nav-link class="uk-navbar-item uk-logo uk-margin-remove uk-padding-remove dynamic" :to="bc.path" exact>{{ bc.title }}</router-link>
    </span>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import {Breadcrumb, breadcrumbsRef} from "@/layout/breadcrumb";
import {Profile, profileRef} from "@/user/profile";

@Component
export default class Breadcrumbs extends Vue {
  get breadcrumbs(): Breadcrumb[] {
    if (!breadcrumbsRef.value) {
      return [];
    }
    return breadcrumbsRef.value;
  }

  get profile(): Profile | undefined {
    return profileRef.value;
  }
}
</script>
