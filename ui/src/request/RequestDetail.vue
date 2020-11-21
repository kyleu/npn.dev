<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right"><router-link :to="'/c/' + this.$route.params.coll"><Icon icon="close" /></router-link></div>
        <h3 class="uk-card-title">
          <Icon class="nav-icon-h3" icon="link" />
          <span>{{ req ? (req.title || req.key) : $route.params.req }}</span>
        </h3>
        <div v-if="req">
          <URLEditor :req="req" />
          <div v-if="different" class="right">
            <button v-style-button class="uk-button uk-button-default uk-margin-small-right mt" @click="reset();">Reset</button>
            <button v-style-button class="uk-button uk-button-default mt" @click="save();">Save Changes</button>
          </div>
          <button v-style-button class="uk-button uk-button-default uk-margin-small-right mt" @click="doCall()">Call</button>
          <ExportActions />
          <button v-style-button class="uk-button uk-button-default uk-margin-small-right mt" @click="deleteRequest()">Delete</button>
        </div>
      </div>
      <router-view />
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {cloneRequest, NPNRequest} from "@/request/model";
import RequestEditor from "@/request/editor/RequestEditor.vue";
import URLEditor from "@/request/editor/URLEditor.vue";
import {diff} from "@/request/prototype/diff";
import ExportActions from "@/request/editor/ExportActions.vue";
import {callResultRef, requestEditingRef, requestOriginalRef, setActiveRequest} from "@/request/state";
import {socketRef} from "@/socket/socket";
import {requestService} from "@/util/services";
import {clientCommands} from "@/util/command";
import Icon from "@/util/Icon.vue";

@Component({ components: {Icon, ExportActions, RequestEditor, URLEditor } })
export default class RequestDetail extends Vue {
  get req(): NPNRequest | undefined {
    setActiveRequest(this.$route.params.coll, this.$route.params.req);
    return requestEditingRef.value;
  }

  get different(): boolean {
    const diffs = diff(requestOriginalRef.value, requestEditingRef.value);
    return diffs.length > 0;
  }

  reset(): void {
    requestEditingRef.value = cloneRequest(requestOriginalRef.value)
  }

  save(): void {
    const s = socketRef.value;
    if (!s) {
      return;
    }
    const e = requestEditingRef.value;
    if (e) {
      const param = {coll: this.$route.params.coll, orig: requestOriginalRef.value?.key || e.key, req: e}
      s.send({svc: requestService.key, cmd: clientCommands.saveRequest, param})
    }
  }

  doCall(): void {
    if (this.$route.name === 'CallResult') {
      callResultRef.value = undefined;
    } else {
      this.$router.push({name: "CallResult", params: {coll: this.$route.params.coll, req: this.$route.params.req}})
    }
  }

  deleteRequest(): void {
    if (confirm('Are you sure you want to delete request [' + this.$route.params.req + ']?')) {
      if (socketRef.value) {
        const param = { coll: this.$route.params.coll, req: this.$route.params.req};
        socketRef.value.send({svc: requestService.key, cmd: clientCommands.deleteRequest, param});
      }
    }
  }
}
</script>
