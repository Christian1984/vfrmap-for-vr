const path = require("path");
const { merge } = require("webpack-merge");

const { panelBaseConfig } = require("./webpack.panel.base.conf");
const { devConfig } = require("./webpack.common.conf");

module.exports = merge(devConfig, panelBaseConfig, {
    output: {
        filename: "[name].js",
        path: path.resolve(
            __dirname,
            "fskneeboard-panel",
            "christian1984-ingamepanel-fskneeboard",
            "html_ui",
            "InGamePanels",
            "FSKneeboardPanel",
            "dev"
        )
    }
});