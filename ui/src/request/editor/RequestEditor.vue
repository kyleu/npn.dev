<template>
  <div v-if="req" class="request-editor uk-card uk-card-body uk-card-default mt">
    <form action="" method="post" onsubmit="console.log('XXXXXXX');return false;">
      <div>
        <ul data-uk-tab="">
          <li><a href="#details">Details</a></li>
          <li><a href="#query">Query</a></li>
          <li><a href="#auth">Auth</a></li>
          <li><a href="#headers">Headers</a></li>
          <li><a href="#body">Body</a></li>
          <li><a href="#options">Options</a></li>
        </ul>
        <ul class="uk-switcher uk-margin">
          <li class="request-details-panel">
            <div class="mt">
              <label class="uk-form-label">
                Key
                <input v-model="req.key" class="uk-input" name="key" type="text" data-lpignore="true" />
              </label>
            </div>

            <div class="mt">
              <label class="uk-form-label">
                Title
                <input v-model="req.title" class="uk-input" name="title" type="text" data-lpignore="true" />
              </label>
            </div>

            <div class="mt">
              <label class="uk-form-label">
                Description
                <textarea v-model="req.description" class="uk-textarea" name="description" data-lpignore="true"></textarea>
              </label>
            </div>
          </li>

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
import {requestEditingRef} from "@/state/state";
import {NPNRequest} from "@/request/model";

@Component({ components: { AuthEditor, BodyEditor, HeadersEditor, OptionsEditor, QueryParamsEditor } })
export default class RequestEditor extends Vue {
  get req(): NPNRequest | undefined {
    return requestEditingRef.value;
  }

  created(): void {
    setBC(this, {path: "/c/" + this.$route.params.coll, title: this.$route.params.coll}, {path: "", title: this.$route.params.req});
  }

  updated(): void {
    setBC(this, {path: "/c/" + this.$route.params.coll, title: this.$route.params.coll}, {path: "", title: this.$route.params.req});
  }
}
</script>
