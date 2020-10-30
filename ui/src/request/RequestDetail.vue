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
          <div class="mt uk-panel">
            <div class="left" style="width:120px;">
              <select :id="req.key + '-method'" class="uk-select" name="method">
                <option selected="selected">TODO</option>;
              </select>
            </div>
            <div :id="req.key + '-link'" class="url-view uk-inline right" style="width:calc(100% - 120px);">
              <a class="uk-form-icon uk-form-icon-flip" href="" onclick="TODO(); call" title="send request" data-uk-icon="icon: play" />
              <div onclick="TODO();request.editor.toggleURLEditor('{req.key}', true);">
                <span :id="req.key + '-urlview'" class="url-link">TODO prototypeToHTML(req.prototype)</span>
              </div>
            </div>
            <div :id="req.key + '-edit'" class="url-input hidden uk-inline right" style="width:calc(100% - 120px);">
              <a class="uk-form-icon uk-form-icon-flip" href="" onclick="TODO();request.editor.toggleURLEditor('{req.key}', false);return false;" title="cancel edit" data-uk-icon="icon: close" />
              <form onsubmit="TODO();return false;">
                <input :id="req.key + '-url'" class="uk-input" name="url" type="text" :value="req.prototype.path" data-lpignore="true" />
              </form>
            </div>
          </div>
          <div id="save-panel" class="right hidden">
            <button class="uk-button uk-button-default uk-margin-small-right mt" onclick="TODO();request.form.reset('coll', 'req.key');">Reset</button>
            <button class="uk-button uk-button-default mt" onclick="TODO();request.form.saveCurrentRequest('coll', 'req.key');">Save Changes</button>
          </div>
          <router-link class="uk-button uk-button-default uk-margin-small-right mt" :to="'/c/' + this.$route.params.coll + '/' + this.$route.params.req + '/call'">Call</router-link>
          <div class="uk-inline">
            <button type="button" class="uk-button uk-button-default uk-margin-small-right mt">Export</button>
            <div id="export-dropdown" data-uk-dropdown="mode: click">
              <ul class="uk-list uk-list-divider" style="margin-bottom: 0;">
                <li>
                  {nav.link({path: path + "/transform/" + k, title: transforms[k], onclk: "UIkit.dropdown(dom.req('#export-dropdown')).hide(false);"})}
                </li>
              </ul>
            </div>
          </div>
          <router-link class="uk-button uk-button-default uk-margin-small-right mt" :to="'/c/' + this.$route.params.coll + '/' + this.$route.params.req + '/delete'">Delete</router-link>
        </div>
      </div>
      <RequestEditor :req="req" />
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {getState, getStateSetBC} from "@/util/vutils";
import RequestSummaryList from "@/request/RequestSummaryList.vue";
import {NPNRequest} from "@/request/model";
import RequestEditor from "@/request/RequestEditor.vue";

@Component({ components: {RequestEditor, RequestSummaryList } })
export default class RequestDetail extends Vue {
  get req(): NPNRequest | undefined {
    const ret = getState(this).getRequestDetail(this.$route.params.coll, this.$route.params.req);
    if ((!ret) && this.$route.params.req) {
      this.$store.commit("send", {svc: "request", cmd: "getRequest", param: {coll: this.$route.params.coll, req: this.$route.params.req}});
    }
    return ret
  }

  created(): void {
    getStateSetBC(this, {path: "/c/" + this.$route.params.coll, title: this.$route.params.coll}, {path: "", title: this.$route.params.req});
  }

  updated(): void {
    getStateSetBC(this, {path: "/c/" + this.$route.params.coll, title: this.$route.params.coll}, {path: "", title: this.$route.params.req});
  }
}
</script>
