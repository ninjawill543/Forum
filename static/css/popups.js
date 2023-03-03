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
    document.getElementById("register2").style.display = "block";
    document.getElementById("register").style.display = "none";
}

function closeRegister() {
    document.getElementById("register2").style.display = "none";
    document.getElementById("register").style.display = "none";
    document.getElementById("login").style.display = "none";
} 

console.log("Register");