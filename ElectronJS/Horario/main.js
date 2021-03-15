const electron = require("electron");
const app = electron.app;
var path = require("path");

const BrowserWindows = electron.BrowserWindow;

app.on("ready", function () {
  let win = new BrowserWindows({
    show: false,
    minHeight: 800,
    minWidth: 1200,
    backgroundColor: '#DFE0E2',
    webPreferences: {
      //Especifica que va a cargar primero antes que todos los scripts de la pagina
      //path.join() junta varias cosas y me las pone en un path
      preload: path.join(__dirname, "preload.js"),
    },
  });

  //Load a local HTML file
  win.loadURL("file://" + __dirname + "/frontend/index.html");

  win.once("ready-to-show", function () {
    win.show();
  });
});

app.on("window-all-closed", function () {
  if (process.plataform != "darwin") app.quit();
});
