function openLogin() {
    document.getElementById("login").style.display = "block";
    document.getElementById("register").style.display = "none";
    document.getElementById("register2").style.display = "none";
    blur();
}
function openRegister() {
    document.getElementById("login").style.display = "none";
    document.getElementById("register").style.display = "block";
    document.getElementById("register2").style.display = "none";
    blur();
}
function openRegister2() {
    document.getElementById("register").style.display = "none";
    document.getElementById("register").style.display = "none";
    document.getElementById("register2").style.display = "block";
    blur();
}
function closeRegister() {
    document.getElementById("register2").style.display = "none";
    document.getElementById("register").style.display = "none";
    document.getElementById("login").style.display = "none";
    unblur();
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
        document.getElementById("lock").disabled = true;
        document.getElementById("lock1").disabled = true;
        document.getElementById("lock2").disabled = true;
    } else{
        document.getElementById("logbutton").style.display = "none";
        document.getElementById("profbutton").style.display = "block";
        document.getElementById("lock").disabled = false;
        document.getElementById("lock1").disabled = false;
        document.getElementById("lock2").disabled = false;
    }
} 

function refreshPage() {
    window.location.reload();
}   

function edit(nmb) {
    console.log(nmb);
}

function blur(){
    document.getElementById("blurrr").style.display = "block";
    document.getElementById("filter").style.filter = "blur(4px)";
    document.getElementById("search").style.filter = "blur(4px)";
    document.getElementById("post").style.filter = "blur(4px)";
}

function unblur(){
    document.getElementById("blurrr").style.display = "none";
    document.getElementById("filter").style.filter = "blur(0px)";
    document.getElementById("search").style.filter = "blur(0px)";
    document.getElementById("post").style.filter = "blur(0px)";
}
