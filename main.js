/* globals console */
'use strict';
var app = require('app');
var BrowserWindow = require('browser-window');

require('crash-reporter').start();

var mainWindow = null;
var backendPort = null;

var exec = require('child_process').exec;
var child = exec('./main');
child.stdout.on('data', function(data) {
  backendPort = data;
  console.log('go: ' + data);
});
child.stderr.on('data', function(data) {
  if (backendPort === null) {
    backendPort = data.split('Listening on')[1].trim();
  }
  console.log('go.err: ' + data);
});
child.on('close', function(code) {
  console.log('go stopped:' + code);
});

// Quit when all windows are closed.
app.on('window-all-closed', function() {
  if (process.platform !== 'darwin') {
    app.quit();
  }
});

app.on('will-quit', function() {
  child.kill();
});

app.on('ready', function() {
  mainWindow = new BrowserWindow({
    width: 1024,
    title: 'Food Plan Organizer',
    height: 786,
    'web-preferences': {
      'web-security': false
    }
  });

  var backendParam = '?backend=' + backendPort;
  mainWindow.loadUrl('http://localhost:9000/index.html' + backendParam);
  // mainWindow.loadUrl('file://' + __dirname + '/dist/index.html' + backendParam);

  mainWindow.on('closed', function() {
    mainWindow = null;
  });
});