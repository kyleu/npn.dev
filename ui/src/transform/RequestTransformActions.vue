<template>
  <div class="uk-inline">
    <button v-style-button type="button" class="uk-button uk-button-default mrs mt">Export</button>
    <div id="export-dropdown" data-uk-dropdown="mode: click">
      <ul class="uk-list uk-list-divider" style="margin-bottom: 0;">
        <li v-for="(v, n) in transforms" :key="n">
          <router-link :to="'/c/' + $route.params.coll + '/' + $route.params.req + '/transform/' + v.key" onclick="UIkit.dropdown(this.parentElement.parentElement.parentElement).hide(false);" :title="v.description">{{ v.title }}</router-link>
        </li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {RequestTransformers, RequestTransformer} from "@/util/transformers";

@Component
export default class RequestTransformActions extends Vue {
  get transforms(): RequestTransformer[] {
    return RequestTransformers;
  }
}
</script>
