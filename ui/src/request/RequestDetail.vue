<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right"><router-link :class="'uk-icon ' + $store.state.profile.linkColor + '-fg'" data-uk-icon="close" :to="'/c/' + this.$route.params.coll"></router-link></div>
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
          <router-link class="uk-button uk-button-default uk-margin-small-right mt" :to="'/c/' + this.$route.params.coll + '/' + req.key + '/call'">Call</router-link>
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
import {getState} from "@/util/vutils";
import RequestSummaryList from "@/request/RequestSummaryList.vue";
import {NPNRequest} from "@/request/model";
import RequestEditor from "@/request/editor/RequestEditor.vue";
import URLEditor from "@/request/editor/URLEditor.vue";
import {diff} from "@/request/diff";
import ExportActions from "@/request/editor/ExportActions.vue";

@Component({ components: {ExportActions, RequestEditor, RequestSummaryList, URLEditor } })
export default class RequestDetail extends Vue {
  get req(): NPNRequest | undefined {
    this.$store.commit("setActiveRequest", {coll: this.$route.params.coll, req: this.$route.params.req});
    const s = getState(this);
    if ((!s.requestEditing) && this.$route.params.req) {
      this.$store.commit("send", {svc: "request", cmd: "getRequest", param: {coll: this.$route.params.coll, req: this.$route.params.req}});
    }
    return s.requestEditing;
  }

  get different(): boolean {
    const s = getState(this);
    const diffs = diff(s.requestOriginal, s.requestEditing);
    console.debug(s.requestOriginal, s.requestEditing);
    if (diffs.length > 0) {
      console.debug(diffs);
    }
    return diffs.length > 0;
  }
}
</script>
