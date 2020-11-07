<template>
  <div v-if="req" class="request-editor uk-card uk-card-body uk-card-default mt">
    <form action="" method="post" onsubmit="console.log('XXXXXXX');return false;">
      <div>
        <ul data-uk-tab="">
          <li><a href="#details" @click="setTab('details')">Details</a></li>
          <li><a href="#query" @click="setTab('query')">Query</a></li>
          <li><a href="#auth" @click="setTab('auth')">Auth</a></li>
          <li><a href="#headers" @click="setTab('headers')">Headers</a></li>
          <li><a href="#body" @click="setTab('body')">Body</a></li>
          <li><a href="#options" @click="setTab('options')">Options</a></li>
        </ul>
        <ul class="uk-switcher uk-margin">
          <li><RequestEditorDetails :req="req" /></li>
          <li><QueryParamsEditor :qp="req.prototype.query" /></li>
          <li><AuthEditor :auth="req.prototype.auth" /></li>
          <li><HeadersEditor :headers="req.prototype.headers" /></li>
          <li><BodyEditor :body="req.prototype.body" /></li>
          <li><OptionsEditor :opts="req.prototype.options" /></li>
        </ul>
      </div>
    </form>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import AuthEditor from "@/auth/AuthEditor.vue";
import BodyEditor from "@/body/BodyEditor.vue";
import HeadersEditor from "@/header/HeadersEditor.vue";
import OptionsEditor from "@/request/editor/OptionsEditor.vue";
import QueryParamsEditor from "@/request/editor/QueryParamsEditor.vue";
import {setBC} from "@/util/vutils";
import {requestEditingRef} from "@/request/state";
import {NPNRequest} from "@/request/model";
import RequestEditorDetails from "@/request/editor/RequestEditorDetails.vue";

@Component({ components: {RequestEditorDetails, AuthEditor, BodyEditor, HeadersEditor, OptionsEditor, QueryParamsEditor } })
export default class RequestEditor extends Vue {
  activeTab = ""

  get req(): NPNRequest | undefined {
    return requestEditingRef.value;
  }

  mounted(): void {
    setBC(this, {path: "/c/" + this.$route.params.coll, title: this.$route.params.coll}, {path: "", title: this.$route.params.req});
  }

  setTab(s: string): void {
    this.activeTab = s;
    console.log("setEditorTab: " + s);
  }
}
</script>
