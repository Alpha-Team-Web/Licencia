import {httpExcGET} from "../AlphaAPI";
import {saveProfileUrlEmployer, saveProfileUrlFreeLancer} from "../urlNames";
import Cookies from "js-cookie";
import {isFreeLancer} from "./JS1";
import {reload} from "../PageRouter";

let usernameField
let shownNameField
let firstNameField
let lastNameField
let emailField
let telephoneNumberField
let descriptionField
let addressField
export function fillForProfileFields() {
    usernameField = document.getElementById("usernameField");
    shownNameField = document.getElementById('showingNameField')
    firstNameField = document.getElementById("firstNameField");
    lastNameField = document.getElementById("lastNameField");
    emailField = document.getElementById("emailField");
    telephoneNumberField = document.getElementById("telephoneNumberField");
    descriptionField = document.getElementById("descriptionField");
    addressField = document.getElementById("addressField");
}

let username;
let shownName;
let firstname;
let lastname;
let email;
let telephoneNumber;
let description;
let address;
export function fillProfileValues(messages) {
    username = messages.username;
    shownName = messages['shown-name']
    firstname = messages['first-name'];
    lastname = messages['last-name'];
    email = messages.email;
    description = messages.description;
    telephoneNumber = messages['phone-number'];
    address = messages.address;
}

export function fillCommonFields() {
    usernameField.value = username;
    shownNameField.value = shownName;
    firstNameField.value = firstname;
    lastNameField.value = lastname;
    emailField.value = email;
    telephoneNumberField.value = telephoneNumber;
    addressField.value = address;
    descriptionField.value = description;
}

export function saveProfile() {
    let getValue = (firstValue, secondValue) => secondValue == null ? firstValue : secondValue;
    const data = {
        'shown-name': getValue(shownName, shownNameField.value),
        'first-name': getValue(firstname, firstNameField.value),
        'last-name': getValue(lastname, lastNameField.value),
        'phone-number': telephoneNumberField.value,
        'address': addressField.value,
        'description': descriptionField.value
    }
    alert('data: ' + JSON.stringify(data))
    telephoneNumber = telephoneNumberField.value;
    address = addressField.value;
    description = descriptionField.value;
    httpExcGET('post', isFreeLancer ? saveProfileUrlFreeLancer : saveProfileUrlEmployer,
        data, successSaveProfile, errorSaveProfile, {
            'Token': Cookies.get('auth')
        })
}
function successSaveProfile(value) {
    alert('Profile Saved Successfully')
    reload();
}
function errorSaveProfile(value) {
    //Error Handling
    alert('We Have An Error')
    alert('error: ' + value.message)
}

