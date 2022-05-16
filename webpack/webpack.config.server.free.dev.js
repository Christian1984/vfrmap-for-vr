const { merge } = require("webpack-merge");

const { devConfig } = require("./webpack.config.common");
const { indexBaseConfig } = require("./webpack.config.server.base");
const { freemiumConfig, mapsFreeConfig } = require("./webpack.config.server.free");

const freeDevConfig = merge(devConfig, freemiumConfig);
const indexDevConfig = merge(devConfig, indexBaseConfig);
const mapsDevConfig = merge(devConfig, mapsFreeConfig);

module.exports = [freeDevConfig, indexDevConfig, mapsDevConfig];