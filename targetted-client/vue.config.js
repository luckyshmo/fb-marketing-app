const path = require("path");

const vueSrc = "./src";

module.exports = {
  lintOnSave: false,
  devServer: {
    host: 'localhost'
  },
  configureWebpack: {
    resolve: {
      alias: {
        "@": path.resolve(__dirname, vueSrc)
      },
    }
  }
}
