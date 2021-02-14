<template>
  <div class="nav-list">
    <div v-if="(!requests) || (requests.length === 0)" class="nav-link">No requests</div>
    <RequestListItem v-for="r in requests" :key="r.key" :req="r" />
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import RequestListItem from "@/request/RequestListItem.vue";
import {collectionSummariesRef} from "@/collection/state";
import {Summary} from "@/request/model";

@Component({ components: { RequestListItem } })
export default class RequestList extends Vue {
  get requests(): Summary[] | undefined {
    return collectionSummariesRef.value.find(x => x.key === "_")?.requests;
  }
}
</script>
