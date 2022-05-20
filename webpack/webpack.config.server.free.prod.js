const path = require("path");
const { merge } = require("webpack-merge");

const { prodConfig } = require("./webpack.config.common");
const { indexBaseConfig } = require("./webpack.config.server.base");
const { freemiumConfig, mapsFreeConfig } = require("./webpack.config.server.free");

const freeProdConfig = merge(prodConfig, freemiumConfig);
const indexProdConfig = merge(prodConfig, indexBaseConfig);
const mapsProdConfig = merge(prodConfig, mapsFreeConfig);

module.exports = [freeProdConfig, indexProdConfig, mapsProdConfig];