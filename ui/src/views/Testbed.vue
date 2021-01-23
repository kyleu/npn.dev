<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right">
          <router-link to="/"><Icon icon="close" /></router-link>
        </div>
        <h3 class="uk-card-title">Testbed</h3>
        <form class="uk-form-stacked" onsubmit="return false;">
          <fieldset class="uk-fieldset">
            <legend class="hidden">connection form</legend>
            <div class="uk-margin-small">
              <label class="uk-form-label" for="testbed-level">Log Level</label>
              <select id="testbed-level" v-model="level" class="uk-select" name="level">
                <option value="debug">Debug</option>
                <option value="info">Info</option>
                <option value="warn">Warn</option>
                <option value="error">Error</option>
              </select>
            </div>
            <div class="uk-margin-small">
              <label class="uk-form-label" for="testbed-input">Message</label>
              <textarea id="testbed-input" v-model="message" class="uk-textarea" name="msg" type="text" />
            </div>
            <div class="mt">
              <button class="uk-button uk-button-default" type="submit" @click="send()">Send</button>
            </div>
          </fieldset>
        </form>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { setBC } from "@/util/vutils";
import Icon from "@/util/Icon.vue";
import {socketRef} from "@/socket/socket";
import {systemService} from "@/util/services";
import {clientCommands} from "@/util/command";

@Component({ components: { Icon } })
export default class Testbed extends Vue {
  level = "info";
  message = "{}";

  send(): void {
    if (socketRef.value) {
      const p = {"t": "log", "k": this.level, "v": this.message};
      socketRef.value.send({channel: systemService.key, cmd: clientCommands.testbed, param: p});
    }
  }

  mounted(): void {
    setBC(this, { path: "", title: "testbed" });
  }
}
</script>
