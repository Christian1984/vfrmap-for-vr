const path = require("path");
const { merge } = require("webpack-merge");
const HtmlWebpackPlugin = require("html-webpack-plugin");

const commonConfig = require("./webpack.common.conf");

module.exports = merge(commonConfig, {
    mode: "production",
    devtool: "source-map",
    entry: {
        index: path.resolve(
            __dirname,
            "fskneeboard-server",
            "vfrmap",
            "html",
            "web",
            "src",
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
            "web",
            "dist"
        )
    },
    plugins: [
        new HtmlWebpackPlugin({
            filename: "[name].html",
            inject: "body",
            template: path.resolve(
                __dirname,
                "fskneeboard-server",
                "vfrmap",
                "html",
                "web",
                "src",
                "index.html"
            )
        })
    ]
});