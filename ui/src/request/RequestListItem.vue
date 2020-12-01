<template>
  <div class="nav-link">
    <router-link v-style-menu-link class="req-link" :to="'/c/_/' + req.key" :title="label">
      <Icon icon="link" class="nav-icon" />
      {{ label }}
    </router-link>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {Summary} from "@/request/model";
import Icon from "@/util/Icon.vue";
@Component({
  components: {Icon}
})
export default class RequestListItem extends Vue {
  @Prop() req!: Summary;

  get label(): string {
    return ((!this.req.title) || this.req.title.length === 0) ? this.req.key : this.req.title;
  }
}
</script>
