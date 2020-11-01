<template>
  <div v-if="this.$parent.req" class="request-editor uk-card uk-card-body uk-card-default mt">
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
                <input v-model="this.$parent.req.key" class="uk-input" name="key" type="text" data-lpignore="true" />
              </label>
            </div>

            <div class="mt">
              <label class="uk-form-label">
                Title
                <input v-model="this.$parent.req.title" class="uk-input" name="title" type="text" data-lpignore="true" />
              </label>
            </div>

            <div class="mt">
              <label class="uk-form-label">
                Description
                <textarea v-model="this.$parent.req.description" class="uk-textarea" name="description" data-lpignore="true"></textarea>
              </label>
            </div>
          </li>

          <li><QueryParamsEditor :qp="this.$parent.req.prototype.query" /></li>
          <li><AuthEditor :auth="this.$parent.req.prototype.auth" /></li>
          <li><HeadersEditor :headers="this.$parent.req.prototype.headers" /></li>
          <li><BodyEditor :body="this.$parent.req.prototype.body" /></li>
          <li><OptionsEditor :opts="this.$parent.req.prototype.options" /></li>
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
import {getStateSetBC} from "@/util/vutils";

@Component({ components: { AuthEditor, BodyEditor, HeadersEditor, OptionsEditor, QueryParamsEditor } })
export default class RequestEditor extends Vue {
  created(): void {
    getStateSetBC(this, {path: "/c/" + this.$route.params.coll, title: this.$route.params.coll}, {path: "", title: this.$route.params.req});
  }

  updated(): void {
    getStateSetBC(this, {path: "/c/" + this.$route.params.coll, title: this.$route.params.coll}, {path: "", title: this.$route.params.req});
  }
}
</script>
