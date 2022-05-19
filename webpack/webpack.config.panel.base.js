const path = require("path");
const { merge } = require("webpack-merge");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const FileManagerPlugin = require("filemanager-webpack-plugin");


const { commonConfig } = require("./webpack.config.common");

const panelSrcPath = path.resolve(__dirname, "..", "fskneeboard-panel", "src");

const panelDistPath = path.resolve(
    __dirname, "..",
    "fskneeboard-panel",
    "christian1984-ingamepanel-fskneeboard",
    "html_ui",
    "InGamePanels",
    "FSKneeboardPanel"
);

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

const panelBaseConfig = merge(commonConfig, {
    entry: {
        FSKneeboardPanel: path.resolve(panelSrcPath, "FSKneeboardPanel.js"),
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
            template: path.resolve(panelSrcPath, "FSKneeboardPanel.html"),
            chunks: ["FSKneeboardPanel"]
        }),
        new FileManagerPlugin({
            events: {
                onStart: {
                    delete: [
                        {
                          source: communityFolderPanelPath,
                          options: {
                            force: true,
                          },
                        },
                    ]
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

module.exports = { panelBaseConfig, panelSrcPath, panelDistPath }