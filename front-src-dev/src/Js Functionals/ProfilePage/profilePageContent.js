import {httpExcGET, httpGet} from "../AlphaAPI";
import {mainPagePath} from "../PagePaths";
import Cookies from 'js-cookie';
import {
    urlGetEmployerProfileInfo,
    urlGetFreelancerProfileInfo
} from "../urlNames";
import {fillLinksValues, fillLinksValuesToInputs, fillRepoContentFields} from "./linksContent";
import {fillProjectValues} from "./projectsContent";
import {fillCommonFields, fillForProfileFields, fillProfileValues} from "./profileContent";
import {goToPage} from "../PageRouter";
import {fillProfileImage} from "./profilePictureContent";

export let isFreeLancer = true;

export function logOut() {
    // Todo
    Cookies.remove('auth')
    goToPage(mainPagePath)
}

export function loadProfileMenu() {
    fillForProfileFields()
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
}

function handleSuccessGetProfileInfo(value) {
    value = value.data
    alert(JSON.stringify(value));
    fillProfileValues(value[0]);
    fillProfileImage(value[1])
    fillProjectValues(value[0]);
    fillCommonFields();
    if (isFreeLancer) {
        fillRepoContentFields();
        fillLinksValues(value[0]);
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
