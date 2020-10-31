<template>
  <div v-if="req" class="uk-card uk-card-body uk-card-default mt">
    CALL!
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {NPNRequest} from "@/request/model";

@Component
export default class RequestCall extends Vue {
  get req(): NPNRequest | undefined {
    return this.$parent.req;
  }
}
</script>
