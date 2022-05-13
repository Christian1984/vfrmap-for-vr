const path = require("path");
const { merge } = require("webpack-merge");
const HtmlWebpackPlugin = require("html-webpack-plugin");

const commonConfig = require("./webpack.common.conf");

module.exports = merge(commonConfig, {
    mode: "production",
    devtool: "source-map",
    entry: {
        FSKneeboardPanel: path.resolve(__dirname, "fskneeboard-panel", "src", "FSKneeboardPanel.js"),
    },
    output: {
        filename: "[name].js",
        path: path.resolve(
            __dirname,
            "fskneeboard-panel",
            "christian1984-ingamepanel-fskneeboard",
            "html_ui",
            "InGamePanels",
            "FSKneeboardPanel"
        )
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