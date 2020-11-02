<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right"><router-link :class="'uk-icon ' + profile.linkColor + '-fg'" data-uk-icon="close" :to="'/c/' + this.$route.params.coll"></router-link></div>
        <h3 class="uk-card-title">
          <span v-if="req">{{ req.title || req.key }}</span>
          <span v-else>{{ $route.params.req }}</span>
        </h3>
        <div v-if="req">
          <URLEditor :req="req" />
          <div v-if="different" class="right">
            <button class="uk-button uk-button-default uk-margin-small-right mt" onclick="TODO();">Reset</button>
            <button class="uk-button uk-button-default mt" onclick="TODO();">Save Changes</button>
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
import RequestSummaryList from "@/request/RequestSummaryList.vue";
import {NPNRequest} from "@/request/model";
import RequestEditor from "@/request/editor/RequestEditor.vue";
import URLEditor from "@/request/editor/URLEditor.vue";
import {diff} from "@/request/diff";
import ExportActions from "@/request/editor/ExportActions.vue";
import {setActiveRequest, requestEditingRef, requestOriginalRef} from "@/request/state";
import {Profile, profileRef} from "@/user/profile";
import { callResultRef } from '@/request/state'

@Component({ components: {ExportActions, RequestEditor, RequestSummaryList, URLEditor } })
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
}
</script>
