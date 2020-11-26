<template>
  <ul class="uk-list uk-list-divider mt">
    <SessionListItem v-for="s in sessions" :key="s.key" :sess="s" />
  </ul>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {SessionSummary} from "@/session/model";
import SessionListItem from "@/session/SessionListItem.vue";
import {sessionSummariesRef} from "@/session/state";

@Component({ components: { SessionListItem  } })
export default class SessionList extends Vue {
  get sessions(): SessionSummary[] {
    return sessionSummariesRef.value
  }
}
</script>
