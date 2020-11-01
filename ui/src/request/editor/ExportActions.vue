<template>
  <div class="uk-inline">
    <button type="button" class="uk-button uk-button-default uk-margin-small-right mt">Export</button>
    <div id="export-dropdown" data-uk-dropdown="mode: click">
      <ul class="uk-list uk-list-divider" style="margin-bottom: 0;">
        <li v-for="(v, n) in transforms" :key="n">
          <router-link :class="profile.linkColor + '-fg'" :to="'/c/' + $route.params.coll + '/' + $route.params.req + '/transform/' + n" onclick="UIkit.dropdown(document.getElementById('export-dropdown')).hide(false);">{{ v }}</router-link>
        </li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import Profile from "@/user/profile";
import {profileRef} from "@/state/state";

@Component
export default class ExportActions extends Vue {
  get profile(): Profile | undefined {
    return profileRef.value;
  }

  get transforms(): Record<string, string> {
    return {
      "http": "HTTP",
      "json": "JSON",
      "curl": "curl"
    };
  }
}
</script>
