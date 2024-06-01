window.addEventListener("load", function(){
    let anchor = document.getElementById("card-output-files");
    if (typeof(anchor) != 'undefined' && anchor != null)
    {
        setTimeout(function () {
            anchor.scrollIntoView({behavior: "smooth"});
        }, 500);
    }
    
});