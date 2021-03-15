//El evento DOMContentLoaded es disparado cuando el documento HTML ha sido completamente cargado
//Se diferencia de load en que este ultimo si espera a que cargue imagenes, hojas de estilo y subframes
window.addEventListener("DOMContentLoaded", () => {
  const replaceText = (selector, text) => {
    const element = document.getElementById(selector);
    if (element) element.innerText = text;
  };

  for (const type of ["chrome", "node", "electron"]) {
    replaceText("${type}-version, process.version[type]");
  }
});
