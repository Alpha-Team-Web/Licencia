import Cookies from "js-cookie";
import {httpExcGET} from "../AlphaAPI";
import {changePasswordUrlEmployer, changePasswordUrlFreeLancer} from "../urlNames";
import {isFreeLancer} from "./profilePageContent";
import {reload} from "../PageRouter";

let oldPasswordField
let newPasswordField
let repeatNewPasswordField
function fillPasswordFields() {
    oldPasswordField = document.getElementById("oldPasswordField");
    newPasswordField = document.getElementById("passwordField");
    repeatNewPasswordField = document.getElementById("repeatPasswordField");
}

export function changePassword() {
    fillPasswordFields();
    if (oldPasswordField.value === "" || newPasswordField.value === "" || repeatNewPasswordField.value === "") {
        alert("you have empty field")
    } else {
        if (newPasswordField.value !== repeatNewPasswordField.value) {
            alert("passwords doesn't match")
        } else {
            let data = {
                'old-password': oldPasswordField.value,
                'new-password': newPasswordField.value,
            }
            let headers = {
                'Content-Type': 'application/json',
                'token': Cookies.get('auth')
            }
            httpExcGET('POST', isFreeLancer ? changePasswordUrlFreeLancer : changePasswordUrlEmployer,
                data, successChangePassword, denyChangePassword, headers)
        }
    }
}

function successChangePassword(value) {
    alert("password changed successfully" + "  value : " + JSON.stringify(value))
    reload();
}

function denyChangePassword(value) {
    alert("password doesn't change" + "  value : " + JSON.stringify(value))
    // handleError
}
