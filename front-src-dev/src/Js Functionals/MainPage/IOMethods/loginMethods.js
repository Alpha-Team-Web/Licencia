import {httpExcGET} from "../../AlphaAPI";
import {urlLogin} from "../../urlNames";
import Cookies from "js-cookie";
import {goToPage} from "../../PageRouter";
import {profilePagePath} from "../../PagePaths";
import {loginInvalidPasswordLabel, loginNotSignedUpKeypointLabel} from "../ioErrors";
import {emptyFields, emptyFieldsFromErrors, hasEmpty} from "../../Utils/handleInputs";
import {setFieldError, showErrorLabel} from "../../Utils/handleErrors";

let loginKeypoint;
let loginPassword;
let loginKind;
let loginCloseModalFunc;
function setLoginFields() {
    loginKeypoint = document.getElementById("login-KeyPoint");
    loginPassword = document.getElementById("login-Password");
    loginKind = document.getElementById("loginKind");
}

export let emptyLoginFields = () => {
    setLoginFields()
    emptyFields(loginKeypoint, loginPassword, loginKind)
}

export let isLoginFieldsEmpty = () => hasEmpty(loginKeypoint, loginPassword, loginKind);
export let emptyLoginFieldsFromErrors = () => emptyFieldsFromErrors(loginKeypoint, loginPassword, loginKind)


function checkLoginFields() {
    setLoginFields()
    emptyLoginFieldsFromErrors();

    let doc = isLoginFieldsEmpty();
    if (doc != null) {
        setFieldError(doc, true)
        showErrorLabel(doc, 'fill It, Dude')
        return false;
    }

    return true;
}

let getLoginDataFromFields = () => {
    setLoginFields();
    return {
        id: loginKeypoint.value,
        password: loginPassword.value
    }
}

export function login(func) {
    loginCloseModalFunc = func;

    if (checkLoginFields()) {
        const promise = httpExcGET('post', urlLogin, getLoginDataFromFields(), handleSuccessLogin, handleErrorLogin, {
            'Content-Type': 'application/json'
        }, {
            key: 'account-type',
            value: loginKind.value
        });
    }
}

function handleSuccessLogin(value) {
// todo go to Profile Menu And Save Auth
    alert("Login Successful")
    Cookies.set("isfreelancer", loginKind.value === "freelancer");
    goToPage(profilePagePath);
    emptyLoginFields();
}

function handleErrorLogin(value) {
    // todo error the fields
    let message = value.message
    alert("Login Failed")
    alert('Server Message: ' + message)

    switch (message) {
        case 'not signed up username':
        case 'not signed up email':
            setFieldError(loginKeypoint)
            setFieldError(loginPassword)
            showErrorLabel(loginKeypoint, loginNotSignedUpKeypointLabel)
            break;
        case 'invalid password':
            setFieldError(loginPassword)
            showErrorLabel(loginPassword, loginInvalidPasswordLabel)
            break;
        default:
            alert("Haven't Handled That Error Before");
            console.log("messageError: '" + value.message + "'")
    }
}
