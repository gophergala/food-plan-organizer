var app = require('app'); // Module to control application life.
var BrowserWindow = require('browser-window'); // Module to create native browser window.

// Report crashes to our server.
require('crash-reporter').start();

// Keep a global reference of the window object, if you don't, the window will
// be closed automatically when the javascript object is GCed.
var mainWindow = null;

// var exec = require('child_process').exec;
// var child = exec('./main');
// child.stdout.on('data', function(data) {
//   console.log('go: ' + data);
// });
// child.stderr.on('data', function(data) {
//   console.log('go.err: ' + data);
// });
// child.on('close', function(code) {
//   console.log('go stopped:' + code);
// });

// // Quit when all windows are closed.
// app.on('window-all-closed', function() {
//   if (process.platform != 'darwin') {
//     app.quit();
//   }
// });

// app.on('will-quit', function() {
//   child.kill();
// });

// This method will be called when atom-shell has done everything
// initialization and ready for creating browser windows.
app.on('ready', function() {
  // Create the browser window.
  mainWindow = new BrowserWindow({
    width: 1024,
    height: 786,
    'web-preferences': {
      'web-security': false
    }
  });

  // and load the index.html of the app.
  mainWindow.loadUrl('http://localhost:9000/index.html');
  // mainWindow.loadUrl('file://' + __dirname + '/dist/index.html');

  // Emitted when the window is closed.
  mainWindow.on('closed', function() {
    // Dereference the window object, usually you would store windows
    // in an array if your app supports multi windows, this is the time
    // when you should delete the corresponding element.
    mainWindow = null;
  });
});