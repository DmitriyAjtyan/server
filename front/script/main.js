"use strict";

var socket = new WebSocket('ws://127.0.0.1/api/second');

function alertShow() {
    alert(1);
};

socket.onopen = function() {
    console.log("Socket connected!");
  };
  
  socket.onclose = function(event) {
    if (event.wasClean) {
      console.log('Connection closed cleanly');
    } else {
      console.log('Connection break');
    }
    console.log('Code: ' + event.code + ' reason: ' + event.reason);
  };
  
  socket.onmessage = function(event) {
    console.log("Get data " + event.data);
  };
  
  socket.onerror = function(error) {
    console.log("ERROR " + error.message);
  };
