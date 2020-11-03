<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right"><router-link :class="'uk-icon ' + profile.linkColor + '-fg'" data-uk-icon="close" to="/"></router-link></div>
        <h3 class="uk-card-title">Collections!</h3>

        <div class="mt">
          <div class="uk-inline" style="width: 100%;">
            <a class="uk-form-icon uk-form-icon-flip" title="cancel edit" data-uk-icon="icon: plus" href="" @click.prevent="addCollection()" />
            <form @submit.prevent="addCollection()">
              <input id="coll-add-input" class="uk-input" type="text" placeholder="Add collection" data-lpignore="true" />
            </form>
          </div>
        </div>

        <CollectionGallery />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import CollectionGallery from "@/collection/CollectionGallery.vue";
import {setBC} from "@/util/vutils";
import {Profile, profileRef} from "@/user/profile";
import {socketRef} from "@/socket/socket";
import {collectionService} from "@/util/services";
import {clientCommands} from "@/util/command";

@Component({components: {CollectionGallery}})
export default class CollectionIndex extends Vue {
  get profile(): Profile | undefined {
    return profileRef.value;
  }

  created(): void {
    setBC(this, {path: "", title: "collections"});
  }

  updated(): void {
    setBC(this, {path: "", title: "collections"});
  }

  addCollection(): void {
    const title = (document.getElementById("coll-add-input") as HTMLInputElement).value.trim();
    if (socketRef.value) {
      socketRef.value.send({svc: collectionService.key, cmd: clientCommands.addCollection, param: title});
    }
  }
}
</script>
