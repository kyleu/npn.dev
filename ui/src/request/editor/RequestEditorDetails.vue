<template>
  <div>
    <div>
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
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {NPNRequest} from "@/request/model";
import {requestEditingRef} from "@/request/state";

@Component
export default class RequestEditorDetails extends Vue {
  get req(): NPNRequest | undefined {
    return requestEditingRef.value;
  }
}
</script>
