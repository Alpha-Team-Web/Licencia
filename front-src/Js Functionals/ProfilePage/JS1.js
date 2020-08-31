const usernameField = document.getElementById("usernameField");
const firstNameField = document.getElementById("firstNameField");
const lastNameField = document.getElementById("lastNameField");
const emailField = document.getElementById("emailField");
const siteAddressField = document.getElementById("siteAddressField");
const telephoneNumberField = document.getElementById("telephoneNumberField");
const passwordField = document.getElementById("passwordField");
const repeatPasswordField = document.getElementById("repeatPasswordField");
const gitHubAccountField = document.getElementById("githubAccountField");
const siteAddressField = document.getElementById("siteAddressField");
const descriptionField = document.getElementById("descriptionField");
const addressField = document.getElementById("addressField");

let username;
let firstname;
let lastname;
let email;
let siteAddress;
let telephoneNumber;
let password;
let gitHubAccount;
let gitHubRepo;
let description;
let address;

function logOut() {
    // Todo
    window.location.href = mainPageName;
}

const gitHubAccountPart = document.getElementById("gitHubAccountPart");

function loadProfileMenu() {
    httpGet(urlGetProfileInfo, {
        'Content-Type': 'application/json',
        'token': Cookies.get('auth')
    }, handleSuccessGetProfileInfo, handleDenyGetProfileInfo);
    if (!Cookies.get('isfreelancer')) {
        gitHubAccountPart.style.display = "none";


    }
}

function handleSuccessGetProfileInfo(value) {
    console.log("message : " + value.message);
    let messages = JSON.parse(value.message);
    username = messages.username;
    firstname = messages.firstname;
    lastname = messages.lastname;
    email = messages.email;
    description = messages.description;
    telephoneNumber = messages['telephone-number'];
    address = messages.address;
    password = messages.password;
    if (Cookies.get('isfreelancer')) {
        gitHubAccount = messages['github-account'];
        gitHubRepo = messages['github-repo'];
        siteAddress = messages['site-address'];
        fillFreelancerSpecialFields();
    } else {
        fillCommonFields();
    }

}

function fillFreelancerSpecialFields() {
    siteAddressField.value = siteAddress;
    //TODO : account and repos;
}

function fillCommonFields() {
    usernameField.value = username;
    firstNameField.value = firstname;
    lastNameField.value = lastname;
    emailField.value = email;
    telephoneNumberField.value = telephoneNumber;
    addressField.value = address;
    descriptionField.value = description;
    passwordField.value = password;
}

function handleDenyGetProfileInfo(value) {
    alert(JSON.stringify(value));
    console.log("raft too handleDenyGetProfileInfo");
}
