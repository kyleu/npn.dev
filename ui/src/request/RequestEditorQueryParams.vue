<template>
  <ul v-if="qp" class="uk-list uk-list-divider">
    <li>
      <div data-uk-grid="">
        <div class="uk-width-1-4">Name</div>
        <div class="uk-width-1-4">Value</div>
        <div class="uk-width-1-2">
          <div class="right">
            <a :class="$store.state.profile.linkColor + '-fg'" href="" onclick="return false" title="new param">
              <span data-uk-icon="icon: plus" />
            </a>
          </div>
          Description
        </div>
      </div>
    </li>
    <li v-for="p of qp" :key="p.k">
      <div data-uk-grid="">
        <div class="uk-width-1-4">
          <input v-model="p.k" class="uk-input" type="text" />
        </div>
        <div class="uk-width-1-4">
          <input v-model="p.v" class="uk-input" type="text" />
        </div>
        <div class="uk-width-1-2">
          <div class="right" style="margin-top: 6px;">
            <a :class="$store.state.profile.linkColor + '-fg'" href="" onclick="return false;" title="remove param">
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
export default class RequestEditorQueryParams extends Vue {
  @Prop() qp!: QueryParam[] | undefined;
}
</script>
