<template>
  <div v-if="(!qp) || qp.length === 0">
    No query parameters defined, why not <a href="" @click.prevent="addParam()">add one</a>?
  </div>
  <ul v-else class="uk-list uk-list-divider">
    <li>
      <div data-uk-grid="">
        <div class="uk-width-1-4">Name</div>
        <div class="uk-width-1-4">Value</div>
        <div class="uk-width-1-2">
          <div class="right">
            <a href="" title="new param" @click.prevent="addParam()"><Icon icon="plus" /></a>
          </div>
          Description
        </div>
      </div>
    </li>
    <li v-for="(p, idx) of qp" :key="idx">
      <div data-uk-grid="">
        <div class="uk-width-1-4">
          <input v-model="p.k" class="uk-input" type="text" />
        </div>
        <div class="uk-width-1-4">
          <input v-model="p.v" class="uk-input" type="text" />
        </div>
        <div class="uk-width-1-2">
          <div class="right" style="margin-top: 6px;">
            <a href="" title="remove param" @click.prevent="removeParam(idx)"><Icon icon="close" /></a>
          </div>
          <input v-model="p.desc" style="width: calc(100% - 36px);" class="uk-input" type="text" />
        </div>
      </div>
    </li>
  </ul>
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
