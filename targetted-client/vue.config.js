const path = require("path");

const vueSrc = "./src";

module.exports = {
  lintOnSave: false,
  configureWebpack: {
    resolve: {
      alias: {
        "@": path.resolve(__dirname, vueSrc)
      },
    }
  }
}
