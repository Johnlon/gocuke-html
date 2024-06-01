let inputOnHTML = document.getElementById("filesToUpload");
let uploadInfo = document.getElementById("uploadInfo");
let uploadContent = document.getElementById("uploadContent");

inputOnHTML.onchange = function () {
  uploadContent.innerHTML = "";
  for (let v of inputOnHTML.files) {
    uploadInfo.style.opacity = "1";
    uploadContent.innerHTML += v.name + "<br>";
  }
};

["dragenter", "dragover"].forEach((eventName) => {
  inputOnHTML.addEventListener(eventName, highlight, false);
});

["dragleave", "drop"].forEach((eventName) => {
  inputOnHTML.addEventListener(eventName, unhighlight, false);
});

function highlight(e) {
  inputOnHTML.classList.add("highlight-multiple");
}

function unhighlight(e) {
  inputOnHTML.classList.remove("highlight-multiple");
}
