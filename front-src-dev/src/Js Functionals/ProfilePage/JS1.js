import {httpExcGET, httpGet} from "../AlphaAPI";
import {mainPagePath} from "../PagePaths";
import Cookies from 'js-cookie';
import {
    changePasswordUrlEmployer,
    changePasswordUrlFreeLancer,
    gitHubUrl, saveGithubUrlFreeLancer,
    saveProfileUrlEmployer,
    saveProfileUrlFreeLancer,
    urlGetEmployerProfileInfo,
    urlGetFreelancerProfileInfo
} from "../urlNames";
import {fillLinksValues, fillLinksValuesToInputs, initGithubRepos} from "./linksContent";
import {fillProjectValues} from "./projectsContent";
import {fillCommonFields, fillForProfileFields, fillProfileValues} from "./profileContent";
import {goToPage} from "../PageRouter";

export let isFreeLancer = true;

export function logOut() {
    // Todo
    goToPage(mainPagePath)
}


export function initTransitionsStart(profileTransition, ...transitions) {
    alert('transitions: ' + transitions)
    // transitions.forEach((value => value.toggleVisibility()))
}

export function loadProfileMenu() {
    fillForProfileFields()
    // alert('IsFreeLancer: ' + Cookies.get('isfreelancer'))
    isFreeLancer = Cookies.get('isfreelancer');
    alert('Cookies: "' + isFreeLancer + "'")
    if (!isFreeLancer) {
        httpGet(urlGetEmployerProfileInfo, {
            'Content-Type': 'application/json',
            'Token': Cookies.get('auth')
        }, handleSuccessGetProfileInfo, handleDenyGetProfileInfo);
    } else {
        httpGet(urlGetFreelancerProfileInfo, {
            'Content-Type': 'application/json',
            'Token': Cookies.get('auth')
        }, handleSuccessGetProfileInfo, handleDenyGetProfileInfo);
    }
    initGithubRepos();
}

function handleSuccessGetProfileInfo(value) {
    let messages = value;
    alert(JSON.stringify(messages));
    fillProfileValues(messages);
    fillProjectValues(messages);
    fillCommonFields();
    if (isFreeLancer) {
        fillLinksValues(messages);
        fillLinksValuesToInputs();
    }
}

function handleDenyGetProfileInfo(value) {
    alert(JSON.stringify(value));
    console.log("raft too handleDenyGetProfileInfo");
}


let profileComponent;
let linksComponent;
function fillComponents() {
    profileComponent = document.getElementById('profileComponent')
    linksComponent = document.getElementById('linksComponent')
}


export function switchProfileToLinks() {
    fillComponents();
    linksComponent.style.display = "block"
    profileComponent.style.display = "none";
}

export function switchLinksToProfile() {
    fillComponents();
    linksComponent.style.display = "none"
    profileComponent.style.display = "block";
}
