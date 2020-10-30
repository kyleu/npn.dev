<template>
  <div class="uk-card uk-card-body uk-card-default mt">
    <h3 class="uk-card-title">Requests</h3>
    <div id="request-list" class="mt">
      <ul class="uk-list uk-list-divider">
        <li v-for="req of requests" :key="req.key">
          <RequestSummaryListItem :coll="coll" :req="req" />
        </li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import {Summary} from "@/request/model";
import RequestSummaryListItem from "@/request/RequestSummaryListItem.vue";

@Component({
  components: {RequestSummaryListItem}
})
export default class RequestSummaryList extends Vue {
  @Prop() coll!: string;
  @Prop() requests!: Summary[];
}
</script>
