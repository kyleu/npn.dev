<template>
  <li>
    <router-link v-style-menu-link class="session-link mt" :to="'/s/' + sess.key">
      {{ label }}
    </router-link>
  </li>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import Icon from "@/util/Icon.vue";
import {SessionSummary} from "@/session/session";

@Component({ components: {Icon} })
export default class SessionListItem extends Vue {
  @Prop() sess!: SessionSummary;

  get label(): string {
    return ((!this.sess.title) || this.sess.title.length === 0) ? this.sess.key : this.sess.title;
  }
}
</script>
