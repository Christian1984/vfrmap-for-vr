const path = require("path");
const { merge } = require("webpack-merge");
const HtmlWebpackPlugin = require("html-webpack-plugin");

const { commonConfig } = require("./webpack.config.common");

const websrcBasePath = path.resolve(__dirname, "..", "fskneeboard-server", "websrc");

const freemiumBaseConfig = merge(commonConfig, {
    output: {
        filename: "[name].js",
        path: path.resolve(
            __dirname, "..",
            "fskneeboard-server",
            "_vendor",
            "premium",
            "webdist"
        ),
        clean: true
    }
});

const indexBaseConfig = merge(commonConfig, {
    entry: {
        index: path.resolve(
            websrcBasePath,
            "index",
            "index.js"
        ),
    },
    output: {
        filename: "[name].js",
        path: path.resolve(
            __dirname, "..",
            "fskneeboard-server",
            "vfrmap",
            "html",
            "webdist"
        ),
        clean: true
    },
    plugins: [
        new HtmlWebpackPlugin({
            filename: "index.html",
            inject: "body",
            template: path.resolve(
                websrcBasePath,
                "index",
                "index.html"
            )
        })
    ]
});

const mapsBaseConfig = merge(commonConfig, {
    output: {
        filename: "[name].js",
        path: path.resolve(
            __dirname, "..",
            "fskneeboard-server",
            "vfrmap",
            "html",
            "freemium",
            "maps",
            "webdist"
        ),
        clean: true
    },
    plugins: [
        new HtmlWebpackPlugin({
            filename: "maps.html",
            inject: "body",
            template: path.resolve(
                websrcBasePath,
                "maps",
                "maps.html"
            )
        })
    ]
});

module.exports = { freemiumBaseConfig, indexBaseConfig, mapsBaseConfig, websrcBasePath };