const path = require("path");
const { merge } = require("webpack-merge");
const HtmlWebpackPlugin = require("html-webpack-plugin");

const { freemiumBaseConfig } = require("./webpack.config.server.base");

const premiumBasePath = path.resolve(__dirname, "fskneeboard-server", "_vendor", "premium_src", "websrc");

const proConfig = merge(freemiumBaseConfig, {
    entry: {
        charts: path.resolve(
            premiumBasePath,
            "charts",
            "charts.js"
        ),
        notepad: path.resolve(
            premiumBasePath,
            "notepad",
            "notepad.js"
        ),
        waypoints: path.resolve(
            premiumBasePath,
            "waypoints",
            "waypoints.js"
        ),
    },
    plugins: [
        new HtmlWebpackPlugin({
            filename: "charts.html",
            inject: "body",
            template: path.resolve(
                premiumBasePath,
                "charts",
                "charts.html"
            ),
            chunks: ["charts"]
        }),
        new HtmlWebpackPlugin({
            filename: "notepad.html",
            inject: "body",
            template: path.resolve(
                premiumBasePath,
                "notepad",
                "notepad.html"
            ),
            chunks: ["notepad"]
        })
    ]
});

module.exports = { proConfig };