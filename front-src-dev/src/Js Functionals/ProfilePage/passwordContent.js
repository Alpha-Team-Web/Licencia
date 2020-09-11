import Cookies from "js-cookie";
import {httpExcGET} from "../AlphaAPI";
import {changePasswordUrlEmployer, changePasswordUrlFreeLancer} from "../urlNames";
import {isFreeLancer} from "./profilePageContent";
import {reload} from "../PageRouter";
import {setFieldError, showErrorLabel} from "../Utils/handleErrors";
import {incorrectOldPasswordLabel, newPasswordMisMatch} from "./Utils/registerErrors";
import {emptyFieldsFromErrors, hasEmpty} from "../Utils/handleInputs";

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

    emptyFieldsFromErrors(oldPasswordField, newPasswordField, repeatNewPasswordField)
    let emptyField = hasEmpty(oldPasswordField, newPasswordField, repeatNewPasswordField)
    if (emptyField !== null) {
        setFieldError(emptyField)
        showErrorLabel(emptyField, 'fill This Dude')
    } else {
        if (newPasswordField.value !== repeatNewPasswordField.value) {
            setFieldError(newPasswordField)
            showErrorLabel(newPasswordField, newPasswordMisMatch)
            setFieldError(repeatNewPasswordField)
            showErrorLabel(repeatNewPasswordField, newPasswordMisMatch)
        } else {
            let data = {
                'old-pass': oldPasswordField.value,
                'new-pass': newPasswordField.value,
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
    switch (value.message) {
        case 'password mismatch':
            setFieldError(oldPasswordField)
            showErrorLabel(oldPasswordField, incorrectOldPasswordLabel)
            break;
        default:
            alert("haven't Handled This Error Before.")
            break;
    }
}
