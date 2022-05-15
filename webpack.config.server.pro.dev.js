const { merge } = require("webpack-merge");

const { devConfig } = require("./webpack.config.common");
const { indexBaseConfig, mapsBaseConfig } = require("./webpack.config.server.base");
const { proConfig } = require("./webpack.config.server.pro");

const proDevConfig = merge(devConfig, proConfig);
const indexDevConfig = merge(devConfig, indexBaseConfig);
const mapsDevConfig = merge(devConfig, mapsBaseConfig);

module.exports = [proDevConfig, indexDevConfig, mapsDevConfig];