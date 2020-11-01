<template>
  <ul class="uk-list uk-list-divider">
    <li>
      <div data-uk-grid="">
        <div class="uk-width-1-4">Name</div>
        <div class="uk-width-1-4">Value</div>
        <div class="uk-width-1-2">
          <div class="right">
            <a :class="profile.linkColor + '-fg'" href="" onclick="return false" title="new header">
              <span data-uk-icon="icon: plus" />
            </a>
          </div>
          Description
        </div>
      </div>
    </li>
    <li v-for="h of headers" :key="h.k">
      <div data-uk-grid="">
        <div class="uk-width-1-4">
          <input v-model="h.k" class="uk-input" type="text" />
        </div>
        <div class="uk-width-1-4">
          <input v-model="h.v" class="uk-input" type="text" />
        </div>
        <div class="uk-width-1-2">
          <div class="right" style="margin-top: 6px;">
            <a :class="profile.linkColor + '-fg'" href="" onclick="return false;" title="remove param">
              <span data-uk-icon="icon: close" />
            </a>
          </div>
          <input v-model="h.desc" style="width: calc(100% - 36px);" class="uk-input" type="text" />
        </div>
      </div>
    </li>
  </ul>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {Header} from "@/header/model";
import Profile from "@/user/profile";
import {profileRef} from "@/state/state";

@Component
export default class HeadersEditor extends Vue {
  @Prop() headers!: Header[] | undefined;

  get profile(): Profile | undefined {
    return profileRef.value;
  }
}
</script>
