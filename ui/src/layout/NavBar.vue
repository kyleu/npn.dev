<template>
  <header>
    <div data-uk-sticky="sel-target: .uk-navbar-container; cls-active: data-uk-navbar-sticky; media: 960">
      <nav id="navbar" :class="'uk-navbar-container ' + profile.navColor + '-bg'" data-uk-navbar>
        <Breadcrumbs />
        <div class="uk-navbar-right">
          <ul class="uk-navbar-nav">
            <li class="uk-margin-small-right">
              <router-link v-if="profile.picture.length === 0 || profile.picture === 'none'" to="/u" data-uk-icon="icon:user" title="Profile"></router-link>
              <router-link v-else to="/u" title="Profile"><img class="uk-border-circle" alt="user profile" :src="profile.picture" /></router-link>
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
import Breadcrumbs from "@/layout/Breadcrumbs.vue";
import {Profile, profileRef} from "@/user/profile";

@Component({ components: { Breadcrumbs } })
export default class NavBar extends Vue {
  get profile(): Profile | undefined {
    return profileRef.value;
  }
}
</script>
