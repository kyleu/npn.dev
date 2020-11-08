<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <h3 class="uk-card-title">Profile</h3>
        <pre>{{ dbg }}</pre>
      </div>
      <div class="uk-card uk-card-body uk-card-default mt">
        <h3 class="uk-card-title">Debug</h3>
        <div class="mt">
          <label class="uk-form-label">
            Nav Color
            <input v-model="profile.navColor" type="text" class="uk-input" />
          </label>
        </div>
        <div class="mt">
          <label class="uk-form-label">
            Link Color
            <input v-model="profile.linkColor" type="text" class="uk-input" />
          </label>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import CollectionGallery from "@/collection/CollectionGallery.vue";
import {setBC} from "@/util/vutils";
import {Profile, profileRef} from "@/user/profile";
import {jsonStr} from "@/util/json";

@Component({ components: { CollectionGallery } })
export default class Settings extends Vue {
  get dbg (): string {
    return jsonStr(profileRef.value);
  }

  get profile(): Profile | undefined {
    return profileRef.value;
  }

  mounted(): void {
    setBC(this);
  }
}
</script>
