<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right"><router-link :to="'/c/' + $route.params.coll"><Icon icon="close" /></router-link></div>
        <h3 class="uk-card-title">
          <Icon class="nav-icon-h3" icon="album" />
          {{ (coll && coll.title.length > 0) ? coll.title : $route.params.coll }}
        </h3>
        <div v-if="result" class="mt">
          <em>{{ transformer.title }} export</em> <a href="" title="copy result to clipboard" @click.prevent="copyText()"><Icon icon="copy" /></a>
          <pre class="export-result"><code ref="output">{{ result.out }}</code></pre>
        </div>
        <div v-else class="mt"><em>Loading...</em></div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {Collection} from "@/collection/collection";
import {setBC} from "@/util/vutils";
import {getCollection} from "@/collection/state";
import Icon from "@/util/Icon.vue";
import {CollectionTransformer, getCollectionTransformer} from "@/util/transformers";
import {CollectionTransformResult} from "@/transform/result";
import {getCollectionTransformResult} from "@/transform/state";

@Component({ components: { Icon } })
export default class CollectionTransform extends Vue {
  get coll(): Collection | undefined {
    return getCollection(this.$route.params.coll);
  }

  get transformer(): CollectionTransformer {
    return getCollectionTransformer(this.$route.params.fmt);
  }

  get result(): CollectionTransformResult | undefined {
    return getCollectionTransformResult(this.$route.params.coll, this.$route.params.fmt);
  }

  copyText(): void {
    const text = (this.$refs["output"] as Element).innerHTML;
    navigator.clipboard.writeText(text);
  }

  mounted(): void {
    let title = this.$route.params.coll;
    if (title === "_") {
      title = "default";
    }
    setBC(this, {path: "/c/" + this.$route.params.coll, title}, {path: "", title: "export"});
  }
}
</script>
