<template>
  <ul class="uk-list uk-list-divider">
    <li v-if="auth">Auth: {{ auth.type }}</li>
  </ul>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {Auth} from "@/auth/model";
import {requestEditingRef} from "@/request/state";

@Component
export default class AuthEditor extends Vue {
  get auth(): Auth | undefined {
    return requestEditingRef.value?.prototype?.auth
  }
}
</script>
