const electron = require("electron");
const app = electron.app;
var path = require("path");

const BrowserWindows = electron.BrowserWindow;

app.on("ready", function () {
  let win = new BrowserWindows({
    show: false,
    minHeight: 800,
    minWidth: 1200,
    backgroundColor: '#364958',

    //Para quitar cosas del menu
    //frame: false,
    //titleBarStyle: "hidden",

    webPreferences: {
      //Especifica que va a cargar primero antes que todos los scripts de la pagina
      //path.join() junta varias cosas y me las pone en un path
      preload: path.join(__dirname, "preload.js"),
      //Previene el error de integración de HTML/CSS con NodeJS
      nodeIntegration: true}
  });

  //win.maximize();
  
  //Load a local HTML file
  win.loadURL("file://" + __dirname + "/frontend/html/index.html");

  //Cuando este lista la pagina, se mostrara lo que tenga aca
  win.once("ready-to-show", function () {
    win.show();
  });
});

//Cuando cierre todo
app.on("window-all-closed", function () {
  if (process.plataform != "darwin") app.quit();
});
