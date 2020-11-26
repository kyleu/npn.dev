<template>
  <div v-if="(!variables) || variables.length === 0">
    No variables defined, why not <a href="" @click.prevent="addVariable()">add one</a>?
  </div>
  <ul v-else class="uk-list uk-list-divider">
    <li>
      <div data-uk-grid="">
        <div class="uk-width-1-4">Name</div>
        <div class="uk-width-1-4">Value</div>
        <div class="uk-width-1-2">
          <div class="right">
            <a href="" title="new variable" @click.prevent="addVariable()"><Icon icon="plus" /></a>
          </div>
          Description
        </div>
      </div>
    </li>
    <li v-for="(p, idx) of variables" :key="idx">
      <div data-uk-grid="">
        <div class="uk-width-1-4">
          <input v-model="p.k" class="uk-input" type="text" />
        </div>
        <div class="uk-width-1-4">
          <input v-model="p.v" class="uk-input" type="text" />
        </div>
        <div class="uk-width-1-2">
          <div class="right" style="margin-top: 6px;">
            <a href="" title="remove variable" @click.prevent="removeVariable(idx)"><Icon icon="close" /></a>
          </div>
          <input v-model="p.desc" style="width: calc(100% - 36px);" class="uk-input" type="text" />
        </div>
      </div>
    </li>
  </ul>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import Icon from "@/util/Icon.vue";
import {sessionEditingRef} from "@/session/state";
import {Variable} from "@/session/model";

@Component({ components: {Icon} })
export default class VariablesEditor extends Vue {
  get variables(): Variable[] | undefined {
    return sessionEditingRef.value?.variables;
  }

  set variables(x: Variable[] | undefined) {
    if(sessionEditingRef.value) {
      sessionEditingRef.value.variables = x || [];
    }
  }

  addVariable(): void {
    if(!this.variables) {
      this.variables = [];
    }
    this.variables.push({k: "", v: ""});
  }

  removeVariable(idx: number): void {
    if (this.variables) {
      this.variables = this.variables.filter((v, i) => i !== idx);
    }
    if ((!this.variables) || this.variables.length === 0) {
      this.variables = undefined;
    }
  }
}
</script>
