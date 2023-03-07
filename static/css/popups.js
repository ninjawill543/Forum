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

function edit(){
    document.getElementById("edit").style.display = "block";
    document.getElementById("toedit").style.display = "none";

    var input = document.getElementById("enter");

// Execute a function when the user presses a key on the keyboard
input.addEventListener("keypress", function(event) {
  // If the user presses the "Enter" key on the keyboard
  if (event.key === "Enter") {
    // Cancel the default action, if needed
    event.preventDefault();
    // Trigger the button element with a click
    document.getElementById("edit").style.display = "none";
    document.getElementById("toedit").style.display = "block";
  }
}); 
}

function editClose(){
    document.getElementById("edit").style.display = "block";
}
