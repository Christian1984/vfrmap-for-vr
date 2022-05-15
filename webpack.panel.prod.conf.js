const path = require("path");
const { merge } = require("webpack-merge");

const { panelBaseConfig } = require("./webpack.panel.base.conf");
const { prodConfig } = require("./webpack.common.conf");

module.exports = merge(prodConfig, panelBaseConfig, {
    output: {
        filename: "[name].js",
        path: path.resolve( // TODO: this should be the community folder
            __dirname,
            "fskneeboard-panel",
            "christian1984-ingamepanel-fskneeboard",
            "html_ui",
            "InGamePanels",
            "FSKneeboardPanel"
        )
    },
});