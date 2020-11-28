<template>
  <div class="uk-overflow-auto">
    <div v-if="body">
      <HTMLBody v-if="body.type === 'html'" :url="url" :config="body.config" />
      <JSONBody v-else-if="body.type === 'json'" :config="body.config" />
      <ImageBody v-else-if="body.type === 'image'" :config="body.config" />
      <RawBody v-else-if="body.type === 'raw'" :config="body.config" />
      <ErrorBody v-else-if="body.type === 'error'" :config="body.config" />
      <div v-else>
        <em>{{ body.type }}</em>
        <h4>TODO</h4>
      </div>
    </div>
    <div v-else>
      <div>no body</div>
      <p><em>{{ url }}!</em></p>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {RBody} from "@/body/model";
import HTMLBody from "@/body/HTMLBody.vue";
import ImageBody from "@/body/ImageBody.vue";
import JSONBody from "@/body/JSONBody.vue";
import RawBody from "@/body/RawBody.vue";
import ErrorBody from "@/body/ErrorBody.vue";

@Component({ components: { HTMLBody, ImageBody, JSONBody, RawBody, ErrorBody } })
export default class BodyResponse extends Vue {
  @Prop() url!: string;
  @Prop() body!: RBody;
}
</script>
