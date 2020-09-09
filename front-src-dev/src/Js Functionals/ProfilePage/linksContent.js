import {gitHubUrl, saveGithubUrlFreeLancer} from "../urlNames";
import Cookies from "js-cookie";
import {httpExcGET} from "../AlphaAPI";
import {isFreeLancer} from "./JS1";
import GithubRepoComponent from "../../Components/ProfilePageComponents/GithubRepoComponent";
import ReactDOM from 'react-dom';
import React from "react";

let siteAddressField
let gitHubAccountField
export function fillRepoContentFields() {
    gitHubAccountField = document.getElementById("githubAccountField");
    siteAddressField = document.getElementById("siteAddressField");
}


export function fillLinksValuesToInputs() {
    siteAddressField.value = siteAddress;
    gitHubAccountField.value = gitHubAccount;
    gitHubAccountChanged()

    if (gitHubRepos !== null) {
        gitHubRepos.forEach((value) => addRepoByName(value))
    }
    //TODO : remaining;
}

let siteAddress;
let gitHubAccount;
let gitHubRepos;
export function fillLinksValues(messages) {
    gitHubAccount = messages['github-account']
    gitHubRepos = messages['github-repos'];
    siteAddress = messages.website;
}

export function submitGitPart() {
    fillRepoContentFields()
    if (isFreeLancer) {
        let data = {
            'website': siteAddressField.value,
            'github-repos': githubRepositoriesByFields,
            'github': gitHubAccountField.value
        }
        siteAddress = siteAddressField.value;
        gitHubAccount = gitHubAccountField.value;
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


// -----------------------------------------------------------------------------------


const REPOS_MAX_SIZE = 3;

let githubRepositoriesByFields = [];
let gitHubRepoDivs = [];

let gitHubReposContainer
let addRepoDiv
let addedRepoInput
let addRepoIconDiv
function fillGitHubReposFields() {
    gitHubReposContainer = document.getElementById('gitHubRepositories');
    addRepoDiv = document.getElementById('addRepoDiv')
    addedRepoInput = document.getElementById('addRepoInput')
    addRepoIconDiv = document.getElementById('plusRepoIconDiv')
}

export function gitHubAccountChanged() {
    fillGitHubReposFields()
    if (gitHubAccountField.value !== "") {
        addRepoIconDiv.style.display = 'block';
    } else {
        addRepoIconDiv.style.display = 'none';
    }
}

export function clickedPlusIcon() {
    fillGitHubReposFields();
    addRepoDiv.style.display = 'block';
    addedRepoInput.focus();
    addRepoIconDiv.style.display = 'none';
}


export function addedRepoInputFocusOut() {
    fillGitHubReposFields();
    if (addedRepoInput.value === "") {
        addRepoDiv.style.display = 'none';
        addRepoIconDiv.style.display = 'block';
    } else {
        addRepoByName(addedRepoInput.value)
    }
}

function addRepoByName(addRepoName) {
    fillGitHubReposFields()
    if (githubRepositoriesByFields.length <= REPOS_MAX_SIZE) {
        let addedRepoDiv = <GithubRepoComponent repoName={addRepoName} repoIndex={githubRepositoriesByFields.length}
                                                href={gitHubUrl + '/' + gitHubAccount + '/' + addRepoName}/>
        githubRepositoriesByFields[githubRepositoriesByFields.length] = addRepoName;
        gitHubRepoDivs[gitHubRepoDivs.length] = addedRepoDiv
        ReactDOM.render(addedRepoDiv, gitHubReposContainer);
    } else {
        alert('gitHubReposSize More Than Specified')
    }

    if (githubRepositoriesByFields.length === REPOS_MAX_SIZE) {
        addRepoIconDiv.style.display = 'none';
    }
}

export function removeRepository(repoIndex) {
    fillGitHubReposFields()
    if (repoIndex < gitHubRepoDivs.length) {
        let removingRepo = gitHubRepoDivs[repoIndex]
        gitHubRepoDivs = gitHubRepoDivs.filter(((value, index) => index !== repoIndex))
        githubRepositoriesByFields = githubRepositoriesByFields.filter(((value, index) => index !== repoIndex))
        gitHubReposContainer.removeChild(removingRepo)
    }
}
