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

          <li><RequestEditorQueryParams :qp="req.prototype.query" /></li>
          <li><RequestEditorAuth :auth="req.prototype.auth" /></li>
          <li><RequestEditorHeaders :headers="req.prototype.headers" /></li>
          <li><RequestEditorBody :body="req.prototype.body" /></li>
          <li><RequestEditorOptions :opts="req.prototype.options" /></li>
        </ul>
      </div>
    </form>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {NPNRequest} from "@/request/model";
import RequestEditorAuth from "@/request/RequestEditorAuth.vue";
import RequestEditorBody from "@/request/RequestEditorBody.vue";
import RequestEditorHeaders from "@/request/RequestEditorHeaders.vue";
import RequestEditorOptions from "@/request/RequestEditorOptions.vue";
import RequestEditorQueryParams from "@/request/RequestEditorQueryParams.vue";

@Component({ components: { RequestEditorAuth, RequestEditorBody, RequestEditorHeaders, RequestEditorOptions, RequestEditorQueryParams } })
export default class RequestEditor extends Vue {
  @Prop() req!: NPNRequest;
}
</script>
