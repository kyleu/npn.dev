<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right"><router-link to="/c"><Icon icon="close" /></router-link></div>
        <h3 class="uk-card-title">
          <Icon class="nav-icon-h3" icon="album" />
          {{ (coll && coll.title.length > 0) ? coll.title : $route.params.coll }}
        </h3>
        <p v-if="coll && coll.description && showEditor.length === 0">{{ coll.description }}</p>
        <div v-if="showEditor.length !== 0" class="mt">
          <div v-if="$route.params.coll !== '_'">
            <label class="uk-form-label">
              Key
              <input v-model="collEdit.key" class="uk-input" name="key" type="text" data-lpignore="true" />
            </label>
          </div>
          <input v-else v-model="collEdit.key" name="key" type="hidden" data-lpignore="true" />

          <div class="mt">
            <label class="uk-form-label">
              Title
              <input v-model="collEdit.title" class="uk-input" name="title" type="text" data-lpignore="true" />
            </label>
          </div>

          <div class="mt">
            <label class="uk-form-label">
              Description
              <textarea v-model="collEdit.description" class="uk-textarea" name="description" data-lpignore="true"></textarea>
            </label>
          </div>

          <button v-style-button class="right uk-button uk-button-default mt" @click="saveCollection()">Save Changes</button>
          <button v-style-button class="right uk-button uk-button-default mrs mt" @click="showEditor = ''">Cancel</button>
          <button v-if="$route.params.coll !== '_'" v-style-button class="uk-button uk-button-default mrs mt" @click="deleteCollection()">Delete</button>
        </div>
        <div v-else>
          <button v-style-button class="uk-button uk-button-default mrs mt" @click="editCollection()">Edit</button>
          <CollectionTransformActions />
        </div>
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
import {getCollection, getCollectionRequestSummaries} from "@/collection/state";
import {socketRef} from "@/socket/socket";
import {collectionService} from "@/util/services";
import {clientCommands} from "@/util/command";
import {jsonClone} from "@/util/json";
import Icon from "@/util/Icon.vue";
import CollectionTransformActions from "@/transform/CollectionTransformActions.vue";

@Component({ components: {CollectionTransformActions, Icon, RequestSummaryList } })
export default class CollectionDetail extends Vue {
  showEditor = "";

  get coll(): Collection | undefined {
    return getCollection(this.$route.params.coll);
  }

  get collEdit(): Collection | undefined {
    return jsonClone(getCollection(this.$route.params.coll));
  }

  get requests(): Summary[] | undefined {
    return getCollectionRequestSummaries(this.$route.params.coll);
  }

  editCollection(): void {
    this.showEditor = this.$route.params.coll;
  }

  deleteCollection(): void {
    if (confirm('Are you sure you want to delete the collection named [' + this.$route.params.coll + ']?')) {
      if (socketRef.value) {
        socketRef.value.send({channel: collectionService.key, cmd: clientCommands.deleteCollection, param: this.$route.params.coll});
      }
    }
  }

  saveCollection(): void {
    if (socketRef.value) {
      socketRef.value.send({
        channel: collectionService.key,
        cmd: clientCommands.saveCollection,
        param: { originalKey: this.$route.params.coll, coll: this.collEdit }
      });
      this.showEditor = "";
    }
  }

  mounted(): void {
    let title = this.$route.params.coll;
    if (title === "_") {
      title = "default";
    }
    setBC(this, {path: "", title});
    if ((this.showEditor.length > 0) && (this.showEditor !== this.$route.params.coll)) {
      this.showEditor = "";
    }
  }
}
</script>
