<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right"><router-link to="/"><Icon icon="close" /></router-link></div>
        <h3 class="uk-card-title"><Icon icon="bookmark" size="24" class="nav-icon-h3" /> Sessions</h3>
        <SessionList />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {setBC} from "@/util/vutils";
import Icon from "@/util/Icon.vue";
import SessionList from "@/session/SessionList.vue";

@Component({components: {SessionList, Icon}})
export default class SessionIndex extends Vue {
  mounted(): void {
    setBC(this, {path: "", title: "sessions"});
  }
}
</script>
