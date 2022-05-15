const path = require("path");
const { merge } = require("webpack-merge");
const HtmlWebpackPlugin = require("html-webpack-plugin");

const { commonConfig } = require("./webpack.config.common");

const baseConfig = merge(commonConfig, {
    mode: "production",
    devtool: "source-map",
});

const websrcBasePath = path.resolve(__dirname, "fskneeboard-server", "websrc");
const freemiumBasePath = path.resolve(websrcBasePath, "freemium");

const freemiumConfig = merge(baseConfig, {
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
    output: {
        filename: "[name].js",
        path: path.resolve(
            __dirname,
            "fskneeboard-server",
            "_vendor",
            "premium",
            "webdist"
        )
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

const indexConfig = merge(baseConfig, {
    entry: {
        index: path.resolve(
            websrcBasePath,
            "index",
            "index.js"
        ),
    },
    output: {
        filename: "[name].js",
        path: path.resolve(
            __dirname,
            "fskneeboard-server",
            "vfrmap",
            "html",
            "webdist"
        )
    },
    plugins: [
        new HtmlWebpackPlugin({
            filename: "index.html",
            inject: "body",
            template: path.resolve(
                websrcBasePath,
                "index",
                "index.html"
            )
        })
    ]
});

const mapsConfig = merge(baseConfig, {
    entry: {
        maps: path.resolve(
            websrcBasePath,
            "maps",
            "maps.js"
        ),
    },
    output: {
        filename: "[name].js",
        path: path.resolve(
            __dirname,
            "fskneeboard-server",
            "vfrmap",
            "html",
            "freemium",
            "maps",
            "webdist"
        )
    },
    plugins: [
        new HtmlWebpackPlugin({
            filename: "maps.html",
            inject: "body",
            template: path.resolve(
                websrcBasePath,
                "maps",
                "maps.html"
            )
        })
    ]
});

module.exports = [freemiumConfig, indexConfig, mapsConfig];