<template>
  <ul class="uk-list uk-list-divider">
    <li>
      <div data-uk-grid="">
        <div class="uk-width-1-4">Name</div>
        <div class="uk-width-1-4">Value</div>
        <div class="uk-width-1-2">
          <div class="right">
            <a href="" title="new param" @click.prevent="addParam()">
              <span data-uk-icon="icon: plus" />
            </a>
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
            <a href="" title="remove param" @click.prevent="removeParam(idx)">
              <span data-uk-icon="icon: close" />
            </a>
          </div>
          <input v-model="p.desc" style="width: calc(100% - 36px);" class="uk-input" type="text" />
        </div>
      </div>
    </li>
  </ul>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {QueryParam} from "@/request/model";

@Component
export default class QueryParamsEditor extends Vue {
  @Prop() qp: QueryParam[] | undefined;

  addParam(): void {
    if(!this.qp) {
      this.qp = [];
    }
    if(this.qp) {
      this.qp.push({k: "", v: ""});
    }
  }

  removeParam(idx: number): void {
    if (this.qp) {
      this.qp = this.qp.splice(idx, 1);
    }
    if (this.qp?.length === 0) {
      this.qp = undefined;
    }
  }
}
</script>
