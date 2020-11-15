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
        <v-swatches v-model="settings.navB" swatches="text-advanced" show-fallback></v-swatches>
        <div @click="swap('navB')">BG</div>
      </div>
      <div class="left">
        <v-swatches v-model="settings.navF" swatches="text-advanced" show-fallback></v-swatches>
        <div @click="swap('navF')">FG</div>
      </div>
    </div>

    <div class="theme-section left uk-text-center">
      <div>Menu</div>
      <div class="left">
        <v-swatches v-model="settings.menuB" swatches="text-advanced" show-fallback></v-swatches>
        <div @click="swap('menuB')">BG</div>
      </div>
      <div class="left">
        <v-swatches v-model="settings.menuF" swatches="text-advanced" show-fallback></v-swatches>
        <div @click="swap('menuF')">FG</div>
      </div>
      <div class="left">
        <v-swatches v-model="settings.menuL" swatches="text-advanced" show-fallback></v-swatches>
        <div @click="swap('menuL')">Link</div>
      </div>
    </div>

    <div class="theme-section left uk-text-center">
      <div>Body</div>
      <div class="left">
        <v-swatches v-model="settings.bodyB" swatches="text-advanced" show-fallback></v-swatches>
        <div @click="swap('bodyB')">BG</div>
      </div>
      <div class="left">
        <v-swatches v-model="settings.bodyL" swatches="text-advanced" show-fallback></v-swatches>
        <div @click="swap('bodyL')">Link</div>
      </div>
    </div>

    <div class="clear"></div>
    <div v-if="tempTheme.length > 0" class="mt">
      <button v-style-button class="uk-button uk-button-default" @click="saveTheme()">DEBUG / Save Theme</button>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import {profileRef, tempThemeRef, UserSettings} from "@/user/profile";

// @ts-ignore
// eslint-disable-next-line
import VSwatches from 'vue-swatches'
import {socketRef} from "@/socket/socket";
import {systemService} from "@/util/services";
import {clientCommands} from "@/util/command";

@Component({ components: { VSwatches } })
export default class ColorPicker extends Vue {
  src = ""

  get tempTheme(): string {
    return tempThemeRef.value;
  }

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

  swap(t: string): void {
    if (this.src.length === 0) {
      this.src = t;
      return;
    }
    const s = this.src;
    this.src = "";
    console.log(`SWAP: ${s} <-> ${t}`);
    const src = this.getColor(s);
    const tgt = this.getColor(t);
    this.setColor(s, tgt);
    this.setColor(t, src);
  }

  saveTheme(): void {
    if (!this.settings) {
      return;
    }
    const theme = `const ${this.tempTheme} = {
  key: "${this.tempTheme}",
  mode: "${this.settings.mode}",
  bodyB: "${this.settings.bodyB}",
  bodyL: "${this.settings.bodyL}",
  navB: "${this.settings.navB}",
  navF: "${this.settings.navF}",
  menuB: "${this.settings.menuB}",
  menuF: "${this.settings.menuF}",
  menuL: "${this.settings.menuL}"
}`
    if (socketRef.value) {
      socketRef.value.send({svc: systemService.key, cmd: clientCommands.testbed, param: {t: "theme", k: this.tempTheme, v: theme}});
    }

    console.log(theme);
  }
}
</script>
