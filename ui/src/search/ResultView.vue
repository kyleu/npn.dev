<template>
  <li class="search-result">
    <div>
      <div class="right"><em>{{ r.loc }}</em></div>
      <router-link :to="url">
        <Icon v-if="r.t === 'collection'" class="nav-icon" icon="folder" />
        <Icon v-else-if="r.t === 'request'" class="nav-icon" icon="link" />
        <Icon v-else class="nav-icon" icon="unknown" />
        {{ r.key }}
      </router-link>
    </div>
    <div class="match">{{ r.prelude }}<strong>{{ q }}</strong>{{ r.postlude }}</div>
  </li>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import Icon from "@/util/Icon.vue";
import {SearchResult} from "@/search/state";
import {logError} from "@/util/log";

@Component({ components: { Icon } })
export default class ResultView extends Vue {
  @Prop() q!: SearchResult
  @Prop() r!: SearchResult

  get url(): string {
    switch(this.r.t) {
      case "collection":
        return "/c/" + this.r.key;
      case "request":
        return "/c/" + this.r.key;
      default:
        logError("unknown search result type [" + this.r.t + "]");
        return "unknown";
    }
  }
}
</script>
