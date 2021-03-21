const electron = require("electron");
const app = electron.app;
const BrowserWindow = electron.BrowserWindow;

function crearWin() {
  let win = new BrowserWindow({
    width: 450,
    height: 650,
    resizable: false,
    webPreferences: {
      nodeIntegration: true,
    },
  });
  win.loadFile( __dirname + "/frontend/html/calculadora.html");
}

app.whenReady().then(crearWin);

app.on("window-all-closed", () => {
  if (process.platform != "darwin") {
    app.quit();
  }
});

app.on("activate", () => {
  if (BrowserWindow.getAllWindows().length === 0) {
    crearWin();
  }
});
