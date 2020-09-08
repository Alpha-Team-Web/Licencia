import {httpExcGET, parseValue} from "../AlphaAPI";
import {isFreeLancer} from "./JS1";
import {saveProfileUrlEmployer, saveProfileUrlFreeLancer} from "../urlNames";

let username
let nickName
let firstName
let lastName
let email
let phoneNumber
let address
let description

function fillVar() {
    username = document.getElementById("username")
    nickName = document.getElementById("nickName")
    firstName = document.getElementById("firstName")
    lastName = document.getElementById("lastName")
    email = document.getElementById("email")
    phoneNumber = document.getElementById("phoneNumber")
    address = document.getElementById("address")
    description = document.getElementById("description")
}

let confirmForm = () => {
    let data = {
        username: username.value,
        nickName: nickName.value,
        firstName: firstName.value,
        lastName: lastName.value,
        email: email.value,
        phoneNumber: phoneNumber.value,
        address: address.value,
        description: description.value
    }

    httpExcGET("POST",isFreeLancer ? saveProfileUrlFreeLancer : saveProfileUrlEmployer,data,successSaveProfile,denySaveProfile,{
        'Content-Type': 'application/json'
    })
}

function successSaveProfile(value){
    alert(value)
}

function denySaveProfile(value){
    value = parseValue(value)
    console.log(value.message)
}