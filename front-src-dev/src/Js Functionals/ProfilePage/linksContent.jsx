import {gitHubUrl, saveGithubUrlFreeLancer} from "../urlNames";
import Cookies from "js-cookie";
import {checkURL, httpExcGET, httpGet} from "../AlphaAPI";
import {isFreeLancer} from "./profilePageContent";
import GithubRepoComponent from "../../Components/ProfilePageComponents/GithubRepoComponent";
import ReactDOM from 'react-dom';
import React from "react";

const REPOS_MAX_SIZE = 3;

let siteAddressField
let gitHubAccountField
let githubRepositoriesByFields = [];
export function fillRepoContentFields() {
    gitHubAccountField = document.getElementById("githubAccountField");
    siteAddressField = document.getElementById("siteAddressField");
}


export function fillLinksValuesToInputs() {
    alert('fuck')
    siteAddressField.value = siteAddress;
    gitHubAccountField.value = gitHubAccount;
    githubRepositoriesByFields = gitHubRepos;
    gitHubAccountChanged()

    if (gitHubRepos !== null) {
        createRepoDivs()
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
        githubRepositoriesByFields = [];
        createRepoDivs();
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
    if (addedRepoInput.value !== "") {
        addRepoByName(addedRepoInput.value)
    }

    addRepoDiv.style.display = 'none';
    addedRepoInput.value = '';
    determineAddRepoDivIcon();
}

async function addRepoByName(addRepoName) {
    let fuck = await checkURL(gitHubUrl + '/' + gitHubAccountField.value + "/" + addRepoName);
    if(!fuck) {
        alert("You Haven't Claimed a Valid Repo")
    } else if (githubRepositoriesByFields.length <= REPOS_MAX_SIZE && fuck) {
        if (githubRepositoriesByFields.includes(addRepoName)) {
            alert('Please Enter A New Repo')
        } else {
            githubRepositoriesByFields[githubRepositoriesByFields.length] = addRepoName;
            createRepoDivs();
        }
    }
}

function createRepoDivs() {
    let repoDivs = githubRepositoriesByFields.map(((value, index) =>
        <GithubRepoComponent repoName={value} id={'Repo' + index}
                             repoIndex={index}
                             href={gitHubUrl + '/' + gitHubAccountField.value + '/' + value} />))
    ReactDOM.render(repoDivs, gitHubReposContainer);
    determineAddRepoDivIcon();
}

let determineAddRepoDivIcon = () => {
    addRepoIconDiv.style.display = githubRepositoriesByFields.length === REPOS_MAX_SIZE ? 'none' : 'block';
}

export function removeRepository(repoIndex) {
    fillGitHubReposFields()
    if (repoIndex < githubRepositoriesByFields.length) {
        githubRepositoriesByFields = githubRepositoriesByFields.filter(((value, index) => index !== repoIndex))
        createRepoDivs()
        addRepoIconDiv.style.display = 'block';
    }
}



export function submitGitPart() {
    fillRepoContentFields()
    if (isFreeLancer) {
        let data = {
            'website': siteAddressField.value,
            'github-repos': githubRepositoriesByFields,
            'github-account': gitHubAccountField.value
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
