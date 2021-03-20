<template>
  <div v-if="req" class="request-editor uk-card uk-card-body uk-card-default mt">
    <form action="" method="post" onsubmit="return false;">
      <div>
        <ul data-uk-tab="">
          <li><a id="tab-request-details" href="#details" @click="setTab('details')">Details</a></li>
          <li><a id="tab-request-query" href="#query" @click="setTab('query')">Query</a></li>
          <li><a id="tab-request-auth" href="#auth" @click="setTab('auth')">Auth</a></li>
          <li><a id="tab-request-headers" href="#headers" @click="setTab('headers')">Headers</a></li>
          <li><a id="tab-request-body" href="#body" @click="setTab('body')">Body</a></li>
          <li><a id="tab-request-options" href="#options" @click="setTab('options')">Options</a></li>
        </ul>
        <ul class="uk-switcher">
          <li><RequestEditorDetails /></li>
          <li><QueryParamsEditor /></li>
          <li><AuthEditor /></li>
          <li><HeadersEditor /></li>
          <li><BodyEditor ref="body" /></li>
          <li><OptionsEditor /></li>
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
  activeTab = "";

  get req(): NPNRequest | undefined {
    return requestEditingRef.value;
  }

  mounted(): void {
    const title = this.$route.params.coll === "_" ? "default" : this.$route.params.coll;
    setBC(this, {path: "/c/" + this.$route.params.coll, title}, {path: "", title: this.$route.params.req});
  }

  setTab(s: string): void {
    this.activeTab = s;
    if (s === "body") {
      (this.$refs["body"] as BodyEditor).refresh();
    }
  }
}
</script>
