<template>
  <div class="nav-section">
    <router-link v-style-menu-link :to="{name: 'SessionIndex'}" onclick="return false;">
      <Icon icon="database" /> {{ sessTitle }}
    </router-link>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import Icon from "@/util/Icon.vue";
import {activeSessionRef, sessionsRef, SessionSummary} from "@/session/session";

@Component({ components: {Icon } })
export default class SessionSelector extends Vue {
  get sessions(): SessionSummary[] {
    return sessionsRef.value;
  }

  get sessTitle(): string {
    let t = activeSessionRef.value;
    if (t.length === 0) {
      t = "_";
    }
    const x = sessionsRef.value.find(x => x.key === t);
    if (!x) {
      if (t === "_") {
        return "Default Session"
      }
      return t;
    }
    return (x.title && x.title.length > 0) ? x.title : x.key;
  }
}
</script>
