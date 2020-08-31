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
        setFieldError(doc)
        setTimeout(() => alert("fill the red box!!"), 1000);
    } else {
        const data = {
            username: signUpUsername.value,
            firstname: signUpFirstName.value,
            lastname: signUpLastName.value,
            email: signUpEmail.value,
            password: signUpPassword.value
        }
        const promise = httpExcGET('post', urlSignUp, data, handleSuccessSignUp, handleErrorSignUp, {
            key: 'account-type',
            value: signupKind.value
        });
    }
}

function handleSuccessSignUp(value) {
    alert("SignUp Successful")
    shideLoginMenu(false)
}

function handleErrorSignUp(reason) {
    // todo error the fields
    alert("SignUp Failed")
}


const loginKeypoint = document.getElementById("login-KeyPoint");
const loginPassword = document.getElementById("login-Password");
const loginKind = document.getElementById("loginKind")

function login() {
    var doc = hasEmpty(loginKeypoint, loginPassword)
    if (doc != null) {
        setFieldError(doc)
        setTimeout(() => alert("fill the red box!!"), 1000);
    } else {
        const data = {
            id: loginKeypoint.value,
            password: loginPassword.value
        }
        const promise = httpExcGET('post', urlLogin, data, handleSuccessLogin, handleErrorLogin, {
            key: 'account-type',
            value: loginKind.value
        });
    }
}

function handleSuccessLogin(value) {
// todo go to Profile Menu And Save Auth
    alert("Login Successful")
    window.location.href = profilePageName;
}

function handleErrorLogin(reason) {
    // todo error the fields
    alert("Login Failed")
}

/*function http(method, url, data, success, deny, ...params) {
    return fetch(url + createQuery(params), {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        headers: {
            'Content-Type': 'application/json'
        },
        redirect: 'follow',
        referrerPolicy: 'no-referrer',
        body: JSON.stringify(data)
    }).then(success)
        .catch(deny);
}*/

function hasEmpty(...args) {
    for (let doc of args) {
        if (doc.value === "") {
            return doc;
        }
    }
    return null;
}

function setFieldError(field) {
    if (!field.parentElement.classList.contains("error")) {
        // field.style.border = "1px solid red";
        field.parentElement.classList.add("error");
    }
}


function printResponse(response) {
    response.json()
        .then(value => console.log("value: '" + value.message + "'"))
        .catch(reason => console.log("reason: " + reason))
}
