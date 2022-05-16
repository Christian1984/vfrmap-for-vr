const path = require("path");
const { merge } = require("webpack-merge");
const HtmlWebpackPlugin = require("html-webpack-plugin");

const { freemiumBaseConfig, mapsBaseConfig, websrcBasePath } = require("./webpack.config.server.base");

const freemiumBasePath = path.resolve(websrcBasePath, "freemium");

const freemiumConfig = merge(freemiumBaseConfig, {
    entry: {
        charts: path.resolve(
            freemiumBasePath,
            "charts",
            "charts.js"
        ),
        notepad: path.resolve(
            freemiumBasePath,
            "notepad",
            "notepad.js"
        ),
        waypoints: path.resolve(
            freemiumBasePath,
            "waypoints",
            "waypoints.js"
        ),
    },
    plugins: [
        new HtmlWebpackPlugin({
            filename: "charts.html",
            inject: "body",
            template: path.resolve(
                freemiumBasePath,
                "template.html"
            ),
            chunks: ["charts"]
        }),
        new HtmlWebpackPlugin({
            filename: "notepad.html",
            inject: "body",
            template: path.resolve(
                freemiumBasePath,
                "template.html"
            ),
            chunks: ["notepad"]
        })
    ]
});

const mapsFreeConfig = merge(mapsBaseConfig, {
    entry: {
        maps: path.resolve(
            websrcBasePath,
            "maps",
            "maps.js"
        ),
    },
});

module.exports = { freemiumConfig, mapsFreeConfig };