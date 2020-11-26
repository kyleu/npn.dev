<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right"><router-link to="/about"><Icon title="about npn" icon="question" /></router-link></div>
        <h3 class="uk-card-title">npn</h3>
        <p>You're using <a href="https://npn.dev">npn</a>, an HTTP client that helps you document and test APIs</p>

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
        <h3 class="uk-card-title">Sessions</h3>
        <p>Each session contains variables used in the request and cookies from responses</p>
        <SessionList />
      </div>
      <div class="uk-card uk-card-body uk-card-default mt">
        <h3 class="uk-card-title">Collections</h3>
        <p>Store related URLs in a collection of requests, which you can run in bulk or share with colleagues</p>
        <CollectionGallery />
      </div>
      <div class="uk-card uk-card-body uk-card-default mt">
        <h3 class="uk-card-title">Recent Requests</h3>
        <p>TODO</p>
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
import SessionList from "@/session/SessionList.vue";

@Component({ components: {SessionList, Icon, CollectionGallery } })
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
