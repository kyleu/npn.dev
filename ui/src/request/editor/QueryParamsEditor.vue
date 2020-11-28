<template>
  <div v-if="(!qp) || qp.length === 0" class="mt">
    No query parameters defined, why not <a href="" @click.prevent="addParam()">add one</a>?
  </div>
  <div v-else class="uk-overflow-auto">
    <table class="uk-table uk-table-divider uk-table-justify">
      <thead>
        <tr>
          <th>Name</th>
          <th>Value</th>
          <th>Description</th>
          <th style="width: 20px;"><a href="" title="new param" @click.prevent="addParam()"><Icon icon="plus" /></a></th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(p, idx) of qp" :key="idx">
          <td><input v-model="p.k" class="uk-input" type="text" data-lpignore="true" /></td>
          <td><input v-model="p.v" class="uk-input" type="text" data-lpignore="true" /></td>
          <td><input v-model="p.desc" class="uk-input" type="text" data-lpignore="true" /></td>
          <td><a href="" title="remove param" @click.prevent="removeParam(idx)"><Icon class="remove-link" icon="close" /></a></td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {QueryParam} from "@/request/model";
import Icon from "@/util/Icon.vue";
import {requestEditingRef} from "@/request/state";

@Component({ components: {Icon} })
export default class QueryParamsEditor extends Vue {
  get qp(): QueryParam[] | undefined {
    return requestEditingRef.value?.prototype.query;
  }

  set qp(x: QueryParam[] | undefined) {
    if(requestEditingRef.value) {
      requestEditingRef.value.prototype.query = x;
    }
  }

  addParam(): void {
    if(!this.qp) {
      this.qp = [];
    }
    this.qp.push({k: "", v: ""});
  }

  removeParam(idx: number): void {
    if (this.qp) {
      this.qp = this.qp.filter((v, i) => i !== idx);
    }
    if ((!this.qp) || this.qp.length === 0) {
      this.qp = undefined;
    }
  }
}
</script>
