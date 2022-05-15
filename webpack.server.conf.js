const path = require("path");
const { merge } = require("webpack-merge");
const HtmlWebpackPlugin = require("html-webpack-plugin");

const { commonConfig } = require("./webpack.common.conf");

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
            filename: "[name].html",
            inject: "body",
            template: path.resolve(
                freemiumBasePath,
                "template.html"
            )
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
            filename: "[name].html",
            inject: "body",
            template: path.resolve(
                websrcBasePath,
                "index",
                "index.html"
            ),
            chunks: ["index"]
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
            filename: "[name].html",
            inject: "body",
            template: path.resolve(
                websrcBasePath,
                "maps",
                "maps.html"
            ),
            chunks: ["maps"]
        })
    ]
});

module.exports = [freemiumConfig, indexConfig, mapsConfig];