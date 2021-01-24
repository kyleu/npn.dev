const IN_PRODUCTION = process.env.NODE_ENV === "production";

module.exports = {
  plugins: [
    IN_PRODUCTION &&
      require("@fullhuman/postcss-purgecss")({
        content: [`./public/**/*.html`, `./src/**/*.vue`],
        defaultExtractor(content) {
          const contentWithoutStyleBlocks = content.replace(
            /<style[^]+?<\/style>/gi,
            ""
          );
          return (
            contentWithoutStyleBlocks.match(
              /[A-Za-z0-9-_/:]*[A-Za-z0-9-_/]+/g
            ) || []
          );
        },
        safelist: [
          /CodeMirror.*/,
          /cm-.*/,
          /log-.*/,
          /uk-dropdown.*/,
          /uk-grid.*/,
          /uk-hidden.*/,
          /uk-navbar.*/,
          /uk-offcanvas.*/,
          /uk-open.*/,
          /uk-sticky.*/,
          /uk-tab.*/,
          /uk-toggle.*/,
          /uk-visible.*/,
          /uk-form-custom.*/,
          /-(leave|enter|appear)(|-(to|from|active))$/,
          /^(?!(|.*?:)cursor-move).+-move$/,
          /^router-link(|-exact)-active$/,
          /data-v-.*/,
        ],
      }),
  ],
};
