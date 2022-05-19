const path = require("path");
const { merge } = require("webpack-merge");
const HtmlWebpackPlugin = require("html-webpack-plugin");

const { commonConfig } = require("./webpack.config.common");
const panelDistPath = path.resolve(
    __dirname, "..",
    "fskneeboard-panel",
    "christian1984-ingamepanel-fskneeboard",
    "html_ui",
    "InGamePanels",
    "FSKneeboardPanel"
);

const panelBaseConfig = merge(commonConfig, {
    entry: {
        FSKneeboardPanel: path.resolve(__dirname, "..", "fskneeboard-panel", "src", "FSKneeboardPanel.js"),
    },
    output: {
        filename: "[name].js",
        path: panelDistPath,
        clean: true
    },
    plugins: [
        new HtmlWebpackPlugin({
            filename: "FSKneeboardPanel.html",
            inject: "head",
            template: path.resolve(__dirname, "..", "fskneeboard-panel", "src", "FSKneeboardPanel.html"),
            chunks: ["FSKneeboardPanel"]
        })
    ]
});

module.exports = { panelBaseConfig, panelDistPath }