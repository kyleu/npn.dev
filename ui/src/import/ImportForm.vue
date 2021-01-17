<template>
  <div class="uk-section uk-section-small">
    <div class="uk-container uk-container-expand uk-position-relative">
      <div class="uk-card uk-card-body uk-card-default">
        <div class="right">
          <router-link to="/"><Icon icon="close" /></router-link>
        </div>
        <h3 class="uk-card-title"><Icon icon="upload" class="nav-icon-h3" /> Import</h3>
        <p>Upload your requests or collections in npn, Postman, or OpenAPI format.</p>

        <div class="mt">
          <ul class="uk-child-width-expand" data-uk-tab>
            <li><a href="#">Upload a File</a></li>
            <li><a href="#">Paste Text</a></li>
          </ul>
          <ul class="uk-switcher">
            <li>
              <form method="post" enctype="multipart/form-data" :action="actionURL">
                <div class="js-upload uk-placeholder uk-text-center">
                  <span class="uk-text-middle">Upload files by dropping them here or </span>
                  <div data-uk-form-custom>
                    <input type="file" name="file" multiple>
                    <span v-style-link>selecting one</span>
                  </div>
                </div>
                <div class="right"><button class="uk-button uk-button-default">Upload File</button></div>
              </form>
            </li>
            <li>
              <form method="post" enctype="application/x-www-form-urlencoded" :action="actionURL">
                <textarea style="height: 87px;" class="uk-textarea" rows="4" name="content" placeholder="Paste your file contents here"></textarea>
                <div class="right mt"><button class="uk-button uk-button-default">Upload Contents</button></div>
              </form>
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
import {hostRef} from "@/socket/socket";

@Component({ components: { Icon } })
export default class ImportForm extends Vue {
  get actionURL(): string {
    const h = hostRef.value;
    console.log(h + "!!!");
    if (h.length > 0) {
      return "http://" + h + "/i";
    }
    return "/i";
  }

  mounted(): void {
    setBC(this, { path: "", title: "import" });
  }
}
</script>
