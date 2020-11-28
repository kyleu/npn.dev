<template>
  <div class="uk-overflow-auto">
    <h4>{{ title }}</h4>
    <table class="uk-table uk-table-divider uk-text-left uk-table-small uk-table-justify">
      <tbody>
        <tr v-for="h in headers" :key="h.k + h.v" :title="h.desc">
          <td class="uk-text-nowrap">{{ h.k }}</td>
          <td class="uk-text-nowrap">{{ h.v }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {Header} from "@/header/model";

@Component
export default class HeadersResponse extends Vue {
  @Prop() title!: string;
  @Prop() headers!: Header[];
}
</script>
