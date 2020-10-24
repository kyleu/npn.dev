<template>
  <header>
    <div data-uk-sticky="sel-target: .uk-navbar-container; cls-active: data-uk-navbar-sticky; media: 960">
      <nav id="navbar" class="uk-navbar-container bluegrey-bg" data-uk-navbar>
        <div id="breadcrumbs" class="uk-navbar-left">
          <a title="<%= npncore.AppName %>" class="uk-navbar-item uk-logo" href="/"><svg width="32px" height="32px" viewBox="-0 0 68 68" xmlns="http://www.w3.org/2000/svg">
            <g fill="none">
              <path id="logo-symbol" d="M 50.655 0 L 50.611 12.31 L 30.603 26.31 M 30.603 42.31 L 50.611 56.31 L 50.611 68.048 M 2 34.371 L 28.902 34.31 M 28.902 17.31 L 28.902 51.31 M 30.303 51.31 L 30.303 17.31 M 40.607 52.31 L 43.609 48.31 L 47.61 54.31 L 40.607 52.31 Z M 8.594 33.31 C 9.364 11.769 33.173 -0.86 51.451 10.577 C 69.728 22.014 68.766 48.94 49.718 59.043 C 31.449 68.734 9.332 55.971 8.594 35.31" style="stroke-width: 5px; paint-order: fill; stroke: rgb(135, 135, 135);" />
            </g>
          </svg></a>
          <a class="uk-navbar-item uk-logo uk-margin-remove uk-padding-remove dynamic" href="/">npn</a>
        </div>
        <div class="uk-navbar-right">
          <ul class="uk-navbar-nav">
            <li class="uk-margin-small-right">
              <a href="/profile" data-uk-icon="icon:user" title="Profile"></a>
            </li>
          </ul>
          <a href="" data-uk-toggle="target: #nav-offcanvas;" data-uk-navbar-toggle-icon="" class="uk-hidden@m uk-icon uk-margin-right"></a>
        </div>
      </nav>
    </div>
  </header>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component
export default class NavBar extends Vue {}
</script>

<style lang="scss">
a {
  &.router-link-exact-active {
    font-weight: bold;
  }
}
</style>
