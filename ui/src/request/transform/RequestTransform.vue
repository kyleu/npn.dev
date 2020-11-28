<template>
  <div class="uk-card uk-card-body uk-card-default mt">
    <div class="right"><router-link :to="'/c/' + this.$route.params.coll + '/' + this.$route.params.req"><Icon icon="close" /></router-link></div>
    <h3 v-if="result" class="uk-card-title">{{ result.req }}: {{ result.fmt }}</h3>
    <h3 v-else class="uk-card-title">Loading...</h3>
    <div v-if="result" class="mt">
      <pre>{{ result.out }}</pre>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {setBCReq} from "@/util/vutils";
import {TransformResult} from "@/request/transform/transformResult";
import Icon from "@/util/Icon.vue";
import {getTransformResult} from "@/request/transform/state";

@Component({ components: {Icon} })
export default class RequestTransform extends Vue {
  get result(): TransformResult | undefined {
    return getTransformResult(this.$route.params.coll, this.$route.params.req, this.$route.params.tx);
  }

  updated(): void {
    setBCReq(this, this.$route.params.tx);
  }
}
</script>
