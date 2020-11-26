<template>
  <div v-if="(!cookies) || cookies.length === 0">
    No cookies defined, why not <a href="" @click.prevent="addCookie()">add one</a>?
  </div>
  <ul v-else class="uk-list uk-list-divider">
    <li>
      <div data-uk-grid="">
        <div class="uk-width-1-4">Name</div>
        <div class="uk-width-1-4">Value</div>
        <div class="uk-width-1-4">Domain</div>
        <div class="uk-width-1-4">
          <div class="right">
            <a href="" title="new cookie" @click.prevent="addCookie()"><Icon icon="plus" /></a>
          </div>
          Path
        </div>
      </div>
    </li>
    <li v-for="(p, idx) of cookies" :key="idx">
      <div data-uk-grid="">
        <div class="uk-width-1-4">
          <input v-model="p.name" class="uk-input" type="text" />
        </div>
        <div class="uk-width-1-4">
          <input v-model="p.value" class="uk-input" type="text" />
        </div>
        <div class="uk-width-1-4">
          <input v-model="p.domain" class="uk-input" type="text" />
        </div>
        <div class="uk-width-1-4">
          <div class="right" style="margin-top: 6px;">
            <a href="" title="remove cookie" @click.prevent="removeCookie(idx)"><Icon icon="close" /></a>
          </div>
          <input v-model="p.path" style="width: calc(100% - 36px);" class="uk-input" type="text" />
        </div>
      </div>
    </li>
  </ul>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import Icon from "@/util/Icon.vue";
import {sessionEditingRef} from "@/session/state";
import {Cookie} from "@/session/model";

@Component({ components: {Icon} })
export default class CookiesEditor extends Vue {
  get cookies(): Cookie[] | undefined {
    return sessionEditingRef.value?.cookies;
  }

  set cookies(x: Cookie[] | undefined) {
    if(sessionEditingRef.value) {
      sessionEditingRef.value.cookies = x || [];
    }
  }

  addCookie(): void {
    if(!this.cookies) {
      this.cookies = [];
    }
    this.cookies.push({name: "", value: "", domain: "", path: ""});
  }

  removeCookie(idx: number): void {
    if (this.cookies) {
      this.cookies = this.cookies.filter((v, i) => i !== idx);
    }
    if ((!this.cookies) || this.cookies.length === 0) {
      this.cookies = undefined;
    }
  }
}
</script>
