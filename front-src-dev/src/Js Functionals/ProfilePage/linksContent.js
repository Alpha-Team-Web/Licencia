import {gitHubUrl, saveGithubUrlFreeLancer} from "../urlNames";
import Cookies from "js-cookie";
import {httpExcGET} from "../AlphaAPI";
import {isFreeLancer} from "./JS1";

export let repoDiv
export let iconDiv
export let repoInput
export let firstRepoDiv
export let secondRepoDiv
export let thirdRepoDiv
export let firstRepoLink
export let secondRepoLink
export let thirdRepoLink
export let gitHubReposDiv
export function fillRepoContentDivs() {
    repoDiv = document.getElementById("addRepoDiv");
    iconDiv = document.getElementById("plusRepoIconDiv");
    repoInput = document.getElementById("addRepoInput");
    firstRepoDiv = document.getElementById("firstRepo");
    secondRepoDiv = document.getElementById("secondRepo");
    thirdRepoDiv = document.getElementById("thirdRepo");
    firstRepoLink = document.getElementById("linkRepo1");
    secondRepoLink = document.getElementById("linkRepo2");
    thirdRepoLink = document.getElementById("linkRepo3");
    gitHubReposDiv = document.getElementById("gitHubRepos");
}

export let siteAddressField
export let githubAccountField
export function fillRepoContentFields() {
    githubAccountField = document.getElementById("githubAccountField");
    siteAddressField = document.getElementById("siteAddressField");
}

export function initGithubRepos() {
    fillRepoContentDivs();
    firstRepoDiv.style.display = 'none';
    secondRepoDiv.style.display = 'none';
    thirdRepoDiv.style.display = 'none';
    iconDiv.style.display = 'none';
}

export function fillLinksValuesToInputs() {
    siteAddressField.value = siteAddress;
    //TODO : remaining;
}


let siteAddress;
let gitHubAccount;
let gitHubRepos;

export function fillLinksValues(messages) {
    gitHubAccount = messages.github;
    gitHubRepos = messages['github-repos'];
    siteAddress = messages.website;
}

export function openAddRepoDiv() {
    fillRepoContentDivs()
    let counter = 0;
    if (firstRepoDiv.style.display !== "none") counter += 1;
    if (secondRepoDiv.style.display !== "none") counter += 1;
    if (thirdRepoDiv.style.display !== "none") counter += 1;
    if (counter === 3) {
        alert("you can have only 3 repository")
        return;
    }
    repoDiv.style.display = "block";
    iconDiv.style.display = "none";
    repoInput.value = "";
    repoInput.focus();
}


export function closeAddRepoDiv() {
    fillRepoContentDivs()
    if (repoInput.value === "") {
        // $ep
        repoDiv.style.display = "none";
        iconDiv.style.display = "block";
        return;
    }
    if (firstRepoDiv.style.display === "none") {
        showRepo(firstRepoDiv, firstRepoLink, repoInput.value)
    } else if (secondRepoDiv.style.display === "none") {
        showRepo(secondRepoDiv, secondRepoLink, repoInput.value)
    } else if (thirdRepoDiv.style.display === "none") {
        showRepo(thirdRepoDiv, thirdRepoLink, repoInput.value)
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

function showRepo(repoDiv, repoLink, text) {
    fillRepoContentFields()
    let textNode = document.createTextNode(text);
    removeAllChild(repoLink);
    repoLink.appendChild(textNode);
    repoLink.href = gitHubUrl + githubAccountField.value + '/' + text
    repoDiv.style.display = "block";
}

export function removeRepo(element) {
    fillRepoContentDivs()
    element.style.display = "none";
    iconDiv.style.display = "block";
}

function removeAllChild(element) {
    while (element.firstChild) {
        element.removeChild(element.firstChild);
    }
}

export function accountGithubChanged() {
    fillRepoContentDivs()
    fillRepoContentFields()
    if (githubAccountField.value === "") {
        firstRepoDiv.style.display = "none";
        secondRepoDiv.style.display = "none";
        thirdRepoDiv.style.display = "none";
        repoDiv.style.display = "none";
        iconDiv.style.display = "none";
    } else {
        iconDiv.style.display = "block";
    }
}

export function submitGitPart() {
    fillRepoContentDivs()
    fillRepoContentFields()
    if (isFreeLancer) {
        let gitLinks = [];
        let size = 0;
        if (firstRepoDiv.style.display !== "none") {
            gitLinks[size] = document.getElementById('linkRepo1').text();
            size += 1;
        }
        if (secondRepoDiv.style.display !== "none") {
            gitLinks[size] = document.getElementById('linkRepo2').text();
            size += 1;
        }
        if (thirdRepoDiv.style.display !== "none") {
            gitLinks[size] = document.getElementById('linkRepo3').text();
            size += 1;
        }
        let data = {
            'website': siteAddressField.value,
            'github-repos': gitLinks,
            'github': githubAccountField.value
        }
        siteAddress = siteAddressField.value;
        gitHubAccount = githubAccountField.value;
        let headers = {
            'Content-Type': 'application/json',
            'token': Cookies.get('auth')
        }
        httpExcGET('POST', saveGithubUrlFreeLancer, data, successGithubPartSubmit, denyGithubPartSubmit, headers);
    }
}

function successGithubPartSubmit(value) {
    alert("post successfully" + " value : " + JSON.stringify(value));
    // eslint-disable-next-line no-restricted-globals
    location.reload();
}

function denyGithubPartSubmit(value) {
    alert("post deny" + " value : " + JSON.stringify(value));
    // handleErrors
}
