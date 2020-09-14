export const defaultHeader = {
    'Content-Type': 'application/json'
}

export const accountTypeParams = (accountType) => {
    return {key: 'account-type', value: accountType};
}

export const getParam = (key, value) => {
    return {key: key, value: value};
}
