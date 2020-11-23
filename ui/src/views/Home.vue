<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <h3 class="uk-card-title">Welcome to npn!</h3>

        <div class="mt">
          <div class="uk-inline" style="width: 100%;">
            <a class="uk-form-icon uk-form-icon-flip" title="Call a URL" href="" @click.prevent="runRequest()"><Icon icon="play" /></a>
            <form @submit.prevent="runRequest()">
              <input id="home-add-input" class="uk-input" type="text" placeholder="Call a URL" data-lpignore="true" />
            </form>
          </div>
        </div>
      </div>
      <div class="uk-card uk-card-body uk-card-default mt">
        <h3 class="uk-card-title">Collections</h3>
        <CollectionGallery />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import CollectionGallery from "@/collection/CollectionGallery.vue";
import {setBC} from "@/util/vutils";
import {socketRef} from "@/socket/socket";
import {requestService} from "@/util/services";
import {clientCommands} from "@/util/command";
import Icon from "@/util/Icon.vue";

@Component({ components: {Icon, CollectionGallery } })
export default class Home extends Vue {
  runRequest(): void {
    const el = document.getElementById("home-add-input") as HTMLInputElement;
    const url = el.value.trim();
    el.value = "";
    if (socketRef.value) {
      socketRef.value.send({svc: requestService.key, cmd: clientCommands.runURL, param: url});
    }
  }

  mounted(): void {
    setBC(this);
  }
}
</script>
