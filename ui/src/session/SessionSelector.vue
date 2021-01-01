<template>
  <div class="nav-list">
    <div class="nav-link">
      <router-link v-style-menu-link :to="{name: 'SessionDetail', params: { sess: sessKey } }" onclick="return false;">
        <Icon icon="bookmark" class="nav-icon" />
        {{ sessTitle }}
      </router-link>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import Icon from "@/util/Icon.vue";
import {SessionSummary} from "@/session/model";
import {activeSessionRef, sessionSummariesRef} from "@/session/state";

@Component({ components: {Icon } })
export default class SessionSelector extends Vue {
  get sessions(): SessionSummary[] {
    return sessionSummariesRef.value;
  }

  get sessKey(): string {
    let t = activeSessionRef.value;
    if (t.length === 0) {
      t = "_";
    }
    return t;
  }

  get sessTitle(): string {
    const t = this.sessKey;
    const x = sessionSummariesRef.value.find(x => x.key === t);
    if (!x) {
      if (t === "_") {
        return "Default Session";
      }
      return t;
    }
    return (x.title && x.title.length > 0) ? x.title : x.key;
  }
}
</script>
