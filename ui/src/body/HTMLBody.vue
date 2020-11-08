<template>
  <div class="uk-overflow-auto">
    <div v-if="mode === 'text'" class="right">
      <a :class="profile.settings.linkColor + '-fg'" href="" title="render preview" @click.prevent="mode = 'preview'">preview</a>
    </div>
    <div v-else class="right">
      <a :class="profile.settings.linkColor + '-fg'" href="" title="render preview" @click.prevent="mode = 'text'">text</a>
    </div>
    <em>HTML</em>

    <div v-if="mode === 'text'">
      <pre class="prism-view mt" style="margin: 0;"><code v-html="highlighted"></code></pre>
    </div>
    <div v-else>
      <div class="mt"><HTMLPreview :url="url" :html="content" /></div>
    </div>
  </div>
</template>

<script lang="ts">
import {Profile, profileRef} from "@/user/profile";
import {Component, Prop, Vue} from "vue-property-decorator";
import {HTMLConfig} from "@/body/model";
import {PrismEditor} from "vue-prism-editor";
import HTMLPreview from "@/body/HTMLPreview.vue";

// @ts-ignore
// eslint-disable-next-line
declare const Prism: any;

@Component({ components: { HTMLPreview, PrismEditor } })
export default class HTMLBody extends Vue {
  @Prop() url!: string
  @Prop() config!: HTMLConfig

  mode = "text"

  get highlighted(): string {
    return Prism.highlight(this.config.content, Prism.languages.html);
  }

  get profile(): Profile | undefined {
    return profileRef.value;
  }
}
</script>
