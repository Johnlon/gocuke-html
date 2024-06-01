function showHide(eleName) {
  let ele = document.getElementById(eleName);
  if (ele.classList.contains("hide")) {
    ele.classList.remove("hide");
    document.querySelector("#" + eleName + "-header .expand-arrow").innerHTML =
      "expand_less";
  } else {
    ele.classList.add("hide");
    document.querySelector("#" + eleName + "-header .expand-arrow").innerHTML =
      "expand_more";
  }
}

filterSelection("all");
function filterSelection(c) {
  let x, i;
  x = document.getElementsByClassName("feature-card");
  if (c == "all") c = "";
  for (i = 0; i < x.length; i++) {
    w3RemoveClass(x[i], "show");
    if (x[i].className.indexOf(c) > -1) w3AddClass(x[i], "show");
  }
}

function w3AddClass(element, name) {
  let i, arr1, arr2;
  arr1 = element.className.split(" ");
  arr2 = name.split(" ");
  for (i = 0; i < arr2.length; i++) {
    if (arr1.indexOf(arr2[i]) == -1) {
      element.className += " " + arr2[i];
    }
  }
}

function w3RemoveClass(element, name) {
  let i, arr1, arr2;
  arr1 = element.className.split(" ");
  arr2 = name.split(" ");
  for (i = 0; i < arr2.length; i++) {
    while (arr1.indexOf(arr2[i]) > -1) {
      arr1.splice(arr1.indexOf(arr2[i]), 1);
    }
  }
  element.className = arr1.join(" ");
}

function hideLoader() {
  document.getElementById("loading").style.visibility = "visibility";
  document.getElementById("loading").style.opacity = "0";
  document.getElementById("loading").style.transition =
    "visibility 0s 1s, opacity 1s linear";
  document.getElementById("loading").style.zIndex = "0";
  setTimeout(function () {
    document.getElementById("loading").style.display = "none";
  }, 1800);
}
if (window.addEventListener) {
  window.addEventListener("load", hideLoader);
} else if (window.attachEvent) {
  window.attachEvent("onload", hideLoader);
} else {
  window.onload = hideLoader;
}

let switchPassed = document.getElementById("switch-passed");
let switchFailed = document.getElementById("switch-failed");
switchPassed.addEventListener("change", (e) => {
  checkSwitchFilterState(switchPassed, switchFailed);
});
switchFailed.addEventListener("change", (e) => {
  checkSwitchFilterState(switchPassed, switchFailed);
});

function checkSwitchFilterState(switchPassed, switchFailed) {
  if (switchPassed.checked && switchFailed.checked) {
    filterSelection("all");
  } else if (switchPassed.checked && !switchFailed.checked) {
    filterSelection("passed");
  } else if (!switchPassed.checked && switchFailed.checked) {
    filterSelection("failed");
  } else {
    filterSelection("none");
  }
}

let chartDisplay = "";
let switchCharts = document.getElementById("switch-charts");
switchCharts.addEventListener("change", (e) => {
  if (switchCharts.checked) {
    chartDisplay = "grid";
  } else {
    chartDisplay = "none";
  }

  let chartCards = document.getElementsByClassName("chart-card");
  for (let i = 0; i < chartCards.length; i++) {
    chartCards[i].style.display = chartDisplay;
  }
});

let metadataDisplay = "";
let switchMetadata = document.getElementById("switch-metadata");
switchMetadata.addEventListener("change", (e) => {
  if (switchMetadata.checked) {
    metadataDisplay = "block";
  } else {
    metadataDisplay = "none";
  }

  document.getElementById("metadata-card").style.display = metadataDisplay;
});

let featureCardWidth = "";
let featureCardMargin = "";
let switchTwoColumns = document.getElementById("switch-two-columns");
switchTwoColumns.addEventListener("change", (e) => {
  if (switchTwoColumns.checked) {
    featureCardWidth = "46.4%";
  } else {
    featureCardWidth = "96.4%";
  }

  let featureCards = document.getElementsByClassName("feature-card");
  for (let i = 0; i < featureCards.length; i++) {
    featureCards[i].style.width = featureCardWidth;
    featureCards[i].style.margin = featureCardMargin;
  }
});

setTimeout(hideLoader, 30 * 1000);

document.onclick = function (e) {
  if (e.target.id !== "sidebar" && e.target.id !== "btn" && e.target.id !== "cancle") {
    if (e.target.offsetParent) {
      if (e.target.offsetParent.id !== "sidebar") {
        closeSidebar();
      }
    } else {
      closeSidebar();
    }
  }
};

function closeSidebar() {
  let ele = document.getElementById("sidebar");
  let left = window.getComputedStyle(ele).getPropertyValue("left");
  if (left === "0px") {
    document.getElementById("cancle").click();
  }
}

function printDiv(ele) {
    let eleId = ele.getAttribute("data-id")
    let divContents = document.getElementById(eleId).innerHTML;
    let a = window.open('', '', 'height400, width=600');
    a.document.write('<html><body>');
    a.document.write('<style>pre{white-space: -moz-pre-wrap;white-space: -pre-wrap;white-space: -o-pre-wrap;white-space: pre-wrap;word-wrap: break-word;}</style>');
    a.document.write(divContents);
    a.document.write('</body></html>');
    a.document.close();
    a.print();
}

// Copy Script
"use strict";const copyjs=function(e,t={reSelect:!0,html:!1,copyFromSelector:!1}){const n=new class{init(){return this.selectionOptions={start:{index:window.getSelection().anchorOffset,node:window.getSelection().anchorNode},end:{index:window.getSelection().focusOffset,node:window.getSelection().focusNode}},!0}reselect(){if(!this.selectionOptions.start.node)return!1;const e=window.getSelection();e.removeAllRanges();const t=document.createRange();return t.setStart(this.selectionOptions.start.node,this.selectionOptions.start.index),t.setEnd(this.selectionOptions.end.node,this.selectionOptions.end.index),e.addRange(t),!0}};if(n.init(),t.copyFromSelector)if(t.html){const t=window.getSelection();t.removeAllRanges();const n=document.createRange(),o=document.querySelector(e).tagName;if("INPUT"===o||"TEXTAREA"===o){const t=document.createElement("div");document.body.append(t),t.innerHTML=document.querySelector(e).value;const n=window.getSelection();n.removeAllRanges();const o=document.createRange();o.selectNodeContents(t),n.addRange(o),document.execCommand("copy"),t.remove()}else n.selectNodeContents(document.querySelector(e)),t.addRange(n),document.execCommand("copy")}else{const t=document.createElement("textarea");document.body.append(t);const n=document.querySelector(e);t.value="TEXTAREA"===n.tagName||"INPUT"===n.tagName?n.value:n.innerText,t.select(),document.execCommand("copy"),t.remove()}else if(t.html){const t=document.createElement("div");document.body.append(t),t.innerHTML=e;const n=window.getSelection();n.removeAllRanges();const o=document.createRange();o.selectNodeContents(t),n.addRange(o),document.execCommand("copy"),t.remove()}else{const t=document.createElement("textarea");document.body.append(t),t.value=e,t.select(),document.execCommand("copy"),t.remove()}return t.reSelect&&n.reselect(),!0};
//# sourceMappingURL=/sm/955c4a48b08203f3686757bd707ce369c35113d2911c42968115c59d40973103.map