import Cookies from 'js-cookie'

export function httpExcGET(method, url, data, handleSuccess, handleDeny, headers, ...params) {
    alert('url: ' + url)
    alert('method: ' + method)
    return fetch(url + createQuery(params), {
        method: method,
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        headers: headers,
        redirect: 'follow',
        referrerPolicy: 'no-referrer',
        body: JSON.stringify(data)
    }).then(response => success(response, handleSuccess, handleDeny))
        .catch(deny);
}

export function httpGet(url, headers, handleSuccess, handleDeny, ...params){
    return fetch(url + createQuery(params), {
        method: 'GET',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        headers: headers,
        redirect: 'follow',
        referrerPolicy: 'no-referrer',
    }).then(response => success(response, handleSuccess, handleDeny))
        .catch(response => deny(response));
}

export function httpExcGetFile(method, url, formData, handleSuccess, handleDeny, headers, ...params) {
    return fetch(url + createQuery(params), {
        method: method,
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        headers: headers,
        redirect: 'follow',
        referrerPolicy: 'no-referrer',
        body: formData
    }).then(response => success(response, handleSuccess, handleDeny))
        .catch(deny);
}

export async function checkURL(url) {
    let responseToURL = await fetch(url, {
        method: 'get',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        redirect: 'follow',
        referrerPolicy: 'no-referrer-when-downgrade',
    })
    return responseToURL.status !== 404;
}

function createQuery(params) {
    let query = "";
    let shekCreator = param => param.key + "=" + param.value;
    if (params.length > 0) {
        query += "?" + shekCreator(params[0]);
        for (let i = 1; i < params.length; i++) query += "&" + shekCreator(params[i]);
    }
    return query;
}


function success(response, handleSuccess, handleError) {
    alert("Connected to Server SuccessFully");
    // todo alerting response message
    if (!(response.headers.get('Token') === null || response.headers.get('Token') === undefined)) {
        alert('token Set.')
        Cookies.set('auth', response.headers.get('Token'));
        alert('auth: ' + Cookies.get('auth'))
    }
    response.json()
        .then(value => {
            if (response.status === 200) {
                // todo go to Profile Menu And Save Auth
                handleSuccess(value)
            } else {
                handleError(value)
                // todo error the fields
            }
        }).catch(reason => handleResponseJsonCatch(reason))
}


function deny(response) {
    alert('Error Connecting To Licencia Server')
    alert('response: ' + response)
    //todo
}

function handleResponseJsonCatch(reason) {
    alert("Raft To Catche Response.json()")
    alert("Reason: " + reason)
}

