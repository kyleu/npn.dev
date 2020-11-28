<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right"><router-link to="/c"><Icon icon="close" /></router-link></div>
        <h3 class="uk-card-title">
          <Icon class="nav-icon-h3" icon="album" />
          {{ (coll && coll.title.length > 0) ? coll.title : $route.params.coll }}
        </h3>
        <em>{{ transformer.title }} export</em>
        <pre v-if="result"><code>{{ result.out }}</code></pre>
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
import {CollectionTransformResult} from "@/request/transform/result";
import {getCollectionTransformResult} from "@/request/transform/state";

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

  mounted(): void {
    let title = this.$route.params.coll;
    if (title === "_") {
      title = "default";
    }
    setBC(this, {path: "/c/" + this.$route.params.coll, title}, {path: "", title: "export"});
  }
}
</script>
