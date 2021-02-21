<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right">
          <router-link to="/i"><Icon icon="close" /></router-link>
        </div>
        <h3 class="uk-card-title"><Icon icon="upload" class="nav-icon-h3" /> Import [{{ res.key }}]</h3>
        <em>{{ res.cfg.files.length }} files</em>
        <div v-if="res" class="mt">
          <table class="uk-table uk-table-divider">
            <thead>
              <tr>
                <th scope="col">Filename</th>
                <th scope="col">Type</th>
                <th scope="col">Collection</th>
                <th scope="col">Requests</th>
                <th scope="col">Cookies</th>
                <th scope="col">Variables</th>
              </tr>
            </thead>
            <tbody>
              <ImportResultRow v-for="(r, idx) in res.results" :key="idx" :file="r" />
            </tbody>
          </table>
        </div>
      </div>
      <div class="uk-card uk-card-body uk-card-default mt">
        <h3 class="uk-card-title">File Details</h3>
        <div v-if="res" class="mt">
          <ul class="uk-list uk-list-divider" data-uk-accordion="multiple: true">
            <ImportResultSection v-for="(r, idx) in res.results" :key="idx" :file="r" />
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { setBC } from "@/util/vutils";
import Icon from "@/util/Icon.vue";
import {importResultRef, setActiveImport} from "@/import/state";
import {ImportResult} from "@/import/model";
import {jsonStr} from "@/util/json";
import ImportResultRow from "@/import/ImportResultRow.vue";
import ImportResultSection from "@/import/ImportResultSection.vue";

@Component({ components: {ImportResultRow, ImportResultSection, Icon } })
export default class ImportResults extends Vue {
  get res(): ImportResult | undefined {
    setActiveImport(this.$route.params.id);
    return importResultRef.value;
  }

  get resJSON(): string {
    return jsonStr(this.res);
  }

  mounted(): void {
    setBC(this, { path: "/i", title: "import" }, { path: "", title: "results" });
  }
}
</script>
