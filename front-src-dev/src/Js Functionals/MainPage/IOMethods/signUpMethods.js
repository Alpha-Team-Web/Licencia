import MainTextField from "../../../Components/MainPageComponents/mainTextField";
import {httpExcGET} from "../../AlphaAPI";
import {urlSignUp} from "../../urlNames";
import {signUpDuplicateEmailLabel, signUpDuplicateUsernameLabel, signUpInvalidEmailLabel} from "../ioErrors";
import {emptyFieldsFromErrors, hasEmpty} from "./Utils/handleInputs";
import {setFieldError, showErrorLabel} from "./Utils/handleErrors";

let signUpUsername
let signUpFirstName
let signUpLastName
let signUpEmail
let signUpPassword
let signUpRepeatPassword
let signUpKind
let signUpCloseModalFunc;
function setSignUpFields() {
    signUpUsername = document.getElementById("SignUp-UserName")
    signUpFirstName = document.getElementById("SignUp-FirstName")
    signUpLastName = document.getElementById("SignUp-LastName")
    signUpEmail = document.getElementById("SignUp-Email")
    signUpPassword = document.getElementById("SignUp-Password")
    signUpRepeatPassword = document.getElementById("SignUp-RepeatPassword")
    signUpKind = document.getElementById("signUpKind")
}

export function emptySignUpFields() {
    setSignUpFields();
    signUpUsername.value = "";
    signUpFirstName.value = "";
    signUpLastName.value = "";
    signUpEmail.value = "";
    signUpPassword.value = "";
    signUpRepeatPassword.value = "";
    signUpKind.value = "";
}

export let isSignUpInputsEmpty = () => hasEmpty(signUpUsername, signUpFirstName, signUpLastName, signUpEmail,
    signUpPassword, signUpRepeatPassword)

export let emptySignUpFieldsFromErrors = () => emptyFieldsFromErrors(signUpUsername, signUpFirstName, signUpLastName,
        signUpEmail, signUpPassword, signUpRepeatPassword);


function checkSignUpFields() {

}

export function signUp(func) {
    signUpCloseModalFunc = func
    setSignUpFields();
    let doc = isSignUpInputsEmpty();
    emptySignUpFieldsFromErrors()
    if (doc != null) {
        setFieldError(doc, true)
        setTimeout(() => alert("fill the red box!!"), 1000);
    } else {
        if (signUpPassword.value !== signUpRepeatPassword.value) {
            alert("Your Passwords Doesn't Match")
            setFieldError(signUpPassword)
            setFieldError(signUpRepeatPassword)
            return;
        }
        const data = {
            username: signUpUsername.value,
            'firstname': signUpFirstName.value,
            'lastname': signUpLastName.value,
            email: signUpEmail.value,
            password: signUpPassword.value
        }
        alert('data: ' + JSON.stringify(data))
        const promise = httpExcGET('post', urlSignUp, data, handleSuccessSignUp, handleErrorSignUp, {
            'Content-Type': 'application/json'
        }, {
            key: 'account-type',
            value: signUpKind.value
        });
    }
}

function handleSuccessSignUp(value) {
    alert("SignUp Successful")
    emptySignUpFields();
    signUpCloseModalFunc()
    // closeTheFuckinModal
}

function handleErrorSignUp(value) {
    // todo error the fields
    alert("SignUp Failed")
    alert('Server Message: ' + value.message)
    switch (value.message) {
        case 'duplicate email':
            setFieldError(signUpEmail);
            showErrorLabel(signUpEmail, signUpDuplicateEmailLabel);
            break;
        case 'invalid email':
            setFieldError(signUpEmail);
            showErrorLabel(signUpEmail, signUpInvalidEmailLabel)
            break;
        case 'duplicate username':
            setFieldError(signUpUsername)
            showErrorLabel(signUpUsername, signUpDuplicateUsernameLabel)
            break;
        default:
            alert("Haven't Handled That Error Before");
            console.log("messageError: '" + value.message + "'")
    }
}
