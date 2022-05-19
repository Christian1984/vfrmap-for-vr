const path = require("path");
const { merge } = require("webpack-merge");
const FileManagerPlugin = require("filemanager-webpack-plugin");

const { panelBaseConfig, panelSrcPath, panelDistPath } = require("./webpack.config.panel.base");
const { devConfig } = require("./webpack.config.common");

module.exports = merge(devConfig, panelBaseConfig, {
    plugins: [
        new FileManagerPlugin({
            events: {
                onEnd: {
                    copy: [
                        {
                            source: path.resolve(panelSrcPath, "index.html"),
                            destination: path.resolve(panelDistPath, "index.html")
                        }
                    ]
                }
            }
        })
    ]
});