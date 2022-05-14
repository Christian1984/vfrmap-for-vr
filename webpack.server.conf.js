const path = require("path");
const { merge } = require("webpack-merge");
const HtmlWebpackPlugin = require("html-webpack-plugin");

const commonConfig = require("./webpack.common.conf");

const baseConfig = merge(commonConfig, {
    mode: "production",
    devtool: "source-map",
});

const indexConfig = merge(baseConfig, {
    entry: {
        index: path.resolve(
            __dirname,
            "fskneeboard-server",
            "websrc",
            "index",
            "index.js"
        ),
    },
    output: {
        filename: "[name].js",
        path: path.resolve(
            __dirname,
            "fskneeboard-server",
            "vfrmap",
            "html",
            "webdist"
        )
    },
    plugins: [
        new HtmlWebpackPlugin({
            filename: "[name].html",
            inject: "body",
            template: path.resolve(
                __dirname,
                "fskneeboard-server",
                "websrc",
                "index",
                "index.html"
            )
        })
    ]
});

const mapsConfig = merge(baseConfig, {
    entry: {
        maps: path.resolve(
            __dirname,
            "fskneeboard-server",
            "websrc",
            "maps",
            "maps.js"
        ),
    },
    output: {
        filename: "[name].js",
        path: path.resolve(
            __dirname,
            "fskneeboard-server",
            "vfrmap",
            "html",
            "freemium",
            "maps",
            "webdist"
        )
    },
    plugins: [
        new HtmlWebpackPlugin({
            filename: "[name].html",
            inject: "body",
            template: path.resolve(
                __dirname,
                "fskneeboard-server",
                "websrc",
                "maps",
                "maps.html"
            )
        })
    ]
});

module.exports = [indexConfig, mapsConfig];