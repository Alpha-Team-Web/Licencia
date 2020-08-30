function showLoginMenu() {
    $('#LogSin-Menu')
        .modal('show')
    ;
    loginMenu();
}

const greenColor = "green"
const whiteColor = "white"
let onLoginPage = false;
const loginButton = document.getElementById("LoginMenuButton")
const signupButton = document.getElementById("SignUpMenuButton")
const signupMenu = document.getElementById("SignUp-Menu")
const LoginMenu = document.getElementById("Login-Menu")

function loginMenu() {
    if (onLoginPage) return;
    onLoginPage = true;
    change(greenColor, whiteColor, "none", "block");
}

function signUpMenu() {
    if (!onLoginPage) {
        return
    }
    onLoginPage = false;
    change(whiteColor, greenColor, "block", "none");
}

function change(color1, color2, display1, display2) {
    loginButton.style.backgroundColor = color1;
    signupButton.style.backgroundColor = color2;
    signupMenu.style.display = display1;
    LoginMenu.style.display = display2;
}



const userNameMaxLength = 30;
const passwordMaxLength = 30;
const emailMaxLength = 50;
const firstNameMaxLength = 30;
const lastNameMaxLength = 40;
function signUpInput() {
    // todo
}
// signUpInput();
