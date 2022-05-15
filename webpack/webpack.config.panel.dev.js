const path = require("path");
const { merge } = require("webpack-merge");

const { panelBaseConfig } = require("./webpack.config.panel.base");
const { devConfig } = require("./webpack.config.common");

module.exports = merge(devConfig, panelBaseConfig, {
    output: {
        filename: "[name].js",
        path: path.resolve(
            __dirname, "..",
            "fskneeboard-panel",
            "christian1984-ingamepanel-fskneeboard",
            "html_ui",
            "InGamePanels",
            "FSKneeboardPanel",
            "dev"
        ),
        clean: true
    }
});