const path = require("path");
const { merge } = require("webpack-merge");
const HtmlWebpackPlugin = require("html-webpack-plugin");

const commonConfig = require("./webpack.common.conf");

module.exports = merge(commonConfig, {
    mode: "production",
    devtool: "source-map",
    entry: {
        root: path.resolve(
            __dirname,
            "fskneeboard-server",
            "vfrmap",
            "html",
            "websrc",
            "index.js"
        ),
    },
    output: {
        filename: "[name].js",
        path: path.resolve(
            __dirname,
            "fskneeboard-server",
            "vfrmap",
            "html"
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
                "websrc",
                "index.html"
            ),
            chunks: ["root"]
        })
    ]
});