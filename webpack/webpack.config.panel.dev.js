const path = require("path");
const { merge } = require("webpack-merge");

const { panelBaseConfig, panelDistPath } = require("./webpack.config.panel.base");
const { devConfig } = require("./webpack.config.common");

const communityFolderPanelPath = path.resolve(
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
);

module.exports = merge(devConfig, panelBaseConfig, {
    plugins: [
        new FileManagerPlugin({
            events: {
                onStart: {
                    delete: [communityFolderPanelPath],
                },
                onEnd: {
                    copy: [
                        {
                            source: panelDistPath,
                            destination: communityFolderPanelPath
                        }
                    ]
                }
            }
        })
    ]
});