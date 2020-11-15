<template>
  <div class="mockup-container mt mr" @click.prevent="applyTheme">
    <div>
      <em>{{ theme.key }}</em>
    </div>
    <div class="mockup" :class="theme.mode === 'light' ? 'uk-dark' : 'uk-light'">
      <div class="bg" :style="{ backgroundColor: theme.bodyB }"></div>
      <div class="nav" :style="{ backgroundColor: theme.navB, color: theme.navF }">
        <div class="txt" :style="{ color: theme.navF }">npn</div>
      </div>
      <div class="menu nav-section" :style="{ backgroundColor: theme.menuB }">
        <div class="nav-header" :style="{ color: theme.menuF }">Items</div>
        <a href="" onclick="return false" :style="{ color: theme.menuL }">Item 1</a>
        <a href="" onclick="return false" :style="{ color: theme.menuL }">Item 2</a>
        <a href="" onclick="return false" :style="{ color: theme.menuL }">Item 3</a>
        <a href="" onclick="return false" :style="{ color: theme.menuL }">Item 4</a>
      </div>
      <div class="card uk-card uk-card-body uk-card-default" :class="theme.mode === 'light' ? 'uk-dark' : 'uk-light'">
        <div>Lorem ipsum dolor sit <a href="" onclick="return false;" :style="{ color: theme.bodyL }">amet</a>, consectetur adipiscing elit.</div>
        <div class="button uk-button uk-button-default" :style="{ borderColor: theme.bodyL, color: theme.bodyL }">Button</div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {ThemeColors} from "@/user/themes";
import {profileRef, tempThemeRef} from "@/user/profile";

@Component
export default class Mockup extends Vue {
  @Prop() theme!: ThemeColors

  applyTheme(): void {
    tempThemeRef.value = this.theme.key;

    if(profileRef.value) {
      profileRef.value.settings = {
        mode: this.theme.mode,
        navB: this.theme.navB,
        navF: this.theme.navF,
        menuB: this.theme.menuB,
        menuF: this.theme.menuF,
        menuL: this.theme.menuL,
        bodyB: this.theme.bodyB,
        bodyL: this.theme.bodyL
      };
    }
  }
}
</script>
