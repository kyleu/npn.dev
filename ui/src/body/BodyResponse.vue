<template>
  <div class="uk-overflow-auto">
    <div v-if="body">
      <HTMLBody v-if="body.type === 'html'" ref="html" :url="url" :config="body.config" />
      <XMLBody v-else-if="body.type === 'xml'" ref="xml" :config="body.config" />
      <JSONBody v-else-if="body.type === 'json'" ref="json" :config="body.config" />
      <ImageBody v-else-if="body.type === 'image'" :config="body.config" />
      <RawBody v-else-if="body.type === 'raw'" :config="body.config" />
      <ErrorBody v-else-if="body.type === 'error'" :config="body.config" />
      <div v-else>
        <h4>{{ body.type }}</h4>
        <em>currently unhandled</em>
      </div>
    </div>
    <div v-else>
      <p><em>no response body</em></p>
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
import XMLBody from "@/body/XMLBody.vue";

@Component({ components: {XMLBody, HTMLBody, ImageBody, JSONBody, RawBody, ErrorBody } })
export default class BodyResponse extends Vue {
  @Prop() url!: string;
  @Prop() body!: RBody;

  refresh(): void {
    this.$nextTick(function() {
      if (this.$refs["html"]) {
        (this.$refs["html"] as HTMLBody).refresh();
      }
      if (this.$refs["xml"]) {
        (this.$refs["xml"] as XMLBody).refresh();
      }
      if (this.$refs["json"]) {
        (this.$refs["json"] as JSONBody).refresh();
      }
    });
  }
}
</script>
