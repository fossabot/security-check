//------------------------------------------------------------------------------
//
// Webpack Configuration, https://webpack.js.org/concepts/configuration
//
//------------------------------------------------------------------------------

const HtmlWebpackPlugin = require("html-webpack-plugin");

module.exports = {
  entry: `${__dirname}/client/index.ts`,
  resolve: {
    extensions: [".ts", ".tsx", ".js"]
  },
  output: {
    path: `${__dirname}/dist/`,
    filename: "all.min.js"
  },
  devServer: {
    historyApiFallback: true
  },
  stats: "errors-only",
  module: {
    rules: [
      {
        test: /\.tsx?$/,
        use: "ts-loader",
        exclude: /node_modules/
      }
    ]
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: `${__dirname}/client/public/index.html`
    })
  ]
};
