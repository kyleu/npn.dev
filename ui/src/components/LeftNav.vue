<template>
  <div>
    <CollectionList />
    <div class="nav-section mt">
      <div class="nav-header">System</div>
      <div class="nav-list">
        <div class="nav-item"><a class="theme" href="/system"><span class="nav-icon" data-uk-icon="icon: cog"></span>System</a></div>
        <div class="nav-item"><a class="theme" href="" onclick="log.toggle();return false;"><span class="nav-icon" data-uk-icon="icon: list"></span>Toggle Log</a></div>
        <div class="nav-item"><a class="theme" href="" onclick="npn.debug();return false;"><span class="nav-icon" data-uk-icon="icon: bolt"></span>Debug</a></div>
        <div class="nav-item"><a class="theme" href="" onclick="npn.testbed();return false;"><span class="nav-icon" data-uk-icon="icon: code"></span>Testbed</a></div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import CollectionList from "@/components/CollectionList.vue";

@Component({ components: { CollectionList } })
export default class LeftNav extends Vue {
}
</script>

<style scoped lang="scss">
.left-nav {
  margin-top: 24px;
}
</style>
