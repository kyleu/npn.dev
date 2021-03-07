<template>
  <div>
    <div class="theme-section left uk-text-center">
      <div>Mode</div>
      <div class="left">
        <div class="mode-swatch" style="background-color: #fff;" @click="setTheme('light')"></div>
        <div @click="setTheme('light')">Light</div>
      </div>
      <div class="left">
        <div class="mode-swatch" style="background-color: #222;" @click="setTheme('dark')"></div>
        <div @click="setTheme('dark')">Dark</div>
      </div>
    </div>

    <div class="theme-section left uk-text-center">
      <div>Nav</div>
      <div class="left">
        <input v-model="settings.navB" class="color-input" type="color" />
        <div @click="swap('navB')">BG</div>
      </div>
      <div class="left">
        <input v-model="settings.navF" class="color-input" type="color" />
        <div @click="swap('navF')">FG</div>
      </div>
    </div>

    <div class="theme-section left uk-text-center">
      <div>Menu</div>
      <div class="left">
        <input v-model="settings.menuB" class="color-input" type="color" />
        <div @click="swap('menuB')">BG</div>
      </div>
      <div class="left">
        <input v-model="settings.menuF" class="color-input" type="color" />
        <div @click="swap('menuF')">FG</div>
      </div>
      <div class="left">
        <input v-model="settings.menuL" class="color-input" type="color" />
        <div @click="swap('menuL')">Link</div>
      </div>
    </div>

    <div class="theme-section left uk-text-center">
      <div>Body</div>
      <div class="left">
        <input v-model="settings.bodyB" class="color-input" type="color" />
        <div @click="swap('bodyB')">BG</div>
      </div>
      <div class="left">
        <input v-model="settings.bodyL" class="color-input" type="color" />
        <div @click="swap('bodyL')">Link</div>
      </div>
    </div>

    <div class="clear"></div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { profileRef, UserSettings } from "@/user/profile";

import { socketRef } from "@/socket/socket";
import { systemService } from "@/util/services";
import { clientCommands } from "@/util/command";
import {logDebug} from "@/util/log";

@Component
export default class Theme extends Vue {
  src = "";

  get settings(): UserSettings | undefined {
    return profileRef.value?.settings;
  }

  setTheme(m: string): void {
    if (profileRef.value) {
      profileRef.value.settings.mode = m;
    }
  }

  getColor(t: string): string {
    if (!this.settings) {
      return "?";
    }
    switch (t) {
      case "bodyB":
        return this.settings.bodyB;
      case "bodyL":
        return this.settings.bodyL;
      case "navB":
        return this.settings.navB;
      case "navF":
        return this.settings.navF;
      case "menuB":
        return this.settings.menuB;
      case "menuF":
        return this.settings.menuF;
      case "menuL":
        return this.settings.menuL;
    }
    return "invalid";
  }

  setColor(t: string, c: string): void {
    if (!this.settings) {
      return;
    }
    switch (t) {
      case "bodyB":
        this.settings.bodyB = c;
        break;
      case "bodyL":
        this.settings.bodyL = c;
        break;
      case "navB":
        this.settings.navB = c;
        break;
      case "navF":
        this.settings.navF = c;
        break;
      case "menuB":
        this.settings.menuB = c;
        break;
      case "menuF":
        this.settings.menuF = c;
        break;
      case "menuL":
        this.settings.menuL = c;
        break;
    }
  }
}
</script>
