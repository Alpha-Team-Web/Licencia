export const greenColor = "green"
export const whiteColor = "white"
let onLoginPage = false;
const loginButton = document.getElementById("LoginMenuButton")
const signupButton = document.getElementById("SignUpMenuButton")
const signupMenu = document.getElementById("SignUp-Menu")
const LoginMenu = document.getElementById("Login-Menu")

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
