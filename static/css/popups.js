function openLogin() {
    document.getElementById("login").style.display = "block";
    document.getElementById("register").style.display = "none";
    document.getElementById("register2").style.display = "none";
}
function openRegister() {
    document.getElementById("login").style.display = "none";
    document.getElementById("register").style.display = "block";
    document.getElementById("register2").style.display = "none";
}
function openRegister2() {
    document.getElementById("register").style.display = "none";
    document.getElementById("register").style.display = "none";
    document.getElementById("register2").style.display = "block";
}
function closeRegister() {
    document.getElementById("register2").style.display = "none";
    document.getElementById("register").style.display = "none";
    document.getElementById("login").style.display = "none";
} 
function openLoginIf() {
    let loggedin = 1;
if (loggedin == 0) {
    document.getElementById("login").style.display = "block";
    document.getElementById("register").style.display = "none";
    document.getElementById("register2").style.display = "none";
    document.getElementById("lock").disabled = true;
    document.getElementById("lock1").disabled = true;
}
    
}
function test(userName) {
    if (userName == "") {
        document.getElementById("logbutton").style.display = "block";
        document.getElementById("profbutton").style.display = "none";
    } else{
        document.getElementById("logbutton").style.display = "none";
        document.getElementById("profbutton").style.display = "block";
    }
} 

function refreshPage(){
    window.location.reload();
}