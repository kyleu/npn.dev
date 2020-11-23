<template>
  <ul class="uk-list uk-list-divider mt">
    <SessionListItem v-for="s in sessions" :key="s.key" :sess="s" />
  </ul>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {sessionsRef, SessionSummary} from "@/session/session";
import SessionListItem from "@/session/SessionListItem.vue";

@Component({ components: { SessionListItem  } })
export default class SessionList extends Vue {
  get sessions(): SessionSummary[] {
    return sessionsRef.value
  }
}
</script>
