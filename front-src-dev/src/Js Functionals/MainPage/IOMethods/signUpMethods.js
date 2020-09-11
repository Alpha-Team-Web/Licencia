import MainInput from "../../../Components/MainPageComponents/mainFormElements/mainInput";
import {httpExcGET} from "../../AlphaAPI";
import {urlSignUp} from "../../urlNames";
import {signUpDuplicateEmailLabel, signUpDuplicateUsernameLabel, signUpInvalidEmailLabel} from "../ioErrors";
import {emptyFields, emptyFieldsFromErrors, hasEmpty} from "./Utils/handleInputs";
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

export let emptySignUpFields = () => {
    setSignUpFields();
    emptyFields(signUpUsername, signUpFirstName, signUpLastName,
        signUpEmail, signUpPassword, signUpRepeatPassword, signUpKind)
}

export let isSignUpInputsEmpty = () => hasEmpty(signUpUsername, signUpFirstName, signUpLastName, signUpEmail,
    signUpPassword, signUpRepeatPassword, signUpKind)

export let emptySignUpFieldsFromErrors = () => emptyFieldsFromErrors(signUpUsername, signUpFirstName, signUpLastName,
        signUpEmail, signUpPassword, signUpRepeatPassword, signUpKind);


function checkSignUpFields() {
    setSignUpFields();
    let doc = isSignUpInputsEmpty();
    emptySignUpFieldsFromErrors()
    if (doc != null) {
        setFieldError(doc, true)
        showErrorLabel(doc, 'fill It, Dude')
        return false;
    }
    return true;
}

let getSignUpDataFromFields = () => {
    setSignUpFields()
    return {
        username: signUpUsername.value,
        'firstname': signUpFirstName.value,
        'lastname': signUpLastName.value,
        email: signUpEmail.value,
        password: signUpPassword.value
    }
}

function checkPasswordSimilarity() {
    if (signUpPassword.value !== signUpRepeatPassword.value) {
        let error = "Your Passwords Doesn't Match";
        setFieldError(signUpPassword)
        setFieldError(signUpRepeatPassword)
        showErrorLabel(signUpPassword, error)
        showErrorLabel(signUpRepeatPassword, error)
        return false;
    }
    return true;
}

export function signUp(func) {
    signUpCloseModalFunc = func

     if(checkSignUpFields() && checkPasswordSimilarity()) {
        const promise = httpExcGET('post', urlSignUp, getSignUpDataFromFields(),
            handleSuccessSignUp, handleErrorSignUp, {
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
