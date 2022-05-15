const { merge } = require("webpack-merge");

const { devConfig } = require("./webpack.config.common");
const { indexBaseConfig, mapsBaseConfig } = require("./webpack.config.server.base");
const { freeConfig } = require("./webpack.config.server.free");

const freeDevConfig = merge(devConfig, freeConfig);
const indexDevConfig = merge(devConfig, indexBaseConfig);
const mapsDevConfig = merge(devConfig, mapsBaseConfig);

module.exports = [freeDevConfig, indexDevConfig, mapsDevConfig];