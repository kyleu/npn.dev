<template>
  <div class="uk-card uk-card-body uk-card-default mt">
    <h3 class="uk-card-title">Requests</h3>

    <div class="mt">
      <div class="uk-inline" style="width: 100%;">
        <a class="uk-form-icon uk-form-icon-flip" title="add request" href="" @click.prevent="addRequest()"><Icon icon="plus" /></a>
        <form @submit.prevent="addRequest()">
          <input id="req-add-input" class="uk-input" type="text" placeholder="Add URL" data-lpignore="true" />
        </form>
      </div>
    </div>

    <div id="request-list" class="mt">
      <ul class="uk-list uk-list-divider">
        <li v-for="req of requests" :key="req.key">
          <RequestSummaryListItem :coll="coll" :req="req" />
        </li>
        <li v-if="(!requests) || requests.length === 0"><em>no requests</em></li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import {Summary} from "@/request/model";
import RequestSummaryListItem from "@/request/RequestSummaryListItem.vue";
import {socketRef} from "@/socket/socket";
import {collectionService} from "@/util/services";
import {clientCommands} from "@/util/command";
import Icon from "@/util/Icon.vue";

@Component({ components: {Icon, RequestSummaryListItem} })
export default class RequestSummaryList extends Vue {
  @Prop() coll!: string;
  @Prop() requests!: Summary[];

  addRequest(): void {
    const el = document.getElementById("req-add-input") as HTMLInputElement;
    const url = el.value.trim();
    el.value = "";
    if (socketRef.value) {
      socketRef.value.send({channel: collectionService.key, cmd: clientCommands.addRequestURL, param: {coll: this.$route.params.coll, url}});
    }
  }
}
</script>
