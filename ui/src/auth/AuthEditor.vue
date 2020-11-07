<template>
  <ul class="uk-list uk-list-divider">
    <li v-if="auth">Auth: {{ auth.type }}</li>
  </ul>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {Auth} from "@/auth/model";

@Component
export default class AuthEditor extends Vue {
  @Prop() auth: Auth | undefined;
}
</script>
