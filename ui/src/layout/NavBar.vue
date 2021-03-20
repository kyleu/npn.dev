<template>
  <header>
    <div data-uk-sticky="sel-target: .uk-navbar-container; cls-active: data-uk-navbar-sticky; media: 960">
      <nav id="navbar" v-style-nav class="uk-navbar-container" data-uk-navbar>
        <Breadcrumbs />
        <div class="uk-navbar-right">
          <ul class="uk-navbar-nav">
            <li>
              <a v-style-nav-link class="uk-navbar-toggle" data-uk-toggle href=""><Icon icon="search" /></a>
              <div ref="searchContainer" class="uk-drop" data-uk-drop="mode: click; pos: left-center; offset: 0">
                <form class="uk-search uk-search-navbar uk-width-1-1" @submit.prevent="onSearch">
                  <input ref="searchInput" class="uk-search-input" type="search" placeholder="Search" autofocus />
                </form>
              </div>
            </li>
            <li v-if="pub" class="header-optional"><a v-style-nav-link href="/">About</a></li>
            <li v-if="pub" class="header-optional"><a v-style-nav-link href="https://github.com/kyleu/npn">GitHub</a></li>
            <li v-if="pub" class="header-optional"><a v-style-nav-link href="/download"><div class="download-link">Download</div></a></li>
            <li class="mrs">
              <router-link v-if="(!profile.picture) || profile.picture.length === 0 || profile.picture === 'none'" v-style-nav-link to="/u" title="Profile"><Icon icon="user" /></router-link>
              <router-link v-else to="/u" title="Profile"><img class="uk-border-circle" alt="user profile" :src="profile.picture" /></router-link>
            </li>
            <li>
              <a v-style-nav-link href="" data-uk-toggle="target: #nav-offcanvas;" class="uk-hidden@m mr"><Icon icon="toggle" /></a>
            </li>
          </ul>
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
import {isPublic} from "@/util/log";
import UIkit from "uikit";

@Component({ components: {Icon, Breadcrumbs } })
export default class NavBar extends Vue {
  get pub(): boolean {
    return isPublic();
  }

  get profile(): Profile | undefined {
    return profileRef.value;
  }

  onSearch(): void {
    const i = this.$refs["searchInput"] as HTMLInputElement;
    i.blur();
    const c = this.$refs["searchContainer"] as HTMLElement;
    const f = UIkit.drop(c) as { hide: (b: boolean) => void };
    f.hide(false);
    const params = {q: i.value};
    this.$router.push({name: "SearchResults", params});
  }
}
</script>
