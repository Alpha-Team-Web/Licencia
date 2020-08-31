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
let projectsId;
let requestedProjectsId;

function logOut() {
    // Todo
    window.location.href = mainPageName;
}

const gitHubAccountPart = document.getElementById("gitHubAccountPart");

function loadProfileMenu() {

    if (!Cookies.get('isfreelancer')) {
        httpGet(urlGetEmployerProfileInfo, {
            'Content-Type': 'application/json',
            'token': Cookies.get('auth')
        }, handleSuccessGetProfileInfo, handleDenyGetProfileInfo);
        gitHubAccountPart.style.display = "none";
    } else {
        httpGet(urlGetFreelancerProfileInfo, {
            'Content-Type': 'application/json',
            'token': Cookies.get('auth')
        }, handleSuccessGetProfileInfo, handleDenyGetProfileInfo);
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
    telephoneNumber = messages.phonenumber;
    address = messages.addr;
    password = messages.password;
    projectsId = messages['project-ids'];
    if (Cookies.get('isfreelancer')) {
        gitHubAccount = messages.github;
        gitHubRepo = messages['github-repos'];
        siteAddress = messages.website;
        requestedProjectsId = messages['req-project-ids'];
        fillFreelancerSpecialFields();
    } else {
        fillCommonFields();
    }

}

function fillFreelancerSpecialFields() {
    siteAddressField.value = siteAddress;
    //TODO : remaining;
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
