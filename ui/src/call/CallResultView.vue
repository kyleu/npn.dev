<template>
  <div class="uk-card uk-card-body uk-card-default mt">
    <div class="right">
      <router-link :to="'/c/' + this.$route.params.coll + '/' + this.$route.params.req">(return to editor) <Icon icon="close" /></router-link>
    </div>
    <div v-for="(cycle, idx) in cycles" :key="idx">
      <div style="clear:both"></div>
      <hr v-if="idx > 0" />
      <ResponsePanel :cycle="cycle" />
    </div>
    <h3 v-if="cycles.length === 0" class="uk-card-title">Loading...</h3>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {setBCReq} from "@/util/vutils";
import {RequestResultCycle, RequestResults} from "@/call/model";
import {Prototype} from "@/request/model";
import ResponsePanel from "@/call/ResponsePanel.vue";
import Icon from "@/util/Icon.vue";
import {getRequestResults} from "@/call/state";

interface CallParam {
  coll: string;
  req: string;
  proto: Prototype;
}

@Component({ components: {Icon, ResponsePanel } })
export default class CallResultView extends Vue {
  private pending: CallParam | undefined;

  get result(): RequestResults | undefined {
    return getRequestResults(this.$route.params.coll, this.$route.params.req);
  }

  get cycles(): RequestResultCycle[] {
    return this.result?.cycles || [];
  }

  mounted(): void {
    setBCReq(this, "call");
  }
}
</script>
