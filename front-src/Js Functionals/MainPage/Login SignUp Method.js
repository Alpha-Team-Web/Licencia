const urlSignUp = "http://localhost:8008/register"
const urlLogin = "http://localhost:8008/login"


const signUpUsername = document.getElementById("SignUp-UserName")
const signUpFirstName = document.getElementById("SignUp-FirstName")
const signUpLastName = document.getElementById("SignUp-LastName")
const signUpEmail = document.getElementById("SignUp-Email")
const signUpPassword = document.getElementById("SignUp-Password")
const signUpRepeatPassword = document.getElementById("SignUp-RepeatPassword")
// const signUpIsFreeLancer = document.getElementById("isFreeLancer-ToggleButton")
const signupKind = document.getElementById("signUpKind")
function signUp() {
    var doc = hasEmpty(signUpUsername, signUpFirstName, signUpLastName, signUpEmail, signUpPassword, signUpRepeatPassword)
    if (doc != null) {
        setFieldError(doc.parentElement.parentElement)
        setTimeout(() => alert("fill the red box!!"), 1000);
    } else {
        const data = {
            username: signUpUsername.value,
            firstname: signUpFirstName.value,
            lastname: signUpLastName.value,
            email: signUpEmail.value,
            password: signUpPassword.value
        }
        /*const response = fetch(urlSignUp, {
            method: 'POST',
            mode: 'cors',
            cache: 'no-cache',
            credentials: 'same-origin',
            headers: {
                'Content-Type': 'application/json'
            },
            params: {
                account_type: signUpIsFreeLancer.value ? "freelancer" : "employer"
            },
            redirect: 'follow',
            referrerPolicy: 'no-referrer',
            body: JSON.stringify(data)
        }).then(successSignUp)
            .catch(denySignUp);*/
        const response = axios.post(urlSignUp, data, {
            params: {
                //TODO account Type
                'account-type' : signupKind.value
            }
        }).then(successSignUp)
            .catch(denySignUp)
    }
}
function successSignUp(res) {
    console.log(res)
    console.log("**** \n" + res.status)

    // todo alerting response message
    if (res.status === 200) {
        // todo go to login menu
    } else {
        // todo error the fields
    }
}
function denySignUp(res) {
    alert('Error Connecting To Licencia Server')
    console.log(res)
}


const loginKeypoint = document.getElementById("login-KeyPoint");
const loginPassword = document.getElementById("login-Password");
const loginKind = document.getElementById("loginKind")
function login() {
    var doc = hasEmpty(loginKeypoint, loginPassword)
    if (doc != null) {
        setFieldError(doc.parentElement.parentElement)
        setTimeout(() => alert("fill the red box!!"), 1000);
    } else {
        const data = {
            id: loginKeypoint.value,
            password: loginPassword.value
        }
        const response = axios.get(urlLogin, data, {
            params: {
                'account-type': loginKind.value
            }
        })
            .then(successLogin)
            .catch(denyLogin);
        Cookies.set('Fuck', "Holy Fucking Shit", {
            domain: "FuckFuckFuck",
            path: "FuckFuck"
        });
        window.location.href = profilePageName;
    }
}
function successLogin(res) {
    console.log("success");
    console.log(res)
    // todo alerting response message
    if (res.status === 200) {
        // todo go to Profile Menu And Save Auth
        window.location.href = profilePageName;
    } else {
        // todo error the fields
    }
}
function denyLogin(res) {
    alert('Error Connecting To Licencia Server')
    console.log(res)
}



function hasEmpty(...args) {
    for (let doc of args) {
        console.log(doc);
        if (doc.value === "") {
            return doc;
        }
    }
    return null;
}


function setFieldError(field) {
    if (field.class == null || !field.class.contains('error')) {
        field.style.border = "1px solid red";
        field.class = field.class == null ? 'error' : field.class + " error";
    }
}
