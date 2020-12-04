<template>
  <div v-if="req.prototype" class="mt uk-panel">
    <div class="left" style="width:120px;">
      <select v-model="req.prototype.method" class="uk-select" name="method">
        <option v-for="m in methods" :key="m.key">{{ m.key }}</option>;
      </select>
    </div>
    <div v-if="!editing" class="url-view uk-inline right" style="width:calc(100% - 120px);">
      <a class="uk-form-icon uk-form-icon-flip" title="send request" style="padding-top: 4px;" @click.prevent="doCall()"><Icon icon="play" /></a>
      <div @click="editing = true">
        <span class="url-link"><span v-for="part in protoParts" :key="part.idx" :style="{color: part.color}" :title="part.t">{{ part.v }}</span> </span>
      </div>
    </div>
    <div v-if="editing" class="url-input uk-inline right" style="width:calc(100% - 120px);">
      <a class="uk-form-icon uk-form-icon-flip" title="cancel edit" href="" @click.prevent="editing = false"><Icon icon="close" /></a>
      <form @submit.prevent="doCall()">
        <input id="url-input-el" v-model="protoString" class="uk-input" name="url" type="text" data-lpignore="true" @blur="editing = false" />
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {allMethods, Method, NPNRequest} from "@/request/model";
import {Part, prototypeToURL, prototypeToURLParts} from "@/request/prototype/url";
import {prototypeFromURL} from "@/request/prototype/prototype";
import {requestEditingRef} from "@/request/state";
import Icon from "@/util/Icon.vue";
import {requestResultsRef} from "@/call/state";

@Component({ components: {Icon} })
export default class URLEditor extends Vue {
  private e = false

  get editing(): boolean {
    return this.e;
  }
  set editing(e: boolean) {
    this.e = e;
    if (e) {
      setTimeout(() => {
        document.getElementById("url-input-el")?.focus();
      }, 0);
    }
  }

  get req(): NPNRequest | undefined {
    return requestEditingRef.value;
  }

  get methods(): Method[] {
    return allMethods;
  }

  get protoString(): string {
    return prototypeToURL(this.req?.prototype);
  }

  set protoString(s: string) {
    if (this.req) {
      const o = this.req.prototype;
      const n = prototypeFromURL(s);
      o.protocol = n.protocol;
      o.domain = n.domain;
      o.port = n.port;
      o.path = n.path;
      o.query = n.query;
      o.fragment = n.fragment;
      if ((!o.auth) || o.auth.type === "basic") {
        o.auth = n.auth;
      }
    }
  }

  get protoParts(): Part[] {
    if (!this.req) {
      return [];
    }
    return prototypeToURLParts(this.req?.prototype);
  }

  doCall(): void {
    if (this.$route.name === 'CallResult') {
      requestResultsRef.value = {id: "", coll: "", req: "", cycles: []};
    } else {
      this.$router.push({name: "CallResult", params: {coll: this.$route.params.coll, req: this.$route.params.req}});
    }
  }
}
</script>
