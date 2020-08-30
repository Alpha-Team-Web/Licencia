const usernameField = document.getElementById("usernameField");
const firstNameField= document.getElementById("firstNameField");
const lastNameField= document.getElementById("lastNameField");
const emailField= document.getElementById("emailField");
const siteAddressField= document.getElementById("siteAddressField");
const telephoneNumberField= document.getElementById("telephoneNumberField");
const passwordField= document.getElementById("passwordField");
const repeatPasswordField= document.getElementById("repeatPasswordField");
const gitHubAccountField = document.getElementById("githubAccountField");
const descriptionField= document.getElementById("descriptionField");
const addressField= document.getElementById("addressField");
function logOut() {
    // Todo
    window.location.href = mainPageName;
}

const gitHubAccountPart = document.getElementById("gitHubAccountPart");
function loadProfileMenu() {
    if (false /* if is Not FreeLancer */) {
        gitHubAccountPart.style.display = "none";
    }
    console.log(Cookies.get('Fuck', {
        domain: "FuckFuckFuck",
        path: "FuckFuck"
    }));
}
