<template>
  <div v-if="req.prototype" class="mt uk-panel">
    <div class="left" style="width:120px;">
      <select v-model="req.prototype.method" class="uk-select" name="method">
        <option v-for="m in methods" :key="m.key">{{ m.key }}</option>;
      </select>
    </div>
    <div class="url-view uk-inline right" style="width:calc(100% - 120px);">
      <a class="uk-form-icon uk-form-icon-flip" href="" onclick="TODO(); call" title="send request" data-uk-icon="icon: play" />
      <div onclick="TODO();">
        <span class="url-link">{{ protoHTML }}</span>
      </div>
    </div>
    <div class="url-input hidden uk-inline right" style="width:calc(100% - 120px);">
      <a class="uk-form-icon uk-form-icon-flip" href="" onclick="TODO();return false;" title="cancel edit" data-uk-icon="icon: close" />
      <form onsubmit="TODO();return false;">
        <input class="uk-input" name="url" type="text" :value="req.prototype.path" data-lpignore="true" />
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {allMethods, Method, NPNRequest} from "@/request/model";
import {prototypeToURL} from "@/request/url";

@Component
export default class URLEditor extends Vue {
  @Prop() req!: NPNRequest

  get methods(): Method[] {
    return allMethods;
  }

  get protoString(): string {
    return prototypeToURL(this.req?.prototype);
  }

  get protoHTML(): string {
    return prototypeToURL(this.req?.prototype);
  }
}
</script>
