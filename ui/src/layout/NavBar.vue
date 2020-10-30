<template>
  <header>
    <div data-uk-sticky="sel-target: .uk-navbar-container; cls-active: data-uk-navbar-sticky; media: 960">
      <nav id="navbar" :class="'uk-navbar-container ' + $store.state.profile.navColor + '-bg'" data-uk-navbar>
        <Breadcrumbs />
        <div class="uk-navbar-right">
          <ul class="uk-navbar-nav">
            <li class="uk-margin-small-right">
              <a v-if="$store.state.profile.picture.length === 0 || $store.state.profile.picture === 'none'" href="/profile" data-uk-icon="icon:user" title="Profile"></a>
              <a v-if="$store.state.profile.picture.length !== 0 && $store.state.profile.picture !== 'none'" href="/profile" title="Profile">
                <img class="uk-border-circle" alt="user profile" :src="$store.state.profile.picture" />
              </a>
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

@Component({ components: { Breadcrumbs } })
export default class NavBar extends Vue {}
</script>
