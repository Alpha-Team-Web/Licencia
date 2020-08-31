const usernameField = document.getElementById("usernameField");
const firstNameField = document.getElementById("firstNameField");
const lastNameField = document.getElementById("lastNameField");
const emailField = document.getElementById("emailField");
const siteAddressField = document.getElementById("siteAddressField");
const telephoneNumberField = document.getElementById("telephoneNumberField");
const passwordField = document.getElementById("passwordField");
const repeatPasswordField = document.getElementById("repeatPasswordField");
const gitHubAccountField = document.getElementById("githubAccountField");
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

const repoDiv = document.getElementById("addRepoDiv");
const iconDiv = document.getElementById("plusRepoIconDiv");
const repoInput = document.getElementById("addRepoInput");
const firstRepoDiv = document.getElementById("firstRepo");
const secondRepoDiv = document.getElementById("secondRepo");
const thirdRepoDiv = document.getElementById("thirdRepo");
const firstRepoLink = document.getElementById("linkRepo1");
const secondRepoLink = document.getElementById("linkRepo2");
const thirdRepoLink = document.getElementById("linkRepo3");
const githubAccountField = document.getElementById("githubAccountField");
const gitHubReposDiv = document.getElementById("gitHubRepos");

function openAddRepoDiv() {
    let counter = 0;
    if (firstRepoDiv.style.display !== "none") counter += 1;
    if (secondRepoDiv.style.display !== "none") counter += 1;
    if (thirdRepoDiv.style.display !== "none") counter += 1;
    if (counter === 3) {
        alert("you can have only 3 repository")
        return;
    }
    repoInput.value = "";
    repoDiv.style.display = "block";
    iconDiv.style.display = "none";
}


function closeAddRepoDiv() {
    if (repoInput.value === "") return;
    let textNode = document.createTextNode(repoInput.value);
    if (firstRepoDiv.style.display === "none") {
        removeAllChild(firstRepoLink);
        firstRepoLink.appendChild(textNode);
        firstRepoDiv.style.display = "block";
    } else if (secondRepoDiv.style.display === "none") {
        removeAllChild(secondRepoLink);
        secondRepoLink.appendChild(textNode);
        secondRepoDiv.style.display = "block";
    } else if (thirdRepoDiv.style.display === "none") {
        removeAllChild(thirdRepoLink);
        thirdRepoLink.appendChild(textNode);
        thirdRepoDiv.style.display = "block";
    } else {
        alert("you can have only 3 repository")
        return;
    }
    repoDiv.style.display = "none";
    iconDiv.style.display = "block";
    repoInput.value = "";
    let counter = 0;
    if (firstRepoDiv.style.display !== "none") counter += 1;
    if (secondRepoDiv.style.display !== "none") counter += 1;
    if (thirdRepoDiv.style.display !== "none") counter += 1;
    if (counter === 3) {
        iconDiv.style.display = "none";
    }
}

function removeRepo(element) {
    element.style.display = "none";
    iconDiv.style.display = "block";
}

function removeAllChild(element) {
    while (element.firstChild) {
        element.removeChild(element.firstChild);
    }
}

function accountGithubChanged() {
    if (githubAccountField.value === "") {
        gitHubReposDiv.style.display = "none";
        firstRepoDiv.style.display = "none";
        secondRepoDiv.style.display = "none";
        thirdRepoDiv.style.display = "none";
    } else {
        gitHubReposDiv.style.display = "block"
    }
}
