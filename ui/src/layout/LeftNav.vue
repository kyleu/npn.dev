<template>
  <div>
    <SessionSelector />
    <div class="nav-section mt">
      <router-link v-style-menu-section class="nav-header" tag="div" :to="{name: 'CollectionDetail', params: {coll: '_'}}">Requests</router-link>
      <RequestList />
    </div>
    <div class="nav-section mt">
      <router-link v-style-menu-section class="nav-header" tag="div" :to="{name: 'CollectionIndex'}">Collections</router-link>
      <CollectionList />
    </div>
    <div class="nav-section mt">
      <div v-style-menu-section>System</div>
      <div class="nav-list">
        <div class="nav-link"><a v-style-menu-link href="" onclick="npn.onDebug();return false"><Icon icon="code" class="nav-icon" /> Debug</a></div>
        <div class="nav-link"><router-link v-style-menu-link to="/testbed"><Icon icon="code" class="nav-icon" /> Testbed</router-link></div>
        <div class="nav-link"><a v-style-menu-link href="" @click.prevent="toggleLog()"><Icon icon="list" class="nav-icon" /> Toggle Log</a></div>
        <div class="nav-link"><router-link v-style-menu-link to="/about"><Icon icon="question" class="nav-icon" /> About</router-link></div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import CollectionList from "@/collection/CollectionList.vue";
import Icon from "@/util/Icon.vue";
import SessionSelector from "@/session/SessionSelector.vue";
import RequestList from "@/request/RequestList.vue";

@Component({ components: {RequestList, Icon, CollectionList, SessionSelector } })
export default class LeftNav extends Vue {
  toggleLog(): void {
    const el = document.getElementById("log-container");
    if (el) {
      el.style.display = el.style.display === "block" ? "none" : "block";
    }
  }
}
</script>
