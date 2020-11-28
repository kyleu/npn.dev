<template>
  <div v-if="(!cookies) || cookies.length === 0" class="mt">
    No cookies defined, why not <a href="" @click.prevent="addCookie()">add one</a>?
  </div>
  <div v-else class="uk-overflow-auto">
    <table class="uk-table uk-table-divider uk-table-justify">
      <thead>
        <tr>
          <th>Name</th>
          <th>Value</th>
          <th>Domain</th>
          <th>Path</th>
          <th style="width: 20px;"><a href="" title="new cookie" @click.prevent="addCookie()"><Icon icon="plus" /></a></th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(p, idx) of cookies" :key="idx">
          <td><input v-model="p.name" class="uk-input" type="text" data-lpignore="true" /></td>
          <td><input v-model="p.value" class="uk-input" type="text" data-lpignore="true" /></td>
          <td><input v-model="p.domain" class="uk-input" type="text" data-lpignore="true" /></td>
          <td><input v-model="p.path" class="uk-input" type="text" data-lpignore="true" /></td>
          <td><a href="" title="remove cookie" @click.prevent="removeCookie(idx)"><Icon class="remove-link" icon="close" /></a></td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import Icon from "@/util/Icon.vue";
import {sessionEditingRef} from "@/session/state";
import {Cookie} from "@/session/model";

@Component({ components: {Icon} })
export default class CookieEditor extends Vue {
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
