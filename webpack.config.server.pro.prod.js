const { merge } = require("webpack-merge");

const { prodConfig } = require("./webpack.config.common");
const { indexBaseConfig, mapsBaseConfig } = require("./webpack.config.server.base");
const { proConfig } = require("./webpack.config.server.pro");

const proProdConfig = merge(prodConfig, proConfig);
const indexProdConfig = merge(prodConfig, indexBaseConfig);
const mapsProdConfig = merge(prodConfig, mapsBaseConfig);

module.exports = [proProdConfig, indexProdConfig, mapsProdConfig];