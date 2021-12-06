const path = require('path')

const vueSrc = './src'

module.exports = {
  lintOnSave: false,
  // devServer: {
  //   host: 'localhost',
  //   proxy: 'http://localhost:8087'
  // },
  configureWebpack: {
    resolve: {
      alias: {
        '@src': path.resolve(__dirname, vueSrc)
      }
    }
  }
}
