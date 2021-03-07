<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <h3 class="uk-card-title">Profile</h3>
        <div class="mt">
          <label class="uk-form-label">Name <input v-model="profile.name" class="uk-input" name="name" type="text" data-lpignore="true" /></label>
        </div>
        <div class="mt">
          <label class="uk-form-label">
            Locale
            <select v-model="profile.locale" class="uk-select">
              <option value="en-US">American English</option>
              <option value="en-US" disabled="disabled">More options soon!</option>
            </select>
          </label>
        </div>
        <div class="mt">
          <label class="uk-form-label">Theme</label>
          <Theme />
        </div>

        <div v-if="different" class="mt">
          <button v-style-button class="uk-button uk-button-default" @click="saveSettings()">Save</button>
          <button v-style-button class="uk-button uk-button-default ml" @click="resetSettings()">Reset</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {setBC} from "@/util/vutils";
import {Profile, profileRef, UserSettings} from "@/user/profile";
import {jsonClone} from "@/util/json";
import {allThemes, ThemeColors} from "@/user/themes";
import Theme from "@/user/Theme.vue";
import {socketRef} from "@/socket/socket";
import {systemService} from "@/util/services";
import {clientCommands} from "@/util/command";

@Component({ components: { Theme } })
export default class ProfileView extends Vue {
  original: Profile | undefined = jsonClone(profileRef.value)

  get themes(): ThemeColors[] {
    return allThemes;
  }

  get profile(): Profile | undefined {
    return profileRef.value;
  }

  get settings(): UserSettings | undefined {
    return profileRef.value?.settings;
  }

  resetSettings(): void {
    if (profileRef.value) {
      profileRef.value = jsonClone(this.original);
    }
  }

  saveSettings(): void {
    if (profileRef.value) {
      this.original = jsonClone(profileRef.value);
      if (socketRef.value) {
        socketRef.value.send({channel: systemService.key, cmd: clientCommands.saveProfile, param: this.profile});
      }
    }
  }

  get different(): boolean {
    if (!this.original) {
      this.original = jsonClone(profileRef.value);
    }

    if (profileRef.value?.name !== this.original?.name) {
      return true;
    }
    if (profileRef.value?.picture !== this.original?.picture) {
      return true;
    }
    if (profileRef.value?.locale !== this.original?.locale) {
      return true;
    }

    const s = profileRef.value?.settings;
    const o = this.original?.settings;
    if (!s || !o) {
      return false;
    }
    if (s.mode !== o.mode) {
      return true;
    }
    if (s.navB !== o.navB) {
      return true;
    }
    if (s.navF !== o.navF) {
      return true;
    }
    if (s.menuB !== o.menuB) {
      return true;
    }
    if (s.menuF !== o.menuF) {
      return true;
    }
    if (s.menuL !== o.menuL) {
      return true;
    }
    if (s.bodyB !== o.bodyB) {
      return true;
    }
    if (s.bodyL !== o.bodyL) {
      return true;
    }
    //
    return false;
  }

  mounted(): void {
    setBC(this, {path: "/u", title: "settings"});
  }
}
</script>
