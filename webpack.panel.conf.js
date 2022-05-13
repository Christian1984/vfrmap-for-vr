const path = require("path");
const { CleanWebpackPlugin } = require("clean-webpack-plugin");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const TerserPlugin = require("terser-webpack-plugin");

module.exports = {
    mode: "production",
    //devtool: "source-map",
    entry: {
        FSKneeboardPanel: path.resolve(__dirname, "fskneeboard-panel", "src", "FSKneeboardPanel.js"),
    },
    output: {
        filename: "[name].js",
        path: path.resolve(
            __dirname,
            "fskneeboard-panel",
            "christian1984-ingamepanel-fskneeboard",
            "html_ui",
            "InGamePanels",
            "FSKneeboardPanel"
        )
    },
    plugins: [
        new CleanWebpackPlugin(),
        new HtmlWebpackPlugin({
            filename: "[name].html",
            inject: "head",
            template: path.resolve(__dirname, "fskneeboard-panel", "src", "FSKneeboardPanel.html"),
            //chunks: ["FSKneeboardPanel"]
        }),
        new MiniCssExtractPlugin({
            filename: '[name].css'
        }),
    ],
    module: {
        rules: [
            /*
            {
                test: /\.(js|jsx)$/,
                exclude: /[\\/]node_modules[\\/]/,
                use: ["babel-loader"]
            },
            */
            {
                test: /\.s[ac]ss$/,
                use: [MiniCssExtractPlugin.loader, "css-loader", "sass-loader"]
            }
        ]
    },
    optimization: {
        /*minimize: true,
        minimizer: [
            new TerserPlugin({
                parallel: true,
                terserOptions: {
                    mangle: true,
                }
            }),
        ],*/
    },
}