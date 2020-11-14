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
          <label class="uk-form-label"><input v-model="settings.mode" type="radio" class="uk-radio" value="light" /> Light</label>
          <label class="uk-form-label"><input v-model="settings.mode" type="radio" class="uk-radio" value="dark" /> Dark</label>
        </div>

        <div class="mt">
          <label class="uk-form-label">Body Background <input v-model="settings.bodyB" type="text" class="uk-input" /></label>
        </div>
        <div class="mt">
          <label class="uk-form-label">Body Link <input v-model="settings.bodyL" type="text" class="uk-input" /></label>
        </div>

        <div class="mt">
          <label class="uk-form-label">Nav Background <input v-model="settings.navB" type="text" class="uk-input" /></label>
        </div>
        <div class="mt">
          <label class="uk-form-label">Nav Foreground <input v-model="settings.navF" type="text" class="uk-input" /></label>
        </div>

        <div class="mt">
          <label class="uk-form-label">Menu Background <input v-model="settings.menuB" type="text" class="uk-input" /></label>
        </div>
        <div class="mt">
          <label class="uk-form-label">Menu Foreground <input v-model="settings.menuF" type="text" class="uk-input" /></label>
        </div>
        <div class="mt">
          <label class="uk-form-label">Menu Link <input v-model="settings.menuL" type="text" class="uk-input" /></label>
        </div>

        <div class="mt">
          <label class="uk-form-label">Current Theme <input v-model="tempTheme" type="text" class="uk-input" /></label>
        </div>
        <ColorPicker />
      </div>
      <div class="uk-card uk-card-body uk-card-default mt">
        <h3 class="uk-card-title">Colors (temporary)</h3>
        <Mockup v-for="t in themes" :key="t.key" :theme="t" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import {setBC} from "@/util/vutils";
import {profileRef, tempThemeRef, UserSettings} from "@/user/profile";
import Mockup from "@/user/Mockup.vue";
import {jsonStr} from "@/util/json";
import {allThemes, ThemeColors} from "@/user/themes";

// @ts-ignore
// eslint-disable-next-line
import VSwatches from 'vue-swatches'
import 'vue-swatches/dist/vue-swatches.css'
import ColorPicker from "@/user/ColorPicker.vue";

@Component({ components: {ColorPicker, Mockup, VSwatches } })
export default class Settings extends Vue {
  get themes(): ThemeColors[] {
    return allThemes;
  }

  get dbg(): string {
    return jsonStr(profileRef.value);
  }

  get tempTheme(): string {
    return tempThemeRef.value;
  }

  get settings(): UserSettings | undefined {
    return profileRef.value?.settings;
  }

  mounted(): void {
    setBC(this, {path: "/u", title: "settings"});
  }
}
</script>
