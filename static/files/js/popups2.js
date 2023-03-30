function openEmail() {
    document.getElementById("eChange").style.display = "block";
    document.getElementById("uChange").style.display = "none";
    document.getElementById("DELETE").style.display = "none";
    blur();
}



function openUsername() {
    document.getElementById("eChange").style.display = "none";
    document.getElementById("uChange").style.display = "block";
    document.getElementById("DELETE").style.display = "none";
    blur();

}



function openDelete() {
    document.getElementById("eChange").style.display = "none";
    document.getElementById("uChange").style.display = "none";
    document.getElementById("DELETE").style.display = "block";
    blur();

}



function closeAll() {
    document.getElementById("eChange").style.display = "none";
    document.getElementById("uChange").style.display = "none";
    document.getElementById("DELETE").style.display = "none";
    unblur();
}

function blur(){
    document.getElementById("blur").style.display = "block";
    document.getElementById("toblur1").style.filter = "blur(4px)";
    document.getElementById("toblur2").style.filter = "blur(4px)";
    document.getElementById("toblur3").style.filter = "blur(4px)";
}

function unblur(){
    document.getElementById("blur").style.display = "none";
    document.getElementById("toblur1").style.filter = "blur(0px)";
    document.getElementById("toblur2").style.filter = "blur(0px)";
    document.getElementById("toblur3").style.filter = "blur(0px)";
}

function logOut(username){
    if (username == "") {
        window.location.replace("/");
    }
}



// function edit(nmb) {
//     console.log(nmb);
    
// }