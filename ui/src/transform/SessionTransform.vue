<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right"><router-link :to="'/s/' + $route.params.sess"><Icon icon="close" /></router-link></div>
        <h3 class="uk-card-title">
          <Icon class="nav-icon-h3" icon="album" />
          {{ (sess && sess.title.length > 0) ? sess.title : ($route.params.sess === "_" ? "Default Session" : $route.params.sess) }}
        </h3>
        <div v-if="result" class="mt">
          <em>Session export</em> <a href="" title="copy result to clipboard" @click.prevent="copyText()"><Icon icon="copy" /></a>
          <pre class="export-result"><code ref="output">{{ result.out }}</code></pre>
        </div>
        <div v-else class="mt"><em>Loading...</em></div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {setBC} from "@/util/vutils";
import Icon from "@/util/Icon.vue";
import {SessionTransformResult} from "@/transform/result";
import {getSessionTransformResult} from "@/transform/state";
import {Session} from "@/session/model";
import {getSessionDetail} from "@/session/state";

@Component({ components: { Icon } })
export default class SessionTransform extends Vue {
  get sess(): Session | undefined {
    return getSessionDetail(this.$route.params.sess);
  }

  get result(): SessionTransformResult | undefined {
    return getSessionTransformResult(this.$route.params.sess);
  }

  copyText(): void {
    const text = (this.$refs["output"] as Element).innerHTML;
    navigator.clipboard.writeText(text);
  }

  mounted(): void {
    let title = this.$route.params.sess;
    if (title === "_") {
      title = "default";
    }
    setBC(this, {path: "/s/" + this.$route.params.sess, title}, {path: "", title: "export"});
  }
}
</script>
