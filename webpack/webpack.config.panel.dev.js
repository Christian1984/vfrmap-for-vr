const path = require("path");
const { merge } = require("webpack-merge");

const { panelBaseConfig } = require("./webpack.config.panel.base");
const { devConfig } = require("./webpack.config.common");

module.exports = merge(devConfig, panelBaseConfig);