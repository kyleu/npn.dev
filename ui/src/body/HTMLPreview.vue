<template>
  <div class="preview-content"></div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";

function previewHTMLFor(html: string, baseURL: string): string {
  const headIdx = html.indexOf("<head");
  if (headIdx > -1) {
    const headEnd = html.indexOf(">", headIdx);
    if (headEnd > -1) {
      const base = `<base href="${baseURL}" target="_blank">`;
      html = html.substr(0, headEnd + 1) + base + html.substr(headEnd + 1);
    }
  }
  return html;
}

export function renderHTMLPreview(iframe: HTMLIFrameElement, html: string, baseURL: string): void {
  iframe.style.width = "100%";
  iframe.style.minHeight = "720px";

  const newHTML = previewHTMLFor(html, baseURL);

  const idoc = iframe.contentDocument || (iframe.contentWindow ? iframe.contentWindow.document : undefined);

  if (idoc) {
    idoc.open();
    idoc.write(newHTML);
    idoc.close();
  }
}

@Component
export default class HTMLPreview extends Vue {
  @Prop() html!: string
  @Prop() url!: string

  mounted(): void {
    if (this.html) {
      const iframe = document.createElement("iframe");
      this.$el.innerHTML = "";
      this.$el.appendChild(iframe);
      renderHTMLPreview(iframe, this.html, this.baseURL);
    } else {
      this.$el.innerHTML = "no preview available";
    }
  }

  get baseURL(): string {
    return new URL(this.url).origin;
  }
}
</script>
