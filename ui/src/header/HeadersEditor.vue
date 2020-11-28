<template>
  <div v-if="(!headers) || headers.length === 0" class="mt">
    No headers defined, why not <a href="" @click.prevent="addHeader()">add one</a>?
  </div>
  <div v-else class="uk-overflow-auto">
    <table class="uk-table uk-table-divider uk-table-justify">
      <thead>
        <tr>
          <th>Name</th>
          <th>Value</th>
          <th>Description</th>
          <th style="width: 20px;"><a href="" title="new header" @click.prevent="addHeader()"><Icon icon="plus" /></a></th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(p, idx) of headers" :key="idx">
          <td><input v-model="p.k" class="uk-input" type="text" data-lpignore="true" /></td>
          <td><input v-model="p.v" class="uk-input" type="text" data-lpignore="true" /></td>
          <td><input v-model="p.desc" class="uk-input" type="text" data-lpignore="true" /></td>
          <td><a href="" title="remove param" @click.prevent="removeHeader(idx)"><Icon class="remove-link" icon="close" /></a></td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {Header} from "@/header/model";
import Icon from "@/util/Icon.vue";
import {requestEditingRef} from "@/request/state";

@Component({ components: {Icon} })
export default class HeadersEditor extends Vue {
  get headers(): Header[] | undefined {
    return requestEditingRef.value?.prototype.headers;
  }

  set headers(x: Header[] | undefined) {
    if(requestEditingRef.value) {
      requestEditingRef.value.prototype.headers = x;
    }
  }

  addHeader(): void {
    if(!this.headers) {
      this.headers = [];
    }
    this.headers.push({k: "", v: ""});
  }

  removeHeader(idx: number): void {
    if (this.headers) {
      this.headers = this.headers.filter((v, i) => i !== idx);
    }
    if ((!this.headers) || this.headers.length === 0) {
      this.headers = undefined;
    }
  }
}
</script>
