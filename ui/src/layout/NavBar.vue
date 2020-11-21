<template>
  <header>
    <div data-uk-sticky="sel-target: .uk-navbar-container; cls-active: data-uk-navbar-sticky; media: 960">
      <nav id="navbar" v-style-nav class="uk-navbar-container" data-uk-navbar>
        <Breadcrumbs />
        <div class="uk-navbar-right">
          <ul class="uk-navbar-nav">
            <li class="uk-margin-small-right">
              <router-link v-if="(!profile.picture) || profile.picture.length === 0 || profile.picture === 'none'" v-style-nav-link to="/u" title="Profile"><Icon icon="user" /></router-link>
              <router-link v-else to="/u" title="Profile"><img class="uk-border-circle" alt="user profile" :src="profile.picture" /></router-link>
            </li>
          </ul>
          <a v-style-nav-link href="" data-uk-toggle="target: #nav-offcanvas;" class="uk-hidden@m uk-margin-right"><Icon icon="toggle" /></a>
        </div>
      </nav>
    </div>
  </header>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import Breadcrumbs from "@/layout/Breadcrumbs.vue";
import {Profile, profileRef} from "@/user/profile";
import Icon from "@/util/Icon.vue";

@Component({ components: {Icon, Breadcrumbs } })
export default class NavBar extends Vue {
  get profile(): Profile | undefined {
    return profileRef.value;
  }
}
</script>
