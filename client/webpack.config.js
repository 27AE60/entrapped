'use strict'

var webpack = require('webpack');

module.exports = {
  entry: [
    'webpack/hot/only-dev-server',
    './src/app.jsx',
    './src/templates/index.html',
    './src/styles/main.css',
    './src/public/font/1942.ttf'
  ],

  output: {
    filename: 'app.js',
    path: __dirname + '/dist'
  },

  debug : false,

  module: {
    loaders: [
      {
        test: /\.jsx?$/,
        exclude: /node_modules/,
        loaders: ['react-hot', 'babel']
      },
      {
        test: /\.js$/,
        exclude: /node_modules/,
        loader: 'babel-loader'
      },
      {
        test: /\.html$/,
        loader: "file?name=[name].[ext]",
      },
      {
        test: /\.css$/,
        loader: "file?name=[name].[ext]",
      },
      {
        test: /\.ttf$/,
        loader: "file?name=[name].[ext]",
      }
    ]
  },

  plugins: [
    new webpack.HotModuleReplacementPlugin(),
    new webpack.NoErrorsPlugin()
  ]
};
