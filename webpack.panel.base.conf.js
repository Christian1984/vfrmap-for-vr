const path = require("path");
const { merge } = require("webpack-merge");
const HtmlWebpackPlugin = require("html-webpack-plugin");

const { commonConfig } = require("./webpack.common.conf");

const panelBaseConfig = merge(commonConfig, {
    entry: {
        FSKneeboardPanel: path.resolve(__dirname, "fskneeboard-panel", "src", "FSKneeboardPanel.js"),
    },
    plugins: [
        new HtmlWebpackPlugin({
            filename: "[name].html",
            inject: "head",
            template: path.resolve(__dirname, "fskneeboard-panel", "src", "FSKneeboardPanel.html"),
            chunks: ["FSKneeboardPanel"]
        })
    ]
});

module.exports = { panelBaseConfig }