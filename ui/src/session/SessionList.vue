<template>
  <div>
    <div class="mt">
      <div class="uk-inline" style="width: 100%;">
        <a class="uk-form-icon uk-form-icon-flip" title="add session" href="" @click.prevent="addSession()"><Icon icon="plus" /></a>
        <form @submit.prevent="addSession()">
          <input id="sess-add-input" class="uk-input" type="text" placeholder="Add session" data-lpignore="true" />
        </form>
      </div>
    </div>
    <ul class="uk-list uk-list-divider mt">
      <li v-for="s in sessions" :key="s.key">
        <div class="left mrs">
          <input v-model="active" name="active-session" type="radio" :value="s.key" class="uk-radio" :title="'set [' + ((!s.title) || s.title.length === 0) ? s.key : s.title + '] as the active session'" />
        </div>
        <router-link class="session-link mt" :to="'/s/' + s.key">
          {{ ((!s.title) || s.title.length === 0) ? s.key : s.title }}
        </router-link>
      </li>
    </ul>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {SessionSummary} from "@/session/model";
import {activeSessionRef, sessionSummariesRef} from "@/session/state";
import {socketRef} from "@/socket/socket";
import {sessionService} from "@/util/services";
import {clientCommands} from "@/util/command";
import Icon from "@/util/Icon.vue";

@Component({ components: { Icon } })
export default class SessionList extends Vue {
  get sessions(): SessionSummary[] {
    return sessionSummariesRef.value;
  }

  addSession(): void {
    const el = document.getElementById("sess-add-input") as HTMLInputElement;
    const title = el.value.trim();
    if (socketRef.value) {
      socketRef.value.send({channel: sessionService.key, cmd: clientCommands.addSession, param: title});
    }
  }

  get active(): string {
    return activeSessionRef.value;
  }
  set active(s: string) {
    activeSessionRef.value = s;
  }
}
</script>
