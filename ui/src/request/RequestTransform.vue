<template>
  <div v-if="req" class="uk-card uk-card-body uk-card-default mt">
    TRANSFORM: {{ this.$route.params.tx }}!
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {NPNRequest} from "@/request/model";
import {getStateSetBCReq} from "@/util/vutils";
import RequestDetail from "@/request/RequestDetail.vue";

@Component
export default class RequestTransform extends Vue {
  get req(): NPNRequest | undefined {
    return (this.$parent as RequestDetail).req;
  }

  created(): void {
    getStateSetBCReq(this, this.$route.params.tx);
  }

  updated(): void {
    getStateSetBCReq(this, this.$route.params.tx);
  }
}
</script>
