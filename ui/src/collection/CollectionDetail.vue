<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right"><router-link :class="'uk-icon ' + profile.linkColor + '-fg'" data-uk-icon="close" to="/c"></router-link></div>
        <h3 class="uk-card-title">
          <span class="nav-icon-h3 uk-icon" data-uk-icon="icon: album"></span>
          {{ coll ? coll.title : $route.params.coll }}
        </h3>
        <p v-if="coll">{{ coll.description }}</p>

        <button class="uk-button uk-button-default uk-margin-small-right mt" @click="editCollection()">Edit</button>
        <button class="uk-button uk-button-default mt" @click="deleteCollection()">Delete</button>
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
import {Profile, profileRef} from "@/user/profile";
import {getCollection, getCollectionRequestSummaries} from "@/collection/state";
import {socketRef} from "@/socket/socket";
import {collectionService} from "@/util/services";
import {clientCommands} from "@/util/command";

@Component({ components: { RequestSummaryList } })
export default class CollectionDetail extends Vue {
  get profile(): Profile | undefined {
    return profileRef.value;
  }

  get coll(): Collection | undefined {
    return getCollection(this.$route.params.coll);
  }

  get requests(): Summary[] | undefined {
    return getCollectionRequestSummaries(this.$route.params.coll);
  }

  mounted(): void {
    setBC(this, {path: "", title: this.$route.params.coll});
  }

  editCollection(): void {
    console.log("TODO!");
  }

  deleteCollection(): void {
    if (confirm('Are you sure you want to delete the collection named [' + this.$route.params.coll + ']?')) {
      if (socketRef.value) {
        socketRef.value.send({svc: collectionService.key, cmd: clientCommands.deleteCollection, param: this.$route.params.coll});
      }
    }
  }
}
</script>
