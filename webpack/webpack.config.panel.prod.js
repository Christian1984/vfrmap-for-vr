const path = require("path");
const { merge } = require("webpack-merge");

const { panelBaseConfig } = require("./webpack.config.panel.base");
const { prodConfig } = require("./webpack.config.common");

module.exports = merge(prodConfig, panelBaseConfig);