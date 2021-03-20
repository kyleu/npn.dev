<template>
  <div>
    <div class="nav-section">
      <router-link v-style-menu-section class="nav-header" tag="div" :to="{name: 'CollectionDetail', params: {coll: '_'}}">Requests</router-link>
      <RequestList />
    </div>
    <div class="nav-section mt">
      <router-link v-style-menu-section class="nav-header" tag="div" :to="{name: 'CollectionIndex'}">Collections</router-link>
      <CollectionList />
    </div>
    <div class="nav-section mt">
      <router-link v-style-menu-section class="nav-header" tag="div" :to="{name: 'SessionIndex'}">Sessions</router-link>
      <SessionSelector />
    </div>
    <div class="nav-section mt">
      <div v-style-menu-section>System</div>
      <div class="nav-list">
        <div class="nav-link"><router-link v-style-menu-link to="/i"><Icon icon="upload" class="nav-icon" /> Import</router-link></div>
        <div class="nav-link"><router-link v-style-menu-link to="/cfg"><Icon icon="cog" class="nav-icon" /> Config</router-link></div>
        <div class="nav-link"><router-link v-style-menu-link to="/help"><Icon icon="question" class="nav-icon" /> Help</router-link></div>
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
}
</script>
