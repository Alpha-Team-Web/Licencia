import {setFieldError, showErrorLabel} from "./handleErrors";

export function hasEmpty(...args) {
    for (let doc of args) {
        if (doc.value === "") {
            return doc;
        }
    }
    return null;
}

export function emptyFieldsFromErrors(...fields) {
    fields.forEach((value => {
        if (value) {
            setFieldError(value, false)
            showErrorLabel(value, false)
        }
    }))
}
