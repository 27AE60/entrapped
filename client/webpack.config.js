'use strict'

var src = '/src',
  dist = '/dist';

module.exports = {
  context: __dirname + '/src',
  entry: [
    './app.jsx',
    './templates/index.html'
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
        loader: 'babel-loader'
      },
      {
        test: /\.html$/,
        loader: "file?name=[name].[ext]",
      },
    ]
  }
};
