const path = require("path");
const { merge } = require("webpack-merge");

const { prodConfig } = require("./webpack.config.common");
const { indexBaseConfig, mapsBaseConfig } = require("./webpack.config.server.base");
const { freeConfig } = require("./webpack.config.server.free");

const freeProdConfig = merge(prodConfig, freeConfig);
const indexProdConfig = merge(prodConfig, indexBaseConfig);
const mapsProdConfig = merge(prodConfig, mapsBaseConfig);

module.exports = [freeProdConfig, indexProdConfig, mapsProdConfig];