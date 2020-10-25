module.exports = {
  configureWebpack: {
    performance: {
      maxEntrypointSize: 1024000,
      maxAssetSize: 512000
    }
  },
  chainWebpack: config => config.optimization.minimize(process.env.NODE_ENV !== 'development')
}
