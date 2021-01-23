<template>
  <div>
    <div class="mt">
      <div class="uk-inline" style="width: 100%;">
        <a class="uk-form-icon uk-form-icon-flip" title="add collection" href="" @click.prevent="addCollection()"><Icon icon="plus" /></a>
        <form @submit.prevent="addCollection()">
          <input id="coll-add-input" class="uk-input" type="text" placeholder="Add collection" data-lpignore="true" />
        </form>
      </div>
    </div>

    <ul class="uk-list uk-list-divider mt">
      <CollectionGalleryItem v-for="cc in collections" :key="cc.coll.key" :cc="cc" />
      <li v-if="(!collections) || collections.length === 0"><em>no collections</em></li>
    </ul>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import CollectionGalleryItem from "@/collection/CollectionGalleryItem.vue";
import {CollectionCount} from "@/collection/collection";
import {collectionsRef} from "@/collection/state";
import {socketRef} from "@/socket/socket";
import {collectionService} from "@/util/services";
import {clientCommands} from "@/util/command";
import Icon from "@/util/Icon.vue";

@Component({ components: {Icon, CollectionGalleryItem } })
export default class CollectionGallery extends Vue {
  get collections(): CollectionCount[] {
    return collectionsRef.value;
  }

  addCollection(): void {
    const el = document.getElementById("coll-add-input") as HTMLInputElement;
    const title = el.value.trim();
    if (socketRef.value) {
      socketRef.value.send({channel: collectionService.key, cmd: clientCommands.addCollection, param: title});
    }
  }
}
</script>
