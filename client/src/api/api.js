'use strict';

var MinesStore = require('../../src/stores/minestore.js');
var AppDispatcher = require('../dispatchers/app_dispatcher.js');

var _ = require('../../src/utility/utils.js');

var apiUrl = "ws://192.168.0.103:7000";

/* websocket reference */
var conn = null;

module.exports = {
  connect: function(nickname) {
    console.log('inside connect');
    if (window["WebSocket"]) {
      conn = new WebSocket(apiUrl + '/players/' + nickname);

      conn.onclose = function(evt) {
        console.log('connection closed.');
      };

      conn.onmessage = function(evt) {
        var data =  _.getObject(evt.data);
      
        //console.log('data', data);
        AppDispatcher.handleServerAction(data);
      };
    } else {
      console.log($("Your browser does not support WebSockets."));
    }
  },

  send: function(data) {
    console.log('sending ', data);
    conn.send(data);
  }
};
