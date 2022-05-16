const { merge } = require("webpack-merge");

const { prodConfig } = require("./webpack.config.common");
const { indexBaseConfig } = require("./webpack.config.server.base");
const { premiumConfig, mapsProConfig } = require("./webpack.config.server.pro");

const proProdConfig = merge(prodConfig, premiumConfig);
const indexProdConfig = merge(prodConfig, indexBaseConfig);
const mapsProdConfig = merge(prodConfig, mapsProConfig);

module.exports = [proProdConfig, indexProdConfig, mapsProdConfig];