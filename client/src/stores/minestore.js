var AppDispatcher = require('../dispatchers/app_dispatcher.js');
var assign = require('object-assign');
var EventEmitter = require('events').EventEmitter;
var AppConstants = require('../constants/app_constants.js');

/*
 * Minesfield.
 *
 * Status:
 *
 *     0 : default, not clicked,
 *    -1 : visited,
 *     2 : life,
 *    -2 : death
 *
 * Size: 7 by 7
 *
 * Mines position (x, y) from array can be calculated by,:
 *    x : index of an element / size of minesfield
 *    y : index of an element % size of minesfield
 */

var mines = [
  0, -1, 0, 0, 0, 0, 0,
  0, 0, 0, 2, 0, 0, -1,
  0, -2, 0, 0, 0, 0, 0,
  0, 0, 0, -1, 0, 0, 0,
  0, 0, 0, -1, 0, 0, 0,
  0, 0, 0, 0, 0, 0, -1,
  0, 0, 0, 0, 0, 0, -1,
];

/* username */
var username = '';

var MinesStore = assign({}, EventEmitter.prototype, {
  emitChange: function() {
    this.emit(CHANGE_EVENT);
  },

  addChangeListener: function(callback) {
    this.on(CHANGE_EVENT, callback);
  },

  removeChangeListener: function(callback) {
    this.removeListener(CHANGE_EVENT, callback);
  },

  'getMines': function() {
    return mines;
  },

  'addUsername': function(name) {
    username = name;
  },

  'getUsername': function() {
    return username;
  }
});

MinesStore.dispatchToken = AppDispatcher.register(function(payload) {
  var action = payload.action;

  switch(action.type) {
    case 'registered':
      console.log('user registered');
      break;

    case 'ready':
      console.log('ready');
      break;

    case 'open':
      console.log('open');
      break;

    case 'result':
      console.log('result');
      break;

    case 'enemy':
      console.log('enemy');
      break;

    default:
      // do nothing
  }
});

module.exports = MinesStore;
