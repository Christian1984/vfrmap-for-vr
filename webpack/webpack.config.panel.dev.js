const path = require("path");
const { merge } = require("webpack-merge");

const { panelBaseConfig } = require("./webpack.config.panel.base");
const { devConfig } = require("./webpack.config.common");

module.exports = merge(devConfig, panelBaseConfig, {
    output: {
        filename: "[name].js",
        path: path.resolve(
            process.env.LOCALAPPDATA,
            "Packages",
            "Microsoft.FlightSimulator_8wekyb3d8bbwe",
            "LocalCache",
            "Packages",
            "Community",
            "christian1984-ingamepanel-fskneeboard",
            "html_ui",
            "InGamePanels",
            "FSKneeboardPanel"
        ),
        clean: true
    }
});