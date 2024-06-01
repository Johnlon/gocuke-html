let inputFile = document.getElementById("fileToUpload");
let uploadInfoEmbed = document.getElementById("uploadInfoEmbed");
let uploadContentEmbed = document.getElementById("uploadContentEmbed");

inputFile.onchange = function () {
  uploadContentEmbed.innerHTML = "";
  uploadInfoEmbed.style.opacity = "1";
  uploadContentEmbed.innerHTML += inputFile.value.replace('C:\\fakepath\\', ' ') + "<br>";
};

let inputFiles = document.getElementById("filesToUpload");
let uploadInfoEmbedMultiple = document.getElementById("uploadInfoEmbedMultiple");
let uploadContentEmbedMultiple = document.getElementById("uploadContentEmbedMultiple");

inputFiles.onchange = function () {
  uploadContentEmbedMultiple.innerHTML = "";
  for (let v of inputFiles.files) {
    uploadInfoEmbedMultiple.style.opacity = "1";
    uploadContentEmbedMultiple.innerHTML += v.name + "<br>";
  }
};

;['dragenter', 'dragover'].forEach(eventName => {
  inputFile.addEventListener(eventName, highlightSingle, false)
  inputFiles.addEventListener(eventName, highlightMultiple, false)
})

;['dragleave', 'drop'].forEach(eventName => {
  inputFile.addEventListener(eventName, unhighlightSingle, false)
  inputFiles.addEventListener(eventName, unhighlightMultiple, false)
})

function highlightSingle(e) {
  inputFile.classList.add('highlight-single')
}

function unhighlightSingle(e) {
  inputFile.classList.remove('highlight-single')
}

function highlightMultiple(e) {
  inputFiles.classList.add('highlight-multiple')
}

function unhighlightMultiple(e) {
  inputFiles.classList.remove('highlight-multiple')
}
