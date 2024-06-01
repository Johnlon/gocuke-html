function hideLoader() {
    document.getElementById("loading").style.visibility = "visibility";
    document.getElementById("loading").style.opacity = "0";
    document.getElementById("loading").style.transition = "visibility 0s 1s, opacity 1s linear";
    document.getElementById("loading").style.zIndex = "0";
    document.getElementById("loading").style.display = "none";
  }
  if (window.addEventListener) {
    window.addEventListener("load", hideLoader);
  } else if (window.attachEvent) {
    window.attachEvent("onload", hideLoader);
  } else {
    window.onload = hideLoader;
  }

function showLoading(){
    document.getElementById("loading").style.visibility = "visible";
    document.getElementById("loading").style.opacity = "1";
    document.getElementById("loading").style.transition = "all 1s linear";
    document.getElementById("loading").style.zIndex = "99999999999";
    document.getElementById("loading").style.display = "block";
    return true;
}