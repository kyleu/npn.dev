<template>
  <div class="uk-overflow-auto">
    <h4>BODY!</h4>
    <p><em>{{ url }}</em></p>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import {RBody} from "@/body/model";

@Component
export default class ResultBody extends Vue {
  @Prop() url!: string
  @Prop() body!: RBody
}
</script>
