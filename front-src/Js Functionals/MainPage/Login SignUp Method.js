const signUpUsername = document.getElementById("SignUp-UserName")
const signUpFirstName = document.getElementById("SignUp-FirstName")
const signUpLastName = document.getElementById("SignUp-LastName")
const signUpEmail = document.getElementById("SignUp-Email")
const signUpPassword = document.getElementById("SignUp-Password")
const signUpRepeatPassword = document.getElementById("SignUp-RepeatPassword")
const signupKind = document.getElementById("signUpKind")

function signUp() {
    var doc = hasEmpty(signUpUsername, signUpFirstName, signUpLastName, signUpEmail, signUpPassword, signUpRepeatPassword)
    if (doc != null) {
        setFieldError(doc, true)
        setTimeout(() => alert("fill the red box!!"), 1000);
    } else {
        if (signUpPassword.value !== signUpRepeatPassword.value) {
            alert("Your Passwords Doesn't Match")
            setFieldError(signUpPassword, true)
            setFieldError(signUpRepeatPassword, true)
            return;
        }
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

function handleErrorSignUp(value) {
    // todo error the fields
    alert("SignUp Failed")
    alert('Server Message: ' + value.message)
    switch (value.messageError) {
        case 'duplicate email':
            setFieldError(signUpEmail, true);
            break;
        case 'duplicate username':
            setFieldError(signUpUsername, true)
            break;
        default:
            alert("Haven't Handled That Error Before");
            console.log("messageError: '" + value.messageError + "'")
    }
}


const loginKeypoint = document.getElementById("login-KeyPoint");
const loginPassword = document.getElementById("login-Password");
const loginKind = document.getElementById("loginKind")

function login() {
    let doc = hasEmpty(loginKeypoint, loginPassword);
    if (doc != null) {
        setFieldError(doc, true)
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
    let splitter = value.message.indexOf(':');
    let messageError = value.message.substring(0, splitter);
    let messageField = value.message.substring(splitter + 1);
    alert("Login Successful")
    alert("Server Message: " + value.message)
    alert("Auth: " + value.messageField)
    Cookies.set('auth', value.messageField)
    Cookies.set('isfreelancer', loginKind.value === 'freelancer')
    window.location.href = profilePageName;
}

function handleErrorLogin(value) {
    // todo error the fields
    let splitter = value.message.indexOf(':');
    let messageError = value.message.substring(0, splitter);
    let messageField = value.message.substring(splitter + 1);
    alert("Login Failed")
    alert('Server Message: ' + value.message)
    switch (value.messageError) {
        case 'not signed up username':
        case 'not signed up email':
            setFieldError(loginKeypoint);
            break;
        case 'invalid password':
            setFieldError(loginPassword)
            break;
        default:
            alert("Haven't Handled That Error Before");
            console.log("messageError: '" + value.messageError + "'")
    }
}

function hasEmpty(...args) {
    for (let doc of args) {
        if (doc.value === "") {
            return doc;
        }
    }
    return null;
}

function setFieldError(field, isError) {
    if ((isError === undefined || isError) && !field.parentElement.classList.contains("error")) {
        // field.style.border = "1px solid red";
        field.parentElement.classList.add("error");
    } else if (!isError && field.parentElement.classList.contains("error")) {
        field.parentElement.classList.remove("error")
    }
}
