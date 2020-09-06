export const greenColor = "green"
export const whiteColor = "white"
let onLoginPage = true;
let loginButton = document.getElementById("LoginMenuButton")
let signupButton = document.getElementById("SignUpMenuButton")
let signupMenu = document.getElementById("SignUp-Menu")
let LoginMenu = document.getElementById("Login-Menu")

function emptyLoginFields() {

}

export function loginMenu() {
    if (onLoginPage) return;
    onLoginPage = true;
    emptyLoginFields();
    change(greenColor, whiteColor, "none", "block");
}

export function signUpMenu() {
    if (!onLoginPage) {
        return
    }
    onLoginPage = false;
    change(whiteColor, greenColor, "block", "none");
}

function change(color1, color2, display1, display2) {
    loginButton = document.getElementById("LoginMenuButton")
    signupButton = document.getElementById("SignUpMenuButton")
    signupMenu = document.getElementById("SignUp-Menu")
    LoginMenu = document.getElementById("Login-Menu")

    loginButton.style.backgroundColor = color1;
    signupButton.style.backgroundColor = color2;
    signupMenu.style.display = display1;
    LoginMenu.style.display = display2;
}



export const userNameMaxLength = 30;
export const passwordMaxLength = 30;
export const emailMaxLength = 50;
export const firstNameMaxLength = 30;
export const lastNameMaxLength = 40;
function signUpInput() {
    // todo
}

// signUpInput();
