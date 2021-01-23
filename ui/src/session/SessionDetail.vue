<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right"><router-link to="/s"><Icon icon="close" /></router-link></div>
        <h3 class="uk-card-title">
          <Icon class="nav-icon-h3" icon="bookmark" />
          {{ (sess && sess.title && sess.title.length > 0) ? sess.title : $route.params.sess }}
        </h3>
        <div v-if="showEditor.length !== 0" class="mt">
          <div v-if="$route.params.sess !== '_'">
            <label class="uk-form-label">
              Key
              <input v-model="sess.key" class="uk-input" name="key" type="text" data-lpignore="true" />
            </label>
          </div>
          <input v-else v-model="sess.key" name="key" type="hidden" data-lpignore="true" />

          <div class="mt">
            <label class="uk-form-label">
              Title
              <input v-model="sess.title" class="uk-input" name="title" type="text" data-lpignore="true" />
            </label>
          </div>
        </div>

        <div v-if="different" class="right">
          <button v-style-button class="uk-button uk-button-default mrs mt" @click="reset();">Reset</button>
          <button v-style-button class="uk-button uk-button-default mt" @click="save();">Save Changes</button>
        </div>

        <div v-if="showEditor.length === 0">
          <button v-style-button class="uk-button uk-button-default mrs mt" @click="editSession()">Edit</button>
          <router-link v-style-button class="uk-button uk-button-default mt" :to="'/x/' + $route.params.sess">Export</router-link>
        </div>
        <div v-else>
          <button v-style-button class="uk-button uk-button-default mrs mt" @click="cancelEdit()">Cancel</button>
          <button v-if="$route.params.sess !== '_'" v-style-button class="uk-button uk-button-default mt" @click="deleteSession()">Delete</button>
        </div>
      </div>
      <div class="uk-card uk-card-body uk-card-default mt">
        <h3 class="uk-card-title">Cookies</h3>
        <CookieEditor />
      </div>
      <div class="uk-card uk-card-body uk-card-default mt">
        <h3 class="uk-card-title">Variables</h3>
        <VariablesEditor />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {setBC} from "@/util/vutils";
import RequestSummaryList from "@/request/RequestSummaryList.vue";
import Icon from "@/util/Icon.vue";
import {Session} from "@/session/model";
import {sessionEditingRef, sessionOriginalRef, setActiveSession} from "@/session/state";
import VariablesEditor from "@/session/VariablesEditor.vue";
import {jsonClone} from "@/util/json";
import {socketRef} from "@/socket/socket";
import {sessionService} from "@/util/services";
import {clientCommands} from "@/util/command";
import {diffSessions} from "@/session/diff";
import CookieEditor from "@/session/CookieEditor.vue";

@Component({ components: {CookieEditor, VariablesEditor, Icon, RequestSummaryList } })
export default class SessionDetail extends Vue {
  showEditor = "";

  get sess(): Session | undefined {
    setActiveSession(this.$route.params.sess);
    return sessionEditingRef.value;
  }

  get different(): boolean {
    const diffs = diffSessions(sessionOriginalRef.value, sessionEditingRef.value);
    return diffs.length > 0;
  }

  deleteSession(): void {
    if (confirm('Are you sure you want to delete the session named [' + this.$route.params.sess + ']?')) {
      if (socketRef.value) {
        socketRef.value.send({channel: sessionService.key, cmd: clientCommands.deleteSession, param: this.$route.params.sess});
      }
    }
  }

  editSession(): void {
    this.showEditor = this.$route.params.sess;
  }

  cancelEdit(): void {
    this.showEditor = "";
  }

  reset(): void {
    // TODO? this.cancelEdit();
    sessionEditingRef.value = jsonClone(sessionOriginalRef.value);
  }

  save(): void {
    const s = socketRef.value;
    if (!s) {
      return;
    }
    const sess = sessionEditingRef.value;
    if (sess) {
      const param = {orig: sessionOriginalRef.value?.key || sess.key, sess: sess};
      s.send({channel: sessionService.key, cmd: clientCommands.saveSession, param});
    }
  }

  mounted(): void {
    let title = this.$route.params.sess;
    if (title === "_") {
      title = "default";
    }
    setBC(this, {path: "/s", title: "sessions"}, {path: "", title});
  }
}
</script>
