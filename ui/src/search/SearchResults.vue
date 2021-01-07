<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right">
          <router-link to="/"><Icon icon="close" /></router-link>
        </div>
        <h3 class="uk-card-title"><Icon icon="search" class="nav-icon-h3" /> [{{ query }}] Search Results</h3>
        <div v-if="results.length === 0" class="mt">
          No results
        </div>
        <ul v-else class="uk-list uk-list-divider">
          <ResultView v-for="x in results" :key="x.t + x.key" :r="x" :q="query" />
        </ul>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { setBC } from "@/util/vutils";
import Icon from "@/util/Icon.vue";
import {onSearchQuery, searchQueryRef, SearchResult, searchResultsRef} from "@/search/state";
import ResultView from "@/search/ResultView.vue";

@Component({ components: {ResultView, Icon } })
export default class SearchResults extends Vue {
  get query(): string {
    return searchQueryRef.value.q;
  }

  get results(): SearchResult[] {
    return searchResultsRef.value;
  }

  updated(): void {
    onSearchQuery(this.$route.params["q"], 0);
  }

  mounted(): void {
    onSearchQuery(this.$route.params["q"], 0);
    setBC(this, { path: "", title: "search" });
  }
}
</script>
