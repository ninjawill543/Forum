function openEmail() {
    document.getElementById("eChange").style.display = "block";
    document.getElementById("pChange").style.display = "none";
    document.getElementById("uChange").style.display = "none";
    document.getElementById("DELETE").style.display = "none";
}



function openUsername() {
    document.getElementById("eChange").style.display = "none";
    document.getElementById("pChange").style.display = "none";
    document.getElementById("uChange").style.display = "block";
    document.getElementById("DELETE").style.display = "none";
}



function openPassword() {
    document.getElementById("eChange").style.display = "none";
    document.getElementById("pChange").style.display = "block";
    document.getElementById("uChange").style.display = "none";
    document.getElementById("DELETE").style.display = "none";
}


function openDelete() {
    document.getElementById("eChange").style.display = "none";
    document.getElementById("pChange").style.display = "none";
    document.getElementById("uChange").style.display = "none";
    document.getElementById("DELETE").style.display = "block";
}



function closeAll() {
    document.getElementById("eChange").style.display = "none";
    document.getElementById("pChange").style.display = "none";
    document.getElementById("uChange").style.display = "none";
    document.getElementById("DELETE").style.display = "none";
}