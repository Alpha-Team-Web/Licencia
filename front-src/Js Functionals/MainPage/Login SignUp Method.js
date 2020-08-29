const urlSignUp = "http://localhost:8008/register"
const urlLogin = "http://localhost:8008/login"


const signUpUsername = document.getElementById("SignUp-UserName")
const signUpFirstName = document.getElementById("SignUp-FirstName")
const signUpLastName = document.getElementById("SignUp-LastName")
const signUpEmail = document.getElementById("SignUp-Email")
const signUpPassword = document.getElementById("SignUp-Password")
const signUpRepeatPassword = document.getElementById("SignUp-RepeatPassword")
const signUpIsFreeLancer = document.getElementById("isFreeLancer-ToggleButton")

function signUp() {
    var doc = hasEmpty(signUpUsername, signUpFirstName, signUpLastName, signUpEmail, signUpPassword, signUpRepeatPassword)
    if (doc != null) {
        doc.style.border = "1px solid red";
        setTimeout(() => alert("fill the red box!!"), 1000);
    } else {
        const data = {
            username: signUpUsername,
            firstname: signUpFirstName,
            lastname: signUpLastName,
            email: signUpEmail,
            password: signUpPassword
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
                'account-type' : (signUpIsFreeLancer.checked ? 'freelancer' : 'employer')
            }
        }).then(successSignUp)
            .catch(denySignUp)
    }
}

function successSignUp(res) {
    console.log(res)
    console.log("ok")
}

function denySignUp(res) {
    console.log("riidii")
    console.log(res)
}


const loginKeypoint = document.getElementById("login-KeyPoint");
const loginPassword = document.getElementById("login-Password");

function login() {
    var doc = hasEmpty(loginKeypoint, loginPassword)
    if (doc != null) {
        doc.style.border = "1px solid red";
        setTimeout(() => alert("fill the red box!!"), 1000);
    } else {
        const data = {
            keyPoint: loginKeypoint,
            password: loginPassword
        }
        /*const response = fetch(urlLogin, {
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
        }).then(successLogin)
            .catch(denyLogin);*/
        const response = axios.get(urlLogin, data)
            .then(successLogin)
            .catch(denyLogin);
        window.location.href = profilePageName;
    }
}

function successLogin(res) {
    console.log("success");
    // Todo
    window.location.href = profilePageName;
}

function denyLogin(res) {
    console.log("riidiii");
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