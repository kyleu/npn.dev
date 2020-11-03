<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right"><router-link :class="'uk-icon ' + profile.linkColor + '-fg'" data-uk-icon="close" :to="'/c/' + this.$route.params.coll"></router-link></div>
        <h3 class="uk-card-title">
          <span class="nav-icon-h3 uk-icon" data-uk-icon="icon: link"></span>
          <span>{{ req ? (req.title || req.key) : $route.params.req }}</span>
        </h3>
        <div v-if="req">
          <URLEditor :req="req" />
          <div v-if="different" class="right">
            <button class="uk-button uk-button-default uk-margin-small-right mt" @click="reset();">Reset</button>
            <button class="uk-button uk-button-default mt" @click="save();">Save Changes</button>
          </div>
          <button class="uk-button uk-button-default uk-margin-small-right mt" @click="doCall()">Call</button>
          <ExportActions />
          <router-link class="uk-button uk-button-default uk-margin-small-right mt" :to="'/c/' + this.$route.params.coll + '/' + req.key + '/delete'">Delete</router-link>
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
import {diff} from "@/request/diff";
import ExportActions from "@/request/editor/ExportActions.vue";
import {setActiveRequest, requestEditingRef, requestOriginalRef} from "@/request/state";
import {Profile, profileRef} from "@/user/profile";
import { callResultRef } from '@/request/state'
import {socketRef} from "@/socket/socket";
import {requestService} from "@/util/services";
import {clientCommands} from "@/util/command";

@Component({ components: {ExportActions, RequestEditor, URLEditor } })
export default class RequestDetail extends Vue {
  get profile(): Profile | undefined {
    return profileRef.value;
  }

  get req(): NPNRequest | undefined {
    setActiveRequest(this.$route.params.coll, this.$route.params.req);
    return requestEditingRef.value;
  }

  doCall(): void {
    if (this.$route.name === 'CallResult') {
      callResultRef.value = undefined;
    } else {
      this.$router.push({name: "CallResult", params: {coll: this.$route.params.coll, req: this.$route.params.req}})
    }
  }

  get different(): boolean {
    const diffs = diff(requestOriginalRef.value, requestEditingRef.value);
    // console.debug(jsonParse(jsonStr(requestOriginalRef.value)), jsonParse(jsonStr(requestEditingRef.value)));
    if (diffs.length > 0) {
      console.debug(diffs);
    }
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
}
</script>
