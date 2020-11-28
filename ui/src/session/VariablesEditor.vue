<template>
  <div v-if="(!variables) || variables.length === 0" class="mt">
    No variables defined, why not <a href="" @click.prevent="addVariable()">add one</a>?
  </div>
  <div v-else class="uk-overflow-auto">
    <table class="uk-table uk-table-divider uk-table-justify">
      <thead>
        <tr>
          <th>Name</th>
          <th>Value</th>
          <th>Description</th>
          <th style="width: 20px;"><a href="" title="new variable" @click.prevent="addVariable()"><Icon icon="plus" /></a></th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(p, idx) of variables" :key="idx">
          <td><input v-model="p.k" class="uk-input" type="text" data-lpignore="true" /></td>
          <td><input v-model="p.v" class="uk-input" type="text" data-lpignore="true" /></td>
          <td><input v-model="p.desc" class="uk-input" type="text" data-lpignore="true" /></td>
          <td><a href="" title="remove variable" @click.prevent="removeVariable(idx)"><Icon class="remove-link" icon="close" /></a></td>
        </tr>
      </tbody>
    </table>
  </div>
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
