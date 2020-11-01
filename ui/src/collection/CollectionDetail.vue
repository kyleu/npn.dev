<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right"><router-link :class="'uk-icon ' + profile.linkColor + '-fg'" data-uk-icon="close" to="/c"></router-link></div>
        <h3 class="uk-card-title">
          <span v-if="coll">Collection [{{ coll.title }}]</span>
          <span v-else>Collection [{{ $route.params.coll }}]</span>
        </h3>
        <p v-if="coll">{{ coll.description }}</p>
      </div>
      <RequestSummaryList :coll="$route.params.coll" :requests="requests" />
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {Collection} from "@/collection/collection";
import {setBC} from "@/util/vutils";
import {Summary} from "@/request/model";
import RequestSummaryList from "@/request/RequestSummaryList.vue";
import Profile from "@/user/profile";
import {getCollection, getCollectionRequestSummaries, profileRef} from "@/state/state";

@Component({ components: { RequestSummaryList } })
export default class CollectionDetail extends Vue {
  get profile(): Profile | undefined {
    return profileRef.value;
  }

  get coll(): Collection | undefined {
    return getCollection(this.$route.params.coll);
  }

  get requests(): Summary[] | undefined {
    const coll = this.$route.params.coll;
    if (coll) {
      const reqs = getCollectionRequestSummaries(this.$route.params.coll);
      if (!reqs) {
        this.$store.commit("send", {svc: "collection", cmd: "getCollection", param: this.$route.params.coll});
      }
      return reqs
    }
    return undefined;
  }

  created(): void {
    setBC(this, {path: "", title: this.$route.params.coll});
  }

  updated(): void {
    setBC(this, {path: "", title: this.$route.params.coll});
  }
}
</script>
