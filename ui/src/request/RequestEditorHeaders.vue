<template>
  <ul class="uk-list uk-list-divider">
    <li v-if="headers">Headers: {{ headers.length }}</li>
  </ul>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {Header} from "@/header/model";

@Component
export default class RequestEditorHeaders extends Vue {
  @Prop() headers!: Header[] | undefined;
}
</script>
