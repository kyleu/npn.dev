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
          <ul id="message-list" class="uk-list uk-list-divider" data-uk-accordion="multiple: true">
            <li v-for="(r, idx) in res.results" :key="idx">
              <a class="uk-accordion-title" href="#">{{ r.filename }} ({{ r.type }})</a>
              <div class="uk-accordion-content">
                <div v-if="r.error" class="uk-alert-danger" style="margin: 12px 0; padding: 6px;">{{ r.error }}</div>
                <pre>{{ r.value }}</pre>
              </div>
            </li>
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

@Component({ components: { Icon } })
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
