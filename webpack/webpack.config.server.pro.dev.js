const { merge } = require("webpack-merge");

const { devConfig } = require("./webpack.config.common");
const { indexBaseConfig } = require("./webpack.config.server.base");
const { premiumConfig, mapsProConfig } = require("./webpack.config.server.pro");

const proDevConfig = merge(devConfig, premiumConfig);
const indexDevConfig = merge(devConfig, indexBaseConfig);
const mapsDevConfig = merge(devConfig, mapsProConfig);

module.exports = [proDevConfig, indexDevConfig, mapsDevConfig];